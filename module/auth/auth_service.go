package auth

import (
	"denitiawan/research-swagger-gomod-gorillamux/common/dto"
)

type AuthService interface {
	Login(requestDto LoginDto) *dto.ImplResponse
}
