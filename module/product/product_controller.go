package product

import (
	baseController "denitiawan/research-swagger-gomod-gorillamux/common/controller"
	"net/http"

	"github.com/rs/zerolog/log"
)

type ProductController struct {
	prodcutService ProductService
}

func NewProductController(service ProductService) *ProductController {
	return &ProductController{
		prodcutService: service,
	}
}

func (controller *ProductController) Create(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Create")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Create", response, request)
}

func (controller *ProductController) Update(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Update")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Update", response, request)
}

func (controller *ProductController) Delete(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Delete")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Delete", response, request)
}

func (controller *ProductController) FindById(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("FindById")
	baseController.CreateWebResponse(http.StatusOK, "OK", "FindById", response, request)
}

func (controller *ProductController) FindAll(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("findAll")
	baseController.CreateWebResponse(http.StatusOK, "OK", "FindAll", response, request)
}
