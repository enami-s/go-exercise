package repository

import "apiapp/product"

type ProductRepository interface {
	GetProducts() ([]*product.Product, error)
	GetProductDetail(productId int) (*product.Product, error)
}
