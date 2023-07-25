package user

import (
	"context"
	"database/sql"
	"denitiawan/research-swagger-gomod-gorillamux/common/database"
	"denitiawan/research-swagger-gomod-gorillamux/common/dto"
	baseDto "denitiawan/research-swagger-gomod-gorillamux/common/dto"
)

type UserRepoImpl struct {
	Context context.Context
	Db      *sql.DB
	Class   string
}

func NewUserRepoImpl(ctx context.Context, db *sql.DB) UserRepo {
	return &UserRepoImpl{
		Context: ctx,
		Db:      db,
		Class:   "UserRepoImpl"}
}

func (t *UserRepoImpl) Create(model *User) *dto.ImplResponse {
	function := "Save()"

	// begin
	tx, err := database.BeginTransaction(t.Context, t.Db)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "begin tx failed", err, nil)
	}

	// rollback
	defer tx.Rollback()

	// insert
	result, err := tx.ExecContext(t.Context,
		"INSERT INTO users (name,username,password) VALUES (?, ?, ?)",
		model.Name, model.Username, model.Password)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "Save Failed", err, nil)
	}

	// get id
	id, err := result.LastInsertId()
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "get id failed", err, nil)
	}

	// commit
	err = tx.Commit()
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "commit failed", err, nil)
	}

	// response
	model.Id = id

	// ok
	return baseDto.NewImplResponse(t.Class, function, "OK", nil, model)

}

func (t *UserRepoImpl) Update(model User) *dto.ImplResponse {
	function := "Update()"

	// begin
	tx, err := database.BeginTransaction(t.Context, t.Db)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "begin tx failed", err, nil)
	}

	// rollback
	defer tx.Rollback()

	// insert
	result, err := tx.Exec("update users set name = ?,username = ?,password = ? where id  = ?;",
		model.Name, model.Username, model.Password, model.Id)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "Update Failed", err, nil)
	}

	result.RowsAffected()

	// commit
	err = tx.Commit()
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "commit failed", err, nil)
	}

	// ok
	return baseDto.NewImplResponse(t.Class, function, "OK", nil, model)
}

func (t *UserRepoImpl) Delete(id int64) *dto.ImplResponse {
	function := "Delete()"

	// begin
	tx, err := database.BeginTransaction(t.Context, t.Db)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "begin tx failed", err, nil)
	}

	// rollback
	defer tx.Rollback()

	// query
	result, err := tx.Exec("delete from users where id = ?;", id)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "Delete Failed", err, id)
	}
	result.RowsAffected()

	// commit
	err = tx.Commit()
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "commit failed", err, nil)
	}

	return baseDto.NewImplResponse(t.Class, function, "OK", err, id)
}

func (t *UserRepoImpl) FindById(id int64) *dto.ImplResponse {
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
	var model User
	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&model.Id, &model.Name, &model.Username, &model.Password)
			if err != nil {
				return baseDto.NewImplResponse(t.Class, function, "rows scan failed", err, nil)
			}
		}
	}

	// response
	return baseDto.NewImplResponse(t.Class, function, "OK", nil, model)
}

func (t *UserRepoImpl) FindAll() *dto.ImplResponse {
	function := "FindAll()"

	// tx begin
	tx, err := database.BeginTransaction(t.Context, t.Db)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "begin tx failed", err, nil)
	}

	// query
	sql := "select u.* from users u; "
	rows, err := tx.Query(sql)
	if err != nil {
		return baseDto.NewImplResponse(t.Class, function, "query failed", err, nil)
	}
	defer rows.Close()

	// row scan
	var models []User
	if rows != nil {
		for rows.Next() {
			var tmp User
			err = rows.Scan(&tmp.Id, &tmp.Name, &tmp.Username, &tmp.Password)
			if err != nil {
				return baseDto.NewImplResponse(t.Class, function, "Data not found", err, nil)
			}
			models = append(models, tmp)
		}
	}

	// response
	return baseDto.NewImplResponse(t.Class, function, "OK", nil, models)
}
