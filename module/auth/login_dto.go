package auth

type LoginDto struct {
	Username string `validate:"required,min=1,max=200" json:"username"`
	Password string `validate:"required,min=1,max=200" json:"password"`
}
