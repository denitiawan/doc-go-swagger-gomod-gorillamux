package auth

import (
	"context"
	"database/sql"
	"denitiawan/research-swagger-gomod-gorillamux/common/database"
	baseDto "denitiawan/research-swagger-gomod-gorillamux/common/dto"
	"denitiawan/research-swagger-gomod-gorillamux/config"
	"denitiawan/research-swagger-gomod-gorillamux/module/user"
	"denitiawan/research-swagger-gomod-gorillamux/security"
)

type AuthRepoImpl struct {
	AppConfig config.AppConfig
	Context   context.Context
	Db        *sql.DB
	Class     string
}

func NewAuthRepoImpl(appConfig config.AppConfig, ctx context.Context, Db *sql.DB) AuthRepo {
	return &AuthRepoImpl{
		AppConfig: appConfig,
		Context:   ctx,
		Db:        Db,
		Class:     "AuthRepoImpl"}
}

func (t *AuthRepoImpl) Login(loginDto LoginDto) *baseDto.ImplResponse {
	function := "Login()"

	// find user
	res := t.FindByUsername(loginDto.Username)
	if res.Error != nil {
		return baseDto.NewImplResponse(t.Class, function, "username is not found", res.Error, nil)
	}

	// verify password
	find, _ := res.Data.(user.User)
	err := security.VerifyPassword(find.Password, loginDto.Password)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "wrong password", err, nil)
	}

	// Generate Token
	token, err := security.GenerateToken(
		t.AppConfig.TokenExpiresIn,
		find.Id,
		t.AppConfig.TokenSecret)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "generate token is failed", err, nil)
	}

	// response
	response := NewTokenDto("Authorization", token)
	return baseDto.NewImplResponse(t.Class, function, "OK", nil, response)

}

func (t *AuthRepoImpl) FindByUsername(username string) *baseDto.ImplResponse {
	function := "FindByUsername()"
	// tx begin
	tx, err := database.BeginTransaction(t.Context, t.Db)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, " begin tx is failed", err, nil)
	}

	// query
	sql := "select u.* from users u where u.username = ?;"
	rows, err := tx.Query(sql, username)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "query is failed", err, nil)
	}
	defer rows.Close()

	// row scan
	var model user.User
	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&model.Id, &model.Name, &model.Username, &model.Password)
			if err != nil {
				return baseDto.NewImplResponse(t.Class, function, "data not found", err, nil)
			}
		}
	}

	// response
	return baseDto.NewImplResponse(t.Class, function, "OK", nil, model)
}

func (t *AuthRepoImpl) FindById(id int64) *baseDto.ImplResponse {
	function := "FindById()"

	// tx begin
	tx, err := database.BeginTransaction(t.Context, t.Db)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "begin tx failed", err, nil)
	}

	// query
	sql := "select u.* from users u where u.id = ?;"
	rows, err := tx.Query(sql, id)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "query failed", err, nil)
	}
	defer rows.Close()

	// row scan
	var model user.User
	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&model.Id, &model.Name, &model.Username, &model.Password)
			if err != nil {
				return baseDto.NewImplResponse(t.Class, function, "Data Not Found", err, nil)
			}
		}
	}

	// response
	return baseDto.NewImplResponse(t.Class, function, "OK", nil, model)
}
