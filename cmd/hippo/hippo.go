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

	port := a.Config.Port
	addr := fmt.Sprintf(":%v", port)

	headersOk := handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	fmt.Println("API Gateway is started and listening on port:", port)
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(r)))

}
