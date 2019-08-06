package config

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/kataras/golog"
)

var (
	SPort         int
	RPort         int
	LogPath       string
	LogName       string
	LogFileExpire time.Duration
	LogFileSplit  time.Duration
)

func init() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		golog.Fatal(err)
	}
	var config struct {
		SPort         int    `json:"sport"`
		RPort         int    `json:"rport"`
		LogPath       string `json:"logpath"`
		LogName       string `json:"logname"`
		LogFileExpire int    `json:"logfileexpire"`
		LogFileSplit  int    `json:"logfilesplit"`
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		golog.Error(err)
	}
	SPort = config.SPort
	RPort = config.RPort
	LogPath = config.LogPath
	LogName = config.LogName
	LogFileExpire = time.Duration(24*config.LogFileExpire) * time.Hour
	LogFileSplit = time.Duration(config.LogFileSplit) * time.Hour
	if LogName == "" || LogPath == "" {
		golog.Fatal("invalid LogPath or LogName")
	}
}
