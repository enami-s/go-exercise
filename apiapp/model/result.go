// パッケージの定義
package model

// Genericsを使用したResulttypeの構造体を定義
type Result[T any] struct {
	Success *T
	Failure error
}

// APIによるデータの取得が失敗したらエラーの結果を作成する
func NewErrorResult[T any](err error) *Result[T] {
	return &Result[T]{Success: nil, Failure: err}
}

// APIによるデータの取得が成功したら成功の結果を作成する
func NewSuccessResult[T any](success *T) *Result[T] {
	return &Result[T]{Success: success, Failure: nil}
}
