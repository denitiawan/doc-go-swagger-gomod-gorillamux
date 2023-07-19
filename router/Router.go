package router

import (
	"github.com/gorilla/mux"
)

func NewRouterInit() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}
