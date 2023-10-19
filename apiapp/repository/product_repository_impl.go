package repository

import (
	"apiapp/product"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Chanを使用して、APIから取得したデータをchannelに送信する関数
func GetProductsAsync() <-chan []*product.Product {
	//channelを定義
	ch := make(chan []*product.Product)

	//定数のをLimitの分だけ取得するように変更
	url := fmt.Sprintf("%s?limit=%d", productUrl, product.Limit)

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			ch <- nil
			return
		}
		defer resp.Body.Close()

		//Bodyを読み込む
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ch <- nil
			return
		}

		var response product.ProductListResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			ch <- nil
			return
		}
		ch <- response.Products
	}()

	return ch
}

// https://dummyjson.com/productsを定数で定義
const productUrl = "https://dummyjson.com/products"

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

// "https://dummyjson.com/productsのBodyのプロダクト情報を20個取得する関数
func (repo *ProductRepositoryImpl) GetProducts() ([]*product.Product, error) {

	//GetProductsAsyncを呼び出し、channelを受け取る
	ch := GetProductsAsync()
	//channelからデータを受け取る
	products := <-ch
	//productsを返り値として返す
	return products, nil
}

func (repo *ProductRepositoryImpl) GetProductDetail(productId int) (*product.Product, error) {
	//フィル名を変数で定義
	var fileName = fmt.Sprintf("%d.json", productId)
	//ローカルのprivate/tmpディレクトリにfilenameと同じファイルが存在するか確認
	body, err := ioutil.ReadFile("/private/tmp/" + fileName)

	//ファイルが存在すれば、bodyをproduct.Productに変換
	if err == nil {
		var product product.Product
		err = json.Unmarshal(body, &product)
		//productを返り値として返す
		return &product, nil
	}

	//GetProductsAsyncを呼び出し、channelを受け取る
	ch := GetProductsAsync()
	//channelからデータを受け取る
	products := <-ch

	//productsの中からproductIdと一致するものを探す
	for _, product := range products {
		if product.ID == productId {
			//一致するものがあれば、productを返り値として返す
			return product, nil
		}
	}

	//一致するものがなければ、nilを返す
	return nil, nil
}

// Marshalを使ってstructからjsonに変換するメソッドを実装
func (repo *ProductRepositoryImpl) EncodeProduct(product *product.Product) ([]byte, error) {
	return json.Marshal(product)
}
