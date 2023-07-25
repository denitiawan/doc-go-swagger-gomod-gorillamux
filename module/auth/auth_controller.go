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

// @Tags			auth
// @Router			/v1/auth/login [post]
// @Summary			login
// @Description		Get JWT token for access all APIs
// @Param			RequestBody body LoginDto true "LoginDto.go"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (c *AuthController) Login(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Login")

	// req body
	reqBody, _ := ioutil.ReadAll(request.Body)
	var dto LoginDto
	json.Unmarshal(reqBody, &dto)

	// service
	res := c.service.Login(dto)
	if res.Error != nil {
		baseController.Error(res.Message, nil, res.ErrorMessage, response, request)
		return
	}

	// ok
	baseController.OK(res.Message, res.Data, "", response, request)

}
