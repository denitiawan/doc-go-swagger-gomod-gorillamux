package user

type UserService interface {
	Create(dto UserDto)
	Update(dto UserDto)
	Delete(dtoId int)
	FindById(dtoId int) UserDto
	FindAll() []UserDto
}
