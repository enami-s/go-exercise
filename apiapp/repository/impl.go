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

func (repo *ProductRepositoryImpl) Execute(productId int) {
	product, err := repo.GetProductDetail(productId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("ID: %d, Title: %s\n", product.ID, product.Title)
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
