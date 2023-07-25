package config

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
)

// path swagger is customable
// path (/*any) is required for load the html page own by swagger
// http://localhost:8810/nexsoft/doc/api/swagger/index.html
func SwaggerRouting(router *mux.Router) {
	prefix := "/nexsoft/doc/api"
	router.PathPrefix(prefix).Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
		//httpSwagger.DeepLinking(true),
		//httpSwagger.DocExpansion("none"),
		//httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

}
