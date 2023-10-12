package main

//repositoryディレクトリをインポート
import (
	"apiapp/repository"
)

func main() {
	repo := repository.NewProductRepository()
	repo.Execute(1)
}
