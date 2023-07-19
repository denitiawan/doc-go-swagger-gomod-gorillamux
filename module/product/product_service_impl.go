package product

import (
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepo ProductRepo
	Validate    *validator.Validate
}

func NewProductServiceImpl(repo ProductRepo, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepo: repo,
		Validate:    validate,
	}
}

func (t *ProductServiceImpl) Create(dto ProductDto) {
	//err := t.Validate.Struct(dto)
	//helper.ErrorPanic(err)
	//productModel := Product{
	//	Name: dto.Name,
	//}
	//t.ProductRepo.Save(productModel)
}

func (t *ProductServiceImpl) Delete(dtoId int) {
	t.ProductRepo.Delete(dtoId)
}

func (t *ProductServiceImpl) FindAll() []ProductDto {
	//result := t.ProductRepo.FindAll()
	//
	var listDto []ProductDto
	//for _, value := range result {
	//	responseDto := ProductDto{
	//		Id:   value.Id,
	//		Name: value.Name,
	//	}
	//	listDto = append(listDto, responseDto)
	//}
	//
	return listDto

}

func (t *ProductServiceImpl) FindById(dtoId int) ProductDto {
	//productData, err := t.ProductRepo.FindById(dtoId)
	//helper.ErrorPanic(err)
	//
	//productResponse := ProductDto{
	//	Id:   productData.Id,
	//	Name: productData.Name,
	//}
	//return productResponse
	return ProductDto{}
}

func (t *ProductServiceImpl) Update(dto ProductDto) {
	//productData, err := t.ProductRepo.FindById(dto.Id)
	//helper.ErrorPanic(err)
	//productData.Name = dto.Name
	//t.ProductRepo.Update(productData)
}
