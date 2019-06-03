package api

import (
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/kataras/golog"
	"github.com/neo-ngd/LogServer/storage"
)

const cache_count = 50

type soServer struct {
	sosrv *socketio.Server
	sto   *storage.Storage
}

func (s *soServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	s.sosrv.ServeHTTP(w, r)
}

func (s *soServer) handler(so socketio.Socket) {
	so.Join("log")
	so.On("log:subscribe", func(name string) {
		if _, ok := findSocket(so); ok {
			removeSubscriber(so)
		}
		addSubscriber(name, so)
		cache := s.sto.GetCached(name)
		for _, v := range cache {
			so.Emit("log:log", v)
		}
		golog.Info("on subscribe: ", name)
	})
	so.On("disconnection", func() {
		removeSubscriber(so)
		golog.Info("on disconnect")
	})
}

func (s *soServer) errHandler(so socketio.Socket, err error) {
	golog.Error(err)
}

func newSoServer(sto *storage.Storage) (*soServer, error) {
	s, e := socketio.NewServer(nil)
	if e != nil {
		golog.Error(e)
		return nil, nil
	}
	server := &soServer{
		sosrv: s,
		sto:   sto,
	}
	server.sosrv.On("connection", server.handler)
	server.sosrv.On("error", server.errHandler)
	return server, nil
}

func (s *soServer) sendLog(name, log string) {
	l := storage.LogBody{
		Name: name,
		Text: log,
	}
	distribute(name, l)
}
