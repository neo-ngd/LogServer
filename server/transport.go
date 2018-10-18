package server

import (
	"net/http"
	"strings"

	"github.com/kataras/golog"
)

type Transport struct {
	friends []string
}

func NewTransport(friends []string) *Transport {
	t := Transport{}
	return &t
}
func (t *Transport) Transfer(log string) {
	for _, v := range t.friends {
		go func(url string) {
			reader := strings.NewReader(log)
			_, err := http.Post(url, "", reader)
			if err != nil {
				golog.Warn(err)
			}
		}(v)
	}
}

func (t *Transport) NeedToTransfer(ip string) bool {
	golog.Info("[transport][needtotransfer] ", ip)
	for _, v := range t.friends {
		fip := strings.Split(v, ":")[0]
		if fip == ip {
			return false
		}
	}
	return true
}
