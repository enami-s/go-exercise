package repository

import (
	"apiapp/product"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// publicでlimitの20を定義することで、外部から参照できるようになる
const Limit = 20

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

// "https://dummyjson.com/productsのBodyのプロダクト情報を全て取得する関数
func (repo *ProductRepositoryImpl) GetProducts() ([]*product.Product, error) {
	//limitを20に設定
	url := fmt.Sprintf("https://dummyjson.com/products?limit=%d", Limit)

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

	var response product.ProductsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Products, nil
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
