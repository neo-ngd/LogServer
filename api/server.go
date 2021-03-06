package api

import (
	"fmt"
	"net/http"

	"github.com/neo-ngd/LogServer/storage"

	"github.com/kataras/golog"
)

type ApiServer struct {
	host string
	port int
	sto  *storage.Storage
	so   *wsServer
	srv  *http.Server
	c    chan storage.LogBody
}

func NewApiServer(host string, port int, p *storage.Storage) *ApiServer {
	wssrv := newWsServer(p)
	s := ApiServer{
		host: host,
		port: port,
		sto:  p,
		so:   wssrv,
		c:    make(chan storage.LogBody, 10),
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/ws/", s.so.handler)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	s.srv = srv
	return &s
}

func (s *ApiServer) Start() {
	s.sto.Regist(s.c)
	golog.Infof("start api srv listening: %d", s.port)
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			golog.Fatal("ListenAndServe: ", err)
		}
	}()
	go func() {
		for {
			l, ok := <-s.c
			if !ok {
				golog.Infof("channel closed.")
				return
			}
			s.so.sendLog(l.Name, l.Text)
		}
	}()
}
