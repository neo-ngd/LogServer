package backend

import (
	"net/http"
	"strings"

	"github.com/kataras/golog"
)

type Transport struct {
	friends []string
}

func NewTransport(friends []string) *Transport {
	t := Transport{
		friends: friends,
	}
	return &t
}
func (t *Transport) Transfer(log string) {
	for _, v := range t.friends {
		go func(iport string) {
			reader := strings.NewReader(log)
			url := "http://" + iport
			req, err := http.NewRequest(http.MethodPost, url, reader)
			if err != nil {
				golog.Error(err)
				return
			}
			req.Header.Set("role", "friend")
			client := http.Client{}
			_, err = client.Do(req)
			if err != nil {
				golog.Warn(err)
			}
		}(v)
	}
}

func (t *Transport) NeedToTransfer(ip string) bool {
	for _, v := range t.friends {
		fip := strings.Split(v, ":")[0]
		if fip == ip {
			return false
		}
	}
	return true
}
