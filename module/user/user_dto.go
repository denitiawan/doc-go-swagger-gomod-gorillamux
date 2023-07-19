package user

type UserDto struct {
	Id       int64  `validate:"min=1,max=10" json:"id"`
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Username string `validate:"required,min=1,max=200" json:"username"`
	Password string `validate:"required,min=1,max=200" json:"password"`
}
