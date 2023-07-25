package auth

import (
	"context"
	"database/sql"
	"denitiawan/research-swagger-gomod-gorillamux/config"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func APIAuth(appConfig config.AppConfig, ctx context.Context, db *sql.DB, router *mux.Router, validate *validator.Validate) {

	// base
	prefix := "/v1/auth"

	// repo,service,controller
	authRepo := NewAuthRepoImpl(appConfig, ctx, db)
	authService := NewAuthServiceImpl(authRepo, validate)
	authController := NewAuthController(authService)

	// endpoint
	router.HandleFunc(prefix+"/login", authController.Login).Methods("POST")
}
