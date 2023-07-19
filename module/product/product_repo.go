package product

type ProductRepo interface {
	Save(model Product)
	Update(model Product)
	Delete(modelId int)
	FindById(modelId int) (model Product, err error)
	FindAll() []Product
}
