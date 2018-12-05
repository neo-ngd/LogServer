package main

import (
	"github.com/neo-ngd/LogServer/backend"
	"github.com/neo-ngd/LogServer/config"
)

const (
	localhost = "0.0.0.0"
)

func main() {
	p := backend.NewPersist(
		config.LogPath,
		config.LogName,
		config.LogFileExpire,
		config.LogFileSplit,
	)
	s := backend.NewBackend(config.Name, localhost, config.Port, config.Friends, p)
	s.Start()
}
