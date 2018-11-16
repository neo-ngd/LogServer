package api

import (
	"fmt"
	"net/http"

	"github.com/kataras/golog"

	socketio "github.com/googollee/go-socket.io"
)

type logBody struct {
	Text string
	Tag  string
	Key  int64
}
type SoServer struct {
	sosrv *socketio.Server
	srv   *http.Server
	port  int
}

func (s *SoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	s.sosrv.ServeHTTP(w, r)
}

func handler(so socketio.Socket) {
	golog.Info("on connection.")
	so.Join("log")
	so.On("log:subscribe", func(msg string) {
		golog.Info("on subscribe", msg)
	})
	so.On("disconnection", func() {
		golog.Info("on disconnect")
	})
}

func errHandler(so socketio.Socket, err error) {
	golog.Info(err)
}

func NewServer(port int) (*SoServer, error) {
	s, e := socketio.NewServer(nil)
	if e != nil {
		golog.Error(e)
		return nil, nil
	}
	s.On("connection", handler)
	s.On("error", errHandler)
	server := &SoServer{
		sosrv: s,
		port:  port,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/socket.io/", server.ServeHTTP)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	server.srv = srv
	return server, nil
}

func (s *SoServer) Start() {
	golog.Infof("Serving webapi at localhost: %d...", s.port)
	go func() {
		golog.Fatal(s.srv.ListenAndServe())
	}()
}

func (s *SoServer) SendLog(log string) {
	golog.Info("sending log: ", log)
	l := logBody{
		Text: log,
	}
	s.sosrv.BroadcastTo("log", "log:log", l)
}
