package user

import (
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepo UserRepo
	Validate *validator.Validate
}

func NewUserServiceImpl(repo UserRepo, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: repo,
		Validate: validate,
	}
}

func (t *UserServiceImpl) Create(dto UserDto) {
	//err := t.Validate.Struct(dto)
	//helper.ErrorPanic(err)
	//userModel := User{
	//	//Name: dto.Name,
	//}
	//t.UserRepo.Save(userModel)
}

func (t *UserServiceImpl) Delete(dtoId int) {
	t.UserRepo.Delete(dtoId)
}

func (t *UserServiceImpl) FindAll() []UserDto {
	//result := t.UserRepo.FindAll()

	var listDto []UserDto
	//for _, value := range result {
	//	responseDto := UserDto{
	//		//Id: sql.NullInt64{Int64: value.Id},
	//	}
	//	listDto = append(listDto, responseDto)
	//}

	return listDto
}

func (t *UserServiceImpl) FindById(dtoId int) UserDto {
	//userData, err := t.UserRepo.FindById(dtoId)
	//helper.ErrorPanic(err)
	//
	//userResponse := UserDto{
	//	Id:   userData.Id,
	//	Name: userData.Name,
	//}
	//return userResponse

	return UserDto{}
}

func (t *UserServiceImpl) Update(dto UserDto) {
	//userData, err := t.UserRepo.FindById(dto.Id)
	//helper.ErrorPanic(err)
	//userData.Name = dto.Name
	//t.UserRepo.Update(userData)

}
