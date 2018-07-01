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

func startupConfig() {

	target, err := url.Parse("http://192.168.1.12:9090")
	if err != nil {
		log.Fatal(err)
	}

	target2, err2 := url.Parse("http://192.168.1.12:9092")
	if err2 != nil {
		log.Fatal(err2)
	}

	http.Handle("/identities/", handlers.LoggingHandler(os.Stdout, httputil.NewSingleHostReverseProxy(target)))
	http.Handle("/auth", handlers.LoggingHandler(os.Stdout, httputil.NewSingleHostReverseProxy(target2)))
	http.Handle("/validate", handlers.LoggingHandler(os.Stdout, httputil.NewSingleHostReverseProxy(target2)))

	//r.Handle("/validate", ValidateMiddleware(TestEndpoint)).Methods("GET")
	fmt.Println("API Gateway starting on port :9091")
	log.Fatal(http.ListenAndServe(":9091", nil))

}

func main() {
	startupConfig()
}
