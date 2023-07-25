package router

import (
	"context"
	"database/sql"
	"denitiawan/research-swagger-gomod-gorillamux/config"
	"denitiawan/research-swagger-gomod-gorillamux/module/auth"
	"denitiawan/research-swagger-gomod-gorillamux/module/user"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func APIRouting(appConfig config.AppConfig, db *sql.DB, router *mux.Router) {

	// # context
	ctx := context.Background()

	// # Validator
	validate := validator.New()

	// # Welcome
	APIWelcome(router)

	// # API Registrations
	auth.APIAuth(appConfig, ctx, db, router, validate)
	user.APIUser(appConfig, ctx, db, router, validate)

}
