package api

import socketio "github.com/googollee/go-socket.io"

var distributor = make(map[string][]socketio.Socket)

func AddSubscriber(name string, s socketio.Socket) {
	ss, ok := distributor[name]
	if !ok {
		distributor[name] = []socketio.Socket{s}
	}
	distributor[name] = append(ss, s)
}

func FindSocket(so socketio.Socket) (string, bool) {
	for name, ss := range distributor {
		for _, s := range ss {
			if s == so {
				return name, true
			}
		}
	}
	return "", false
}
func RemoveSubscriber(so socketio.Socket) {
	name, ok := FindSocket(so)
	if !ok {
		return
	}
	ss := distributor[name]
	ts := make([]socketio.Socket, len(ss)-1)
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

func Distribute(name string, log logBody) {
	ss, ok := distributor[name]
	if ok {
		for _, s := range ss {
			s.Emit("log:log", log)
		}
	}
}
