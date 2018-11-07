package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/kataras/golog"
)

var (
	Friends       []string
	Cnport        int
	Webport       int
	LogPath       string
	LogName       string
	LogFileExpire time.Duration
	LogFileSplit  time.Duration
	Name          string
)

func init() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		golog.Fatal(err)
	}
	var config struct {
		Name          string   `json:"name"`
		Cnport        int      `json:"cnport"`
		Webport       int      `json:"webport"`
		Friends       []string `json:"sendto"`
		LogPath       string   `json:"logpath"`
		LogName       string   `json:"logname"`
		LogFileExpire int      `json:"logfileexpire"`
		LogFileSplit  int      `json:"logfilesplit"`
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		golog.Error(err)
	}
	Cnport = config.Cnport
	Webport = config.Webport
	Friends = config.Friends
	LogPath = config.LogPath
	LogName = config.LogName
	LogFileExpire = time.Duration(24*config.LogFileExpire) * time.Hour
	LogFileSplit = time.Duration(config.LogFileSplit) * time.Hour
	Name = config.Name
	if LogName == "" || LogPath == "" {
		golog.Fatal("invalid LogPath or LogName")
	}
	if Name == "" || Name == "local" {
		golog.Fatal("please set a name. tip: dont use \"local\"")
	}
	for _, v := range Friends {
		arr := strings.Split(v, ":")
		if (arr[0] == "127.0.0.1" || arr[0] == "" ||
			arr[0] == "localhost" || arr[0] == "0.0.0.0") &&
			(arr[1] == string(Cnport) || arr[1] == string(Webport)) {
			golog.Fatal("you cannt send log to yourself.")
		}
	}
}
