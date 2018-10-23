package backend

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/kataras/golog"
)

type Backend struct {
	name string
	host string
	port int
	tran *Transport
	pers *Persist
}

func NewBackend(name, host string, port int, friends []string, p *Persist) *Backend {
	s := Backend{
		name: name,
		host: host,
		port: port,
		tran: NewTransport(name, friends),
		pers: p,
	}
	return &s
}

func (s *Backend) Start() {
	golog.Infof("start listening: %d", s.port)
	http.HandleFunc("/", s.handler)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.host, s.port), nil); err != nil {
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
	name := r.Header.Get("From")
	if name == "" {
		name = "local"
	}
	log := string(bytes)
	//persist
	s.pers.Add(fmt.Sprintf("[%s]", name), log+"\n")
	//transfer
	if name == "local" {
		s.tran.Transfer(log)
	}
}

func (s *Backend) GetRemoteIp(r *http.Request) string {
	remoteAddr := r.RemoteAddr
	remoteIP, _, _ := net.SplitHostPort(remoteAddr)

	if remoteIP == "::1" {
		remoteIP = "127.0.0.1"
	}
	return remoteIP
}
