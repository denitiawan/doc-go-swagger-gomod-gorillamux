package user

type User struct {
	Id       int64  `db:"users.id" fieldtag:"view"`
	Name     string `db:"users.name" fieldtag:"view"`
	Username string `db:"users.username" fieldtag:"view"`
	Password string `db:"users.password" fieldtag:"view"`
}

func NewUser(id int64, name string, username string, password string) *User {
	return &User{
		Id:       id,
		Name:     name,
		Username: username,
		Password: password,
	}
}
