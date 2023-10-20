package repository

import (
	"apiapp/product"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// /private/tmp/のproductId.jsonを取得する関数
func GetfileProduct(productId int) (*product.Product, error) {
	//channelを定義
	ch := make(chan *product.Product)
	errCh := make(chan error)

	//読み込んで変換している処理をgo funcにして非同期処理にする
	go func() {
		defer close(ch)
		defer close(errCh)
		//フィル名を変数で定義
		var fileName = fmt.Sprintf("%d.json", productId)
		//ローカルのprivate/tmpディレクトリにfilenameと同じファイルが存在するか確認
		body, err := ioutil.ReadFile("/private/tmp/" + fileName)
		//ファイルが存在しなければ、エラーを返す
		if err != nil {
			errCh <- err
			return
		}
		var product *product.Product
		err = json.Unmarshal(body, &product)
		if err != nil {
			errCh <- err
			return
		}
		//チャネルに格納
		ch <- product
	}()

	select {
	case product := <-ch:
		return product, nil
	case err := <-errCh:
		return nil, err
	}
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

// https://dummyjson.com/productsを定数で定義
const productUrl = "https://dummyjson.com/products"

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repo *ProductRepositoryImpl) GetProductDetail(productId int) (*product.Product, error) {
	var product *product.Product
	//A(productId)の内容があれば、A(productId)を返す
	product, err := GetfileProduct(productId)
	if err == nil {
		return product, nil
	}

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

	err = json.Unmarshal(body, &product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// Marshalを使ってstructからjsonに変換するメソッドを実装
func (repo *ProductRepositoryImpl) EncodeProduct(product *product.Product) ([]byte, error) {
	return json.Marshal(product)
}
