package product

import (
	"database/sql"
)

type ProductRepoImpl struct {
	Db *sql.DB
}

func NewProductRepoImpl(Db *sql.DB) ProductRepo {
	return &ProductRepoImpl{Db: Db}
}

func (t *ProductRepoImpl) Delete(modelId int) {

}

func (t *ProductRepoImpl) FindAll() []Product {
	var model []Product
	return model
}

func (t *ProductRepoImpl) FindById(modelId int) (model Product, err error) {
	var ent Product
	return ent, nil
}

func (t *ProductRepoImpl) Save(model Product) {

}

func (t *ProductRepoImpl) Update(model Product) {

}
