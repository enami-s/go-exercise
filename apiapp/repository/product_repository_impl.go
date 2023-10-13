package repository

import (
	"apiapp/product"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repo *ProductRepositoryImpl) GetProductDetail(productId int) (*product.Product, error) {
	url := fmt.Sprintf("https://dummyjson.com/products/%d", productId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var product product.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
