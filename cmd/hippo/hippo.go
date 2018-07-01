package hippo

import (
	"fmt"
	"github/bradtumy/hippo/config"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/handlers"
)

// New is a test
func New() string {
	// test
	a := "test"
	return a
}

func Startup(cfg config.Config) {

	Port := cfg.Port

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
	fmt.Println("API Gateway starting on port: ", Port)
	log.Fatal(http.ListenAndServe(Port, nil))

}
