package main

import (
	"forum/internal/app"
	"forum/internal/config"
	"log"
)

func main() {
	config := config.MustLoad()

	init, err := app.NewApp(config)
	if err != nil {
		log.Fatal("failed to start app")
	}

	init.Run()
}
