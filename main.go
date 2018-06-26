package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("http://192.168.1.12:9090/identities")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/identities/", http.StripPrefix("/identities/", httputil.NewSingleHostReverseProxy(target)))
	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./Documents"))))
	fmt.Println("API Gateway starting on port :9091")
	log.Fatal(http.ListenAndServe(":9091", nil))
}
