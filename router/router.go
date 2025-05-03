package router

import (
	"tabela-fipe-golang/router/routers"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routers.SetupRouter(r)
}
