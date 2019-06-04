package api

import (
	"github.com/gorilla/websocket"
	"github.com/neo-ngd/LogServer/storage"
)

var distributor = make(map[string][]*websocket.Conn)

func addSubscriber(name string, s *websocket.Conn) {
	ss, ok := distributor[name]
	if !ok {
		distributor[name] = []*websocket.Conn{s}
	}
	distributor[name] = append(ss, s)
}

func findSocket(so *websocket.Conn) (string, bool) {
	for name, ss := range distributor {
		for _, s := range ss {
			if s == so {
				return name, true
			}
		}
	}
	return "", false
}
func removeSubscriber(so *websocket.Conn) {
	name, ok := findSocket(so)
	if !ok {
		return
	}
	ss := distributor[name]
	ts := make([]*websocket.Conn, len(ss)-1)
	j := 0
	for _, v := range ss {
		if v == so {
			continue
		}
		ts[j] = v
		j++
	}
	distributor[name] = ts
}

func distribute(name string, log storage.LogBody) {
	ss, ok := distributor[name]
	if ok {
		for _, s := range ss {
			s.WriteJSON(log)
		}
	}
	ss, ok = distributor["all"]
	if ok {
		for _, s := range ss {
			s.WriteJSON(log)
		}
	}
}
