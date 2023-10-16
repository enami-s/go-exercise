// mainのテストコードの記述
package main

//GetProductDetailをインポート
import (
	"apiapp/repository"
	"fmt"
	"testing"
)

var Debug bool = false

// プロダクト詳細取得に引数を指定しない場合は1の結果が取得できるテスト関数の作成
func TestGetProductsDetail(t *testing.T) {
	if Debug {
		t.Skip("スキップする")
	}

	printedOutput := ""
	capturePrint = func(a ...interface{}) (n int, err error) {
		printedOutput += fmt.Sprint(a...)
		return len(printedOutput), nil
	}

	//repositoryをインポート
	repo := repository.NewProductRepository()

	//executeを実行
	execute(repo)

	//出力された内容が正しいか確認
	if printedOutput != "商品詳細\nID: 1, Title: iPhone 9\n" {
		t.Errorf("正しい結果が出力されませんでした: %s", printedOutput)
	}

}
