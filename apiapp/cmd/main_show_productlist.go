package main

//repositoryディレクトリをインポート
import (
	"apiapp/product"
	"apiapp/repository"
	"fmt"
)

func execute(repo repository.ProductRepository) {
	// 商品IDを指定
	productId := 1

	product, err := repo.GetProductDetail(productId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("商品詳細\nID: %d, Title: %s\n", product.ID, product.Title)
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

	//Limitの定義、デフォルトは20
	limit := 20

	products, err := repo.GetProducts(limit)
	showProducts(products, err)

	execute(repo)
}
