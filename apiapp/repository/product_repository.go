package repository

import "apiapp/product"

type ProductRepository interface {
	//引数にLimitも指定できるようにGetProductを変更
	GetProducts() ([]*product.Product, error)
	GetProductDetail(productId int) (*product.Product, error)
	EncodeProduct(product *product.Product) ([]byte, error)
}
