package backend

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/kataras/golog"
)

type Backend struct {
	host string
	port int
	tran *Transport
	pers *Persist
}

func NewBackend(host string, port int, friends []string, p *Persist) *Backend {
	s := Backend{
		host: host,
		port: port,
		tran: NewTransport(friends),
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
	golog.Info(string(bytes))
	s.pers.Add(string(bytes) + "\n")
	remoteIP := s.GetRemoteIp(r)
	if s.tran.NeedToTransfer(remoteIP) {
		go s.tran.Transfer(string(bytes))
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
