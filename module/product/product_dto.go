package product

type ProductDto struct {
	Id   int    `validate:"min=1,max=10" json:"id"`
	Name string `validate:"required,min=1,max=200" json:"name"`
}
