package repository

type ProductRepository interface {
	Execute(productId int)
}
