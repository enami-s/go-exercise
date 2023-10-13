package repository

import "apiapp/product"

type ProductRepository interface {
	GetProductDetail(productId int) (*product.Product, error)
}
