package product

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func APIProduct(db *sql.DB, router *mux.Router, validate *validator.Validate) {
	// base
	prefix := "/v1/product"

	// repo,service,controller
	productRepo := NewProductRepoImpl(db)
	productService := NewProductServiceImpl(productRepo, validate)
	productController := NewProductController(productService)

	// endpoint
	router.HandleFunc(prefix+"/list", productController.FindAll).Methods("GET")
	router.HandleFunc(prefix+"/view/{Id}", productController.FindById).Methods("GET")
	router.HandleFunc(prefix+"/save", productController.Create).Methods("POST")
	router.HandleFunc(prefix+"/update/{Id}", productController.Update).Methods("PUT")
	router.HandleFunc(prefix+"/delete/{Id}", productController.Delete).Methods("DELETE")
}
