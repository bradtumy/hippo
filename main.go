package main

import (
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
	log.Fatal(http.ListenAndServe(":9091", nil))
}
