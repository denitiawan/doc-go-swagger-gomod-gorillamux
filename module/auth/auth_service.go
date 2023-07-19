package auth

import (
	"denitiawan/research-swagger-gomod-gorillamux/module/user"
)

type AuthService interface {
	Login(requestDto LoginDto) user.UserDto
}
