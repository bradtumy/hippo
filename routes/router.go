package routes

import (
	"encoding/json"
	"github/bradtumy/hippo/cmd/hippo"
	"github/bradtumy/hippo/handlers"
	"github/bradtumy/hippo/middleware"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

// User ...
// Custom object which can be stored in the claims
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Users(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, "decoded")
	var user User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}

// NewRouter creates a new router
func NewRouter(h *hippo.Hippo) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", middleware.Logger(handlers.HomeHandler))
	r.HandleFunc("/auth", middleware.Logger(handlers.AuthenticateHandler))
	r.HandleFunc("/users/{id}/credentials", middleware.Logger(handlers.ValidateTokenMiddleware(Users)))

	return r
}
