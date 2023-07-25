package user

import "denitiawan/research-swagger-gomod-gorillamux/common/dto"

type UserService interface {
	Create(requestDto UserDto) *dto.ImplResponse
	Update(id int64, requestDto UserDto) *dto.ImplResponse
	Delete(id int64) *dto.ImplResponse
	FindById(id int64) *dto.ImplResponse
	FindAll() *dto.ImplResponse
}
