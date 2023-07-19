package auth

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func APIAuth(db *sql.DB, router *mux.Router, validate *validator.Validate) {
	// base
	prefix := "/v1/auth"

	// repo,service,controller
	authRepo := NewAuthRepoImpl(db)
	authService := NewAuthServiceImpl(authRepo, validate)
	authController := NewAuthController(authService)

	// endpoint
	router.HandleFunc(prefix+"/login", authController.Login).Methods("POST")
}
