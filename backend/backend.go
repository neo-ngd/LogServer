package backend

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/neo-ngd/LogServer/api"

	"github.com/kataras/golog"
)

type Backend struct {
	name string
	host string
	port int
	tran *Transport
	pers *Persist
	api  *api.SoServer
	srv  *http.Server
}

func NewBackend(name, host string, port int, friends []string, p *Persist) *Backend {
	apisrv, err := api.NewServer()
	if err != nil {
		golog.Fatal(err)
	}
	s := Backend{
		name: name,
		host: host,
		port: port,
		tran: NewTransport(name, friends),
		pers: p,
		api:  apisrv,
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/log", s.handler)
	mux.HandleFunc("/socket.io/", apisrv.ServeHTTP)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	s.srv = srv
	return &s
}

func (s *Backend) Start() {
	golog.Infof("start listening: %d", s.port)

	if err := s.srv.ListenAndServe(); err != nil {
		golog.Fatal("ListenAndServe: ", err)
	}

}

func (s *Backend) handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		w.WriteHeader(http.StatusOK)
	}()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(string(bytes))
		return
	}
	log := string(bytes)
	name := r.Header.Get("From")
	if name == "" {
		name = s.name
		s.tran.Transfer(log)
	}
	//persist
	s.pers.Add(fmt.Sprintf("[%s]", name), log+"\n")

	//send to web
	go s.api.SendLog(name, log)
}

func (s *Backend) GetRemoteIp(r *http.Request) string {
	remoteAddr := r.RemoteAddr
	remoteIP, _, _ := net.SplitHostPort(remoteAddr)

	if remoteIP == "::1" {
		remoteIP = "127.0.0.1"
	}
	return remoteIP
}
