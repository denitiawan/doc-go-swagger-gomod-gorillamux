package user

import "database/sql"

type User struct {
	Id       sql.NullInt64  `db:"users.id" fieldtag:"view"`
	Name     sql.NullString `db:"users.name" fieldtag:"view"`
	Username sql.NullString `db:"users.username" fieldtag:"view"`
	Password sql.NullString `db:"users.password" fieldtag:"view"`
}
