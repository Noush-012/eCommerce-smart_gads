package main

import (
	"log"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/config"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/di"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Load configuration failed")
	}

	server, err := di.InitiateAPI(config)

	if err != nil {
		log.Fatal("Failed to start server")
	}
	server.Run()

}
