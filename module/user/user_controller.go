package user

import (
	baseController "denitiawan/research-swagger-gomod-gorillamux/common/controller"
	"net/http"

	"github.com/rs/zerolog/log"
)

type UserController struct {
	prodcutService UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{
		prodcutService: service,
	}
}

func (controller *UserController) Create(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Create")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Create", response, request)
}

func (controller *UserController) Update(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Update")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Update", response, request)
}

func (controller *UserController) Delete(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Delete")
	baseController.CreateWebResponse(http.StatusOK, "OK", "Delete", response, request)
}

func (controller *UserController) FindById(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Update")
	baseController.CreateWebResponse(http.StatusOK, "OK", "FindById", response, request)

}

func (controller *UserController) FindAll(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("FindAll")
	baseController.CreateWebResponse(http.StatusOK, "OK", "FindAll", response, request)
}
