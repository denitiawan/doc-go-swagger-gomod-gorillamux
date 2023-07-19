package auth

import (
	"database/sql"
	"denitiawan/research-swagger-gomod-gorillamux/module/user"
)

type AuthRepoImpl struct {
	Db *sql.DB
}

func NewAuthRepoImpl(Db *sql.DB) AuthRepo {
	return &AuthRepoImpl{Db: Db}
}

func (t *AuthRepoImpl) Login(username string, password string) (tags user.User, err error) {
	var tag user.User

	return tag, nil
}
