package routes

import (
	"net/http"

	"github.com/book-wise/middlewares"
	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func ConfigUserRoutes(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, LoginRoute)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
