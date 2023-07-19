package user

type UserRepo interface {
	Save(model User)
	Update(model User)
	Delete(modelId int)
	FindById(modelId int) (model User, err error)
	FindAll() []User
}
