package user

import "denitiawan/research-swagger-gomod-gorillamux/common/dto"

type UserRepo interface {
	Create(model *User) *dto.ImplResponse
	Update(model User) *dto.ImplResponse
	Delete(id int64) *dto.ImplResponse
	FindById(id int64) *dto.ImplResponse
	FindAll() *dto.ImplResponse
}
