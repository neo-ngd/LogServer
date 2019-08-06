package backend

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"

	"github.com/kataras/golog"
	"github.com/neo-ngd/LogServer/storage"
)

type Backend struct {
	host    string
	port    int
	sto     *storage.Storage
	srv     *http.Server
	unkowns map[string]string
}

func NewBackend(host string, port int, p *storage.Storage) *Backend {

	b := Backend{
		host:    host,
		port:    port,
		sto:     p,
		unkowns: make(map[string]string, 7),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/log", b.handler)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	b.srv = srv
	return &b
}

func (s *Backend) Start() {
	golog.Infof("start receive srv listening: %d", s.port)
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
		ip := s.GetRemoteIp(r)
		if n, ok := s.unkowns[ip]; ok {
			name = n
		} else {
			name = fmt.Sprintf("unkown_%d", rand.Int31())
			s.unkowns[ip] = name
		}
	}
	//persist
	go s.sto.Append(fmt.Sprintf("%s", name), log+"\n")
}

func (s *Backend) GetRemoteIp(r *http.Request) string {
	remoteAddr := r.RemoteAddr
	remoteIP, _, _ := net.SplitHostPort(remoteAddr)

	if remoteIP == "::1" {
		remoteIP = "127.0.0.1"
	}
	return remoteIP
}
