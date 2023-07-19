package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

// http://localhost:8899/swagger/index.html
func SwaggerRouting(host string, router *mux.Router) {
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(host+"/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}
