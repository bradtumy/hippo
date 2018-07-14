package main

import (
	"fmt"
	"github/bradtumy/hippo/config"
	"github/bradtumy/hippo/handlers"
	"github/bradtumy/hippo/middleware"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Start() {

	cfg, err := config.New("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	target_url := cfg.Target.Url
	target_port := cfg.Target.Port

	target_addr := fmt.Sprintf(":%v", target_port)

	remote, err := url.Parse(target_url + target_addr)
	if err != nil {
		panic(err)
	}

	port := cfg.Port
	addr := fmt.Sprintf(":%v", port)

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", middleware.Logger(handlers.Proxyhandler(proxy)))

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	Start()
}
