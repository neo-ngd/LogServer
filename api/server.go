package api

import (
	"net/http"

	"github.com/kataras/golog"

	socketio "github.com/googollee/go-socket.io"
)

const cache_count = 50

type logBody struct {
	Text string
	Name string
	Key  int64
}
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
	golog.Info("on connection.")
	so.Join("log")
	so.On("log:subscribe", func(msg string) {
		cache := s.cache.GetCached()
		for _, v := range cache {
			so.Emit("log:log", v)
		}
		golog.Info("on subscribe", msg)
	})
	so.On("disconnection", func() {
		golog.Info("on disconnect")
	})
}

func (s *SoServer) errHandler(so socketio.Socket, err error) {
	golog.Error(err)
}

func NewServer() (*SoServer, error) {
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

func (s *SoServer) SendLog(name, log string) {
	l := logBody{
		Name: name,
		Text: log,
	}
	s.cache.Append(l)
	s.sosrv.BroadcastTo("log", "log:log", l)
}
