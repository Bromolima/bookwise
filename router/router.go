package router

import (
	"github.com/book-wise/router/routes"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigUserRoutes(r)
}
