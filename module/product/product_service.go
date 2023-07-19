package product

type ProductService interface {
	Create(dto ProductDto)
	Update(dto ProductDto)
	Delete(dtoId int)
	FindById(dtoId int) ProductDto
	FindAll() []ProductDto
}
