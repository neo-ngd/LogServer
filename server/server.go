package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/kataras/golog"
)

type Server struct {
	host string
	port int
	tran *Transport
}

func NewServer(host string, port int, friends []string) *Server {
	s := Server{
		host: host,
		port: port,
		tran: NewTransport(friends),
	}
	return &s
}

func (s *Server) Start() {
	log.Printf("server start listening: %d", s.port)
	http.HandleFunc("/", s.handler)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.host, s.port), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		w.WriteHeader(http.StatusOK)
	}()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(bytes))
	golog.Info(r.URL.String(), r.RequestURI, r.RemoteAddr)
	remoteIP := s.GetRemoteIp(r)
	if s.tran.NeedToTransfer(remoteIP) {
		go s.tran.Transfer(string(bytes))
	}
}

func (s *Server) GetRemoteIp(r *http.Request) string {
	remoteAddr := r.RemoteAddr
	remoteIP, _, _ := net.SplitHostPort(remoteAddr)

	if remoteIP == "::1" {
		remoteIP = "127.0.0.1"
	}
	return remoteIP
}
