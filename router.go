package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route defines elements needed to describe a route
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Route is a collection of Route items
type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/movie/amazon/{amazon_id}",
		GetMovie,
	},
}

// NewRouter creates a new router and registers all routes
// and handlers responsible for processing each of the routes
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(handler)
	}

	return router
}
