package user

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func APIUser(db *sql.DB, router *mux.Router, validate *validator.Validate) {

	// base
	prefix := "/v1/user"

	// repo,service,controller
	userRepo := NewUserRepoImpl(db)
	userService := NewUserServiceImpl(userRepo, validate)
	userController := NewUserController(userService)

	// endpoint
	router.HandleFunc(prefix+"/list", userController.FindAll).Methods("GET")
	router.HandleFunc(prefix+"/view/{Id}", userController.FindById).Methods("GET")
	router.HandleFunc(prefix+"/save", userController.Create).Methods("POST")
	router.HandleFunc(prefix+"/update/{Id}", userController.Update).Methods("PUT")
	router.HandleFunc(prefix+"/delete/{Id}", userController.Delete).Methods("DELETE")
}
