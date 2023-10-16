// mainのテストコードの記述
package main

import "testing"

var Debug bool = false

// プロダクト詳細取得に引数を指定しない場合は1の結果が取得できるテスト関数の作成
func TestIsOne(t *testing.T) {
	if Debug {
		t.Skip("スキップする")
	}

	//プロダクト詳細取得に引数を指定しない場合は1の結果が取得できるテスト関数の作成

}
