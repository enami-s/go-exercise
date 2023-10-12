package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type ProductRepository interface {
	GetProductDetail(productId int) (*Product, error)
}

type ProductRepositoryImpl struct{}

func (repo *ProductRepositoryImpl) GetProductDetail(productId int) (*Product, error) {
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

	var product Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func main() {
	repo := &ProductRepositoryImpl{}
	product, err := repo.GetProductDetail(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("ID: %d, Title: %s\n", product.ID, product.Title)
}
