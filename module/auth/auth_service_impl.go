package auth

import (
	"denitiawan/research-swagger-gomod-gorillamux/common/helper"
	"denitiawan/research-swagger-gomod-gorillamux/module/user"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepo AuthRepo
	Validate *validator.Validate
}

func NewAuthServiceImpl(repo AuthRepo, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepo: repo,
		Validate: validate,
	}
}

func (t *AuthServiceImpl) Login(requestDto LoginDto) user.UserDto {

	// validate
	err := t.Validate.Struct(requestDto)
	helper.ErrorPanic(err)

	// repo
	//data, err := t.AuthRepo.Login(requestDto.Username, requestDto.Password)
	//helper.ErrorPanic(err)

	//response := user.UserDto{
	//	Id:       data.Id,
	//	Name:     data.Name,
	//	Username: data.Username,
	//	Password: data.Password,
	//}
	//return response

	return user.UserDto{}
}
