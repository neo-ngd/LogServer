package api

import (
	"net/http"

	"github.com/kataras/golog"

	socketio "github.com/googollee/go-socket.io"
)

const cache_count = 50

type SoServer struct {
	sosrv *socketio.Server
	cache *logcache
}

func (s *SoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	s.sosrv.ServeHTTP(w, r)
}

func (s *SoServer) handler(so socketio.Socket) {
	so.Join("log")
	so.On("log:subscribe", func(name string) {
		if _, ok := FindSocket(so); ok {
			RemoveSubscriber(so)
		}
		AddSubscriber(name, so)
		cache := s.cache.GetCached(name)
		for _, v := range cache {
			so.Emit("log:log", v)
		}
		golog.Info("on subscribe: ", name)
	})
	so.On("disconnection", func() {
		RemoveSubscriber(so)
		golog.Info("on disconnect")
	})
}

func (s *SoServer) errHandler(so socketio.Socket, err error) {
	golog.Error(err)
}

func newSoServer() (*SoServer, error) {
	s, e := socketio.NewServer(nil)
	if e != nil {
		golog.Error(e)
		return nil, nil
	}
	server := &SoServer{
		sosrv: s,
		cache: NewLogCache(cache_count),
	}
	server.sosrv.On("connection", server.handler)
	server.sosrv.On("error", server.errHandler)
	return server, nil
}

func (s *SoServer) sendLog(name, log string) {
	l := logBody{
		Name: name,
		Text: log,
	}
	s.cache.Append(name, l)
	Distribute(name, l)
}
