package api

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/kataras/golog"
	"github.com/neo-ngd/LogServer/api"
)

type ApiServer struct {
	host string
	port int
	tran *Transport
	sto  *storage.Storage
	so  *soserver
	srv  *http.Server
}

func NewApiServer(host string, port int, p *storage.Storage) *ApiServer {
	sosrv, err := newSoServer()
	if err != nil {
		golog.Fatal(err)
	}
	s := Backend{
		host: host,
		port: port,
		sto: 	p,
		so:  sosrv
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/socket.io/", s.so.ServeHTTP)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	s.srv = srv
	return &s
}

func (s *ApiServer)Start() {
	golog.Infof("start api srv listening: %d", s.port)

	if err := s.srv.ListenAndServe(); err != nil {
		golog.Fatal("ListenAndServe: ", err)
	}
}
