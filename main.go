package main

import (
	"github.com/neo-ngd/LogServer/api"
	"github.com/neo-ngd/LogServer/backend"
	"github.com/neo-ngd/LogServer/config"
	"github.com/neo-ngd/LogServer/storage"
)

const (
	localhost = "0.0.0.0"
)

func main() {
	storage := storage.NewStorage(
		config.LogPath,
		config.LogName,
		config.LogFileExpire,
		config.LogFileSplit,
	)
	s := api.NewApiServer(localhost, config.SPort, storage)
	b := backend.NewBackend(config.Name, localhost, config.RPort, config.Friends, storage)
	b.Start()
}
