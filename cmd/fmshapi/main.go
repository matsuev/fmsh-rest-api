package main

import (
	"fmsh-rest-api/internal/config"
	"log"
)

func main() {
	println("FMSH REST API Server")

	cfg, err := config.New("./configs/config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	app := NewApplication(cfg)

	app.Start()

	app.Shutdown()
}
