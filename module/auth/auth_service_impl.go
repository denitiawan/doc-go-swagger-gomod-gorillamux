package auth

import (
	"denitiawan/research-swagger-gomod-gorillamux/common/dto"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImpl struct {
	AuthRepo AuthRepo
	Validate *validator.Validate
	Class    string
}

func NewAuthServiceImpl(repo AuthRepo, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepo: repo,
		Validate: validate,
		Class:    "AuthServiceImpl",
	}
}

func (t *AuthServiceImpl) Login(requestDto LoginDto) *dto.ImplResponse {

	function := "Create()"

	err := t.Validate.Struct(requestDto)
	if err != nil {
		return dto.NewImplResponse(t.Class, function, "validate dto failed", err, nil)
	}

	// repo
	res := t.AuthRepo.Login(requestDto)
	if res.Error != nil {
		return dto.NewImplResponse(t.Class, function, res.Message, res.Error, nil)
	}

	// ok
	return dto.NewImplResponse(t.Class, function, res.Message, nil, res.Data)
}
