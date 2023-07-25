package user

import (
	"context"
	"database/sql"
	"denitiawan/research-swagger-gomod-gorillamux/config"
	"denitiawan/research-swagger-gomod-gorillamux/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func APIUser(appConfig config.AppConfig,
	ctx context.Context,
	db *sql.DB,
	router *mux.Router,
	validate *validator.Validate) {

	// base
	prefix := "/v1/user"

	// repo,service,controller
	userRepo := NewUserRepoImpl(ctx, db)
	userService := NewUserServiceImpl(userRepo, validate)
	userController := NewUserController(ctx, userService)

	// endpoint
	router.HandleFunc(prefix+"/list", middleware.Authorization(appConfig, userController.FindAll)).Methods("GET")
	router.HandleFunc(prefix+"/view/{Id}", middleware.Authorization(appConfig, userController.FindById)).Methods("GET")
	router.HandleFunc(prefix+"/save", middleware.Authorization(appConfig, userController.Create)).Methods("POST")
	router.HandleFunc(prefix+"/update/{Id}", middleware.Authorization(appConfig, userController.Update)).Methods("PUT")
	router.HandleFunc(prefix+"/delete/{Id}", middleware.Authorization(appConfig, userController.Delete)).Methods("DELETE")
}
