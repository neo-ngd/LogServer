package api

import (
	"net/http"

	"github.com/kataras/golog"

	socketio "github.com/googollee/go-socket.io"
)

type logBody struct {
	Text string
	Name string
	Key  int64
}
type SoServer struct {
	sosrv *socketio.Server
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

func NewServer() (*SoServer, error) {
	s, e := socketio.NewServer(nil)
	if e != nil {
		golog.Error(e)
		return nil, nil
	}
	s.On("connection", handler)
	s.On("error", errHandler)
	server := &SoServer{
		sosrv: s,
	}
	return server, nil
}

func (s *SoServer) SendLog(name, log string) {
	l := logBody{
		Name: name,
		Text: log,
	}
	s.sosrv.BroadcastTo("log", "log:log", l)
}
