package repository

import (
	"apiapp/model"
)

type ProductRepository interface {
	//引数にLimitも指定できるようにGetProductを変更
	GetProducts() ([]*model.Product, error)
	GetProductDetail(productId int) (*model.Product, error)
	EncodeProduct(product *model.Product) ([]byte, error)
	FetchProductDetailByResult(productId int) *model.Result[model.Product]
}
