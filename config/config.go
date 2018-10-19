package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kataras/golog"
)

//Friends Friends are log-backends you want to send logs to
var Friends []string

//Port server port
var Port int

func init() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		golog.Fatal(err)
	}
	var config struct {
		Port    int      `json:"port"`
		Friends []string `json:"friends"`
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		golog.Error(err)
	}
	Port = config.Port
	Friends = config.Friends
}
