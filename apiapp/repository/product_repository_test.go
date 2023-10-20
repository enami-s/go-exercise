package repository

import (
	"testing"
)

func TestResult型でのプロダクト詳細取得メソッドに2のプロダクトIDを付与すると2の結果が取得できる(t *testing.T) {
	//NewProductRepositoryでProductRepositoryを取得
	repo := NewProductRepository()

	//GetProductDetailで引数2を指定した時の結果を取得
	product := repo.FetchProductDetailByResult(2)

	//エラーが発生した場合はエラーを出力
	if product.Failure != nil {
		t.Fatal(product.Failure)
	}

	//結果が2であることを確認
	if product.Success.ID != 2 {
		t.Fatal("expected 2 but got", product.Success.ID)
	}
}

func Testプロダクト詳細取得メソッドに1のプロダクトIDを付与すると1の結果が取得できる(t *testing.T) {
	//NewProductRepositoryでProductRepositoryを取得
	repo := NewProductRepository()

	//GetProductDetailで引数1を指定した時の結果を取得
	product, err := repo.GetProductDetail(1)

	//エラーが発生した場合はエラーを出力
	if err != nil {
		t.Fatal(err)
	}

	//結果が1であることを確認
	if product.ID != 1 {
		t.Fatal("expected 1 but got", product.ID)
	}
}

func Testプロダクト詳細取得メソッドに2のプロダクトIDを付与すると2の結果が取得できる(t *testing.T) {
	//NewProductRepositoryでProductRepositoryを取得
	repo := NewProductRepository()

	//GetProductDetailで引数2を指定した時の結果を取得
	product, err := repo.GetProductDetail(2)

	//エラーが発生した場合はエラーを出力
	if err != nil {
		t.Fatal(err)
	}

	//結果が2であることを確認
	if product.ID != 2 {
		t.Fatal("expected 2 but got", product.ID)
	}
}

func Testプロダクト一覧取得メソッドを実行すると20件のproductを取得できる(t *testing.T) {
	//NewProductRepositoryでProductRepositoryを取得
	repo := NewProductRepository()

	//GetProductsでプロダクト一覧を取得
	products, err := repo.GetProducts()

	//エラーが発生した場合はエラーを出力
	if err != nil {
		t.Fatal(err)
	}

	//プロダクト一覧の件数が20件であることを確認
	if len(products) != 20 {
		t.Fatal("expected 20 but got", len(products))
	}
}
