package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Create a router with all routes defined
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
