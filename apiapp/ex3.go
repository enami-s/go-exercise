package main

import (
	"apiapp/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetProductDetail() (*model.Product, error) {
	resp, err := http.Get("https://dummyjson.com/products/1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var product model.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func main() {
	product, err := GetProductDetail()
	if err != nil {
		fmt.Println("Failure:", err)
		return
	}

	fmt.Printf("ID: %d, Title: %s\n", product.ID, product.Title)
}
