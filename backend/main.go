package main

import (
	"github.com/praveenvoonna/weather-app/backend/config"
	"github.com/praveenvoonna/weather-app/backend/server"
)

func main() {
	config.LoadEnv()
	server.StartServer()
}
