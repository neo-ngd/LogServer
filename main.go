package main

import (
	"log-backend/config"
	"log-backend/server"
)

const (
	localhost = "0.0.0.0"
)

func main() {
	s := server.NewServer(localhost, config.Port, config.Friends)
	s.Start()
}
