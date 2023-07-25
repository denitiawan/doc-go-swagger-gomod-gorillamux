package auth

import (
	"denitiawan/research-swagger-gomod-gorillamux/common/dto"
)

type AuthRepo interface {
	Login(loginDto LoginDto) *dto.ImplResponse
	FindByUsername(username string) *dto.ImplResponse
	FindById(id int64) *dto.ImplResponse
}
