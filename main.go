package main

import (
	"denitiawan/research-swagger-gomod-gorillamux/common/error"
	"denitiawan/research-swagger-gomod-gorillamux/config"
	_ "denitiawan/research-swagger-gomod-gorillamux/docs"
	"denitiawan/research-swagger-gomod-gorillamux/router"
	"github.com/rs/zerolog/log"
	"net/http"
)

// --------------------------------------------------
// annotations		: Annotation used for Swagger-UI
//					and will be mapping to folder and files (./root/docs/**)
// docs import		: import ( _ "denitiawan/research-swagger-gomod-gorillamux/docs" )
//					will be used for update all values on all files inside that folder
//					when you run syntax (swag init)
// url swagger-ui 	: http://localhost:5050/nexsoft/doc/api/swagger/index.html
// --------------------------------------------------
// @version		1.1.0
// @title 		Demo Swagger-UI (GO+GORILLA MUX) for Nexsoft Project
// @description Implement swagger-ui on Go project with Gorilla Mux (web framework) + JWT Authorization
// @host 		localhost:5050
// @BasePath 	/

// ------showing authorize button (but validation jwt is not working)---------
// @Security Authorization
// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
// @schemes http
// ------showing authorize button (but validation jwt is not working)---------
func main() {

	log.Info().Msg("Try to start Application!")

	// App Config
	appConfig, err := config.LoadConfig(".", "dev")
	error.ErrorPanic(err)

	// # Database Connection
	db := config.DatabaseConnection(appConfig)

	// # Database Migration
	config.DatabaseMigration(appConfig, db)

	// # Routes Config
	appRoute := router.NewRouterInit()

	// # Swagger
	config.SwaggerRouting(appRoute)

	// # API Registration
	router.APIRouting(appConfig, db, appRoute)

	// # Http Handle
	http.Handle("/", appRoute)

	// # Server
	server := &http.Server{
		Addr:    ":" + appConfig.ServerPort,
		Handler: appRoute,
	}

	log.Info().Msg("Yea Boy!.. Application is running!")

	// # Serve
	err = server.ListenAndServe()
	error.ErrorPanic(err)

	defer db.Close()
}
