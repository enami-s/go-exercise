package main

//repositoryディレクトリをインポート
import (
	"apiapp/model"
	"apiapp/repository"
	"fmt"
)

func execute(repo repository.ProductRepository) {
	// 商品IDを指定
	productId := 2

	product := repo.FetchProductDetailByResult(productId)

	// 商品の情報を表示
	fmt.Println("## 商品詳細\nID: ", product.Success.ID, ", Title: ", product.Success.Title, "\n")
	fmt.Println("## encode")
	encoded, err := repo.EncodeProduct(product.Success)

	if err != nil {
		fmt.Println("Failure:", err)
		return
	}

	fmt.Println(string(encoded) + "\n")
}

// 商品の一覧を表示する関数
func showProducts(products []*model.Product, err error) {
	if err != nil {
		fmt.Println("Failure fetching products:", err)
		return
	}

	fmt.Println("## decode\n### 商品一覧")
	// 取得したプロダクトの情報を表示
	for _, product := range products {
		fmt.Printf("Title: %s, ID; %d\n", product.Title, product.ID)
	}

}

func main() {

	repo := repository.NewProductRepository()

	// 商品一覧を取得
	products, err := repo.GetProducts()
	showProducts(products, err)

	// 商品詳細を取得
	execute(repo)
}
