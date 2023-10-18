package main

import (
	"apiapp/product"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://dummyjson.com/productsを定数で定義
const productUrl = "https://dummyjson.com/products"

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

// "https://dummyjson.com/productsのBodyのプロダクト情報を全て取得する関数
func (repo *ProductRepositoryImpl) GetProducts() ([]*product.Product, error) {

	//定数のをLimitの分だけ取得するように変更
	url := fmt.Sprintf("%s?limit=%d", productUrl, product.Limit)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//Bodyを読み込む
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response product.ProductListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Products, nil
}

func (repo *ProductRepositoryImpl) GetProductDetail(productId int) (*product.Product, error) {
	url := fmt.Sprintf("%s/%d", productUrl, productId)
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
