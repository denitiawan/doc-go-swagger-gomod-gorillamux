package auth

import (
	"denitiawan/research-swagger-gomod-gorillamux/module/user"
)

type AuthRepo interface {
	Login(username string, password string) (model user.User, err error)
}
