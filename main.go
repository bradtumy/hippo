package main

import (
	"github/bradtumy/hippo/cmd/hippo"
	"github/bradtumy/hippo/config"
	"github/bradtumy/hippo/routes"

	"log"
)

func main() {
	cfg, err := config.New("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	hippo := hippo.New(cfg)
	router := routes.NewRouter(hippo)
	hippo.Startup(router)
}
