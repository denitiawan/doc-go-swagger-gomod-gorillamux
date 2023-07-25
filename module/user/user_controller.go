package user

import (
	"context"
	baseController "denitiawan/research-swagger-gomod-gorillamux/common/controller"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserController struct {
	service UserService
	ctx     context.Context
}

func NewUserController(contx context.Context, service UserService) *UserController {
	return &UserController{
		service: service,
		ctx:     contx,
	}
}

// @Tags			user
// @Router			/v1/user/save [post]
// @Summary			save
// @Description		save
// @Param			RequestBody body user.UserDto true "UserDto.go"
// @Param			Authorization header string true "Authorization"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (c *UserController) Create(response http.ResponseWriter, request *http.Request) {
	// req body
	reqBody, _ := ioutil.ReadAll(request.Body)
	var dto UserDto
	json.Unmarshal(reqBody, &dto)

	// service
	res := c.service.Create(dto)
	if res.Error != nil {
		baseController.Error(res.Message, nil, res.ErrorMessage, response, request)
		return
	}

	// ok
	baseController.OK(res.Message, res.Data, "", response, request)
}

// @Tags			user
// @Router			/v1/user/update/{id} [put]
// @Summary			update
// @Description		update
// @Param			Authorization header string true "Authorization"
// @Param			id  path int true "id"
// @Param			RequestBody body UserDto true "UserDto.go"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (c *UserController) Update(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Update")

	// req Id
	vars := mux.Vars(request)
	param, err := strconv.Atoi(vars["Id"])
	if err != nil {
		baseController.BadRequest("Invalid request", nil, "", response, request)
		return
	}
	dtoId, _ := strconv.ParseInt(strconv.Itoa(param), 0, 64)

	// req body
	reqBody, _ := ioutil.ReadAll(request.Body)
	var dto UserDto
	json.Unmarshal(reqBody, &dto)

	// service
	res := c.service.Update(dtoId, dto)
	if res.Error != nil {
		baseController.Error(res.Message, nil, res.ErrorMessage, response, request)
		return
	}

	baseController.OK(res.Message, res.Data, "", response, request)
}

// @Tags			user
// @Router			/v1/user/delete/{id} [delete]
// @Summary			delete
// @Description		delete
// @Param			Authorization header string true "Authorization"
// @Param			id  path int true "id"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (c *UserController) Delete(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Delete")
	// req Id
	vars := mux.Vars(request)
	param, err := strconv.Atoi(vars["Id"])
	if err != nil {
		baseController.BadRequest("Invalid request", nil, "", response, request)
		return
	}
	dtoId, _ := strconv.ParseInt(strconv.Itoa(param), 0, 64)

	// service
	res := c.service.Delete(dtoId)
	if res.Error != nil {
		baseController.Error(res.Message, nil, res.ErrorMessage, response, request)
		return
	}

	baseController.OK(res.Message, res.Data, "", response, request)
}

// @Tags			user
// @Router			/v1/user/view/{id} [get]
// @Summary			view
// @Description		view
// @Param			Authorization header string true "Authorization"
// @Param			id  path int true "id"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (c *UserController) FindById(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("Update")

	// req Id
	vars := mux.Vars(request)
	param, err := strconv.Atoi(vars["Id"])
	if err != nil {
		baseController.BadRequest("Invalid request", nil, "", response, request)
		return
	}
	dtoId, _ := strconv.ParseInt(strconv.Itoa(param), 0, 64)

	// service
	res := c.service.FindById(dtoId)
	if res.Error != nil {
		baseController.Error(res.Message, nil, res.ErrorMessage, response, request)
		return
	}

	baseController.OK(res.Message, res.Data, "", response, request)

}

// @Tags			user
// @Router			/v1/user/list [get]
// @Summary			list
// @Description		list
// @Param			Authorization header string true "Authorization"
// @Produce			application/json
// @Success			200 {object} user.UserDto{} "Response Success (UserDto.go)"
func (c *UserController) FindAll(response http.ResponseWriter, request *http.Request) {
	log.Info().Msg("FindAll")

	// service
	res := c.service.FindAll()
	if res.Error != nil {
		baseController.Error(res.Message, nil, res.ErrorMessage, response, request)
		return
	}

	// ok
	baseController.OK(res.Message, res.Data, "", response, request)
}
