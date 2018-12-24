package main

import (
	"github.com/gorilla/mux"
)

// NewRouter returns a new router with routes configuration
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Methods...).
			Path(route.Path).
			HandlerFunc(route.Handler)
	}

	return router
}
