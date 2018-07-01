package main

import (
	"github/bradtumy/hippo/cmd/hippo"
	"github/bradtumy/hippo/config"

	"log"
)

func main() {
	cfg, err := config.New("config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	//hippo := hippo.New()
	//router := routes.NewRouter(hippo)
	hippo.Startup(cfg)
}
