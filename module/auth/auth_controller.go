package auth

import (
	baseController "denitiawan/research-swagger-gomod-gorillamux/common/controller"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

type AuthController struct {
	service AuthService
}

func NewAuthController(service AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (controller *AuthController) Login(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Login")

	reqBody, _ := ioutil.ReadAll(request.Body)

	var dto LoginDto
	json.Unmarshal(reqBody, &dto)
	data := controller.service.Login(dto)

	baseController.CreateWebResponse(http.StatusOK, "OK", data, response, request)
}
