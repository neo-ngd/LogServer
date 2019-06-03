package backend

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/kataras/golog"
	"github.com/neo-ngd/LogServer/api"
	"github.com/neo-ngd/LogServer/storage"
)

type Backend struct {
	name string
	host string
	port int
	tran *Transport
	sto  *storage.Storage
	srv  *http.Server
}

func NewBackend(name, host string, port int, friends []string, p *storage.Storage) *Backend {

	b := Backend{
		name: name,
		host: host,
		port: port,
		tran: NewTransport(name, friends),
		sto:  p
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/log", s.handler)
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
	go s.sto.Add(fmt.Sprintf("[%s]", name), log+"\n")
}

func (s *Backend) GetRemoteIp(r *http.Request) string {
	remoteAddr := r.RemoteAddr
	remoteIP, _, _ := net.SplitHostPort(remoteAddr)

	if remoteIP == "::1" {
		remoteIP = "127.0.0.1"
	}
	return remoteIP
}
