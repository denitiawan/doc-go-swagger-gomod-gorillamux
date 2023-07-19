package user

import (
	"database/sql"
)

type UserRepoImpl struct {
	Db *sql.DB
}

func NewUserRepoImpl(Db *sql.DB) UserRepo {
	return &UserRepoImpl{Db: Db}
}

func (t *UserRepoImpl) Delete(modelId int) {
}

func (t *UserRepoImpl) FindAll() []User {
	var ent []User
	return ent
}

func (t *UserRepoImpl) FindById(modelId int) (model User, err error) {
	var ent User
	return ent, nil
}

func (t *UserRepoImpl) Save(model User) {
}

func (t *UserRepoImpl) Update(model User) {
}
