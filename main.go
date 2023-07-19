package main

import (
	"denitiawan/research-swagger-gomod-gorillamux/common/helper"
	"denitiawan/research-swagger-gomod-gorillamux/config"
	"denitiawan/research-swagger-gomod-gorillamux/router"
	"github.com/rs/zerolog/log"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	log.Info().Msg("Started Server!")

	// # Database Connection
	db := config.DatabaseConnection()

	// # Database Migration
	//config.DatabaseMigration(db)

	// # Routes Config
	appRoute := router.NewRouterInit()

	// # Swagger
	router.SwaggerRouting("http://localhost:8899", appRoute)

	// # API Registration
	router.APIRouting(db, appRoute)

	// # Http Handle
	http.Handle("/", appRoute)

	// # Server
	server := &http.Server{
		Addr:    ":8899",
		Handler: appRoute,
	}

	// # Serve
	err := server.ListenAndServe()
	helper.ErrorPanic(err)

	defer db.Close()
}
