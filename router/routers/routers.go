package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

func SetupRouter(r *mux.Router) *mux.Router {
	routes := fipeRoutes
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}
	return r
}
