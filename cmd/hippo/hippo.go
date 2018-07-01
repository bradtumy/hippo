package hippo

import (
	"fmt"
	"github/bradtumy/hippo/config"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Hippo struct {
	Config config.Config
	//Database *database.MySQLDB
	//Redis    *database.RedisDB
}

// New is a test
func New(cfg config.Config) *Hippo {
	// test
	return &Hippo{cfg}
}

// Startup provides all of the config details and runs hippo
func (a *Hippo) Startup(r *mux.Router) {

	Port := a.Config.Port
	headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	/*
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

	*/

	//r.Handle("/validate", ValidateMiddleware(TestEndpoint)).Methods("GET")
	fmt.Println("Hippo is listening on port: ", Port)
	log.Fatal(http.ListenAndServe(Port, handlers.CORS(originsOk, headersOk, methodsOk)(r)))

}
