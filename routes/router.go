package routes

import (
	"fmt"
	"github/bradtumy/hippo/cmd/hippo"
	"github/bradtumy/hippo/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO:  move Handlers to handlers package

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

// NewRouter creates a new router
func NewRouter(h *hippo.Hippo) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", middleware.Logger(HomeHandler))

	return r
}
