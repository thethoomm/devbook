package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri           string
	Method        string
	Function      func(http.ResponseWriter, *http.Request)
	Authenticated bool
}

// Put all routes in the router
func Config(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postRoutes...)

	for _, route := range routes {

		if route.Authenticated {
			r.HandleFunc(route.Uri,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	return r
}
