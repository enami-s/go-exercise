package main

//repositoryディレクトリをインポート
import (
	"apiapp/product"
	"apiapp/repository"
	"fmt"
)

func Execute(repo repository.ProductRepository) {
	// 商品IDを指定
	productId := 1

	product, err := repo.GetProductDetail(productId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("商品詳細\nID: %d, Title: %s\n", product.ID, product.Title)
}

// 商品の一覧を取得する関数
func List(products []*product.Product, err error) {
	if err != nil {
		fmt.Println("Error fetching products:", err)
		return
	}

	fmt.Println("商品一覧")
	// 取得したプロダクトの情報を表示
	for _, product := range products {
		fmt.Printf("Title: %s\n", product.Title)
	}
}

func main() {

	repo := repository.NewProductRepository()

	products, err := repo.GetProducts()
	List(products, err)

	Execute(repo)
}
