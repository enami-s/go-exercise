package main

//repositoryディレクトリをインポート
import (
	"apiapp/product"
	"apiapp/repository"
	"fmt"
)

var capturePrint func(a ...interface{}) (n int, err error) = fmt.Print

func execute(repo repository.ProductRepository) {
	// 商品IDを指定
	productId := 1

	product, err := repo.GetProductDetail(productId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	capturePrint("商品詳細\nID: ", product.ID, ", Title: ", product.Title, "\n")
}

// 商品の一覧を表示する関数
func showProducts(products []*product.Product, err error) {
	if err != nil {
		fmt.Println("Error fetching products:", err)
		return
	}

	fmt.Println("商品一覧")
	// 取得したプロダクトの情報を表示
	for _, product := range products {
		fmt.Printf("Title: %s, ID; %d\n", product.Title, product.ID)
	}
}

func main() {

	repo := repository.NewProductRepository()

	products, err := repo.GetProducts()
	showProducts(products, err)

	execute(repo)
}
