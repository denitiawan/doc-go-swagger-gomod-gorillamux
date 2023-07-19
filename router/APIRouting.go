package router

import (
	"database/sql"
	"denitiawan/research-swagger-gomod-gorillamux/module/auth"
	"denitiawan/research-swagger-gomod-gorillamux/module/product"
	"denitiawan/research-swagger-gomod-gorillamux/module/user"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func APIRouting(db *sql.DB, router *mux.Router) {

	// # Validator
	validate := validator.New()

	// # Welcome
	APIWelcome(router)

	// # API Registrations
	auth.APIAuth(db, router, validate)
	user.APIUser(db, router, validate)
	product.APIProduct(db, router, validate)

}
