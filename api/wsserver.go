package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/kataras/golog"
	"github.com/neo-ngd/LogServer/storage"
)

type wsServer struct {
	upgrader websocket.Upgrader
	sto      *storage.Storage
}

func newWsServer(sto *storage.Storage) *wsServer {
	return &wsServer{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		sto: sto,
	}
}

func (ws *wsServer) handler(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		golog.Println(err)
		return
	}
	//... Use conn to send and receive messages.
	go func(conn *websocket.Conn) {
		for {
			mtype, data, err := conn.ReadMessage()
			if err != nil {
				golog.Error(err)
				conn.Close()
				return
			}
			if mtype == websocket.TextMessage {
				name := string(data)
				golog.Infof("subscriber: %s", name)
				if _, ok := findSocket(conn); ok {
					removeSubscriber(conn)
				}
				addSubscriber(name, conn)
				cache := ws.sto.GetCached(name)
				for _, v := range cache {
					data, err := json.Marshal(v)
					if err != nil {
						continue
					}
					conn.WriteMessage(websocket.TextMessage, data)
				}
			}
		}
	}(conn)
}

func (ws *wsServer) sendLog(name, log string) {
	l := storage.LogBody{
		Name: name,
		Text: log,
	}
	distribute(name, l)
}
