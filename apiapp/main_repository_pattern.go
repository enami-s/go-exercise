package main

//repositoryディレクトリをインポート
import (
	"apiapp/repository"
	"fmt"
)

func execute(repo repository.ProductRepository, productId int) {
	product, err := repo.GetProductDetail(productId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("ID: %d, Title: %s\n", product.ID, product.Title)
}

func main() {
	repo := repository.NewProductRepository()
	execute(repo, 1)
}
