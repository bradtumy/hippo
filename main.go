package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	target, err := url.Parse("http://192.168.1.12:9090")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/identities/", handlers.LoggingHandler(os.Stdout, httputil.NewSingleHostReverseProxy(target)))
	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./Documents"))))
	fmt.Println("API Gateway starting on port :9091")
	log.Fatal(http.ListenAndServe(":9091", nil))
}
