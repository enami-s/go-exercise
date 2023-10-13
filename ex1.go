package main

import "fmt"

func main() {
	var num int = 100
	fmt.Printf("Intは整数を扱う際に使用する。%d", num)

	var s string = "Hello World"
	fmt.Printf("stringは文字列を扱う場合に使用する。 %s\n", s)

	var f float64 = 3.14
	fmt.Printf("floatは小数を扱う際に使用する。%f\n", f)

	var u uint = 24
	fmt.Printf("Unitは正の数を扱う際に使用する。%d\n", u)

	var c complex128 = 1 + 2i
	fmt.Printf("complexは複素数を扱う場合に使用する。 %v\n", c)

	var b bool = true
	fmt.Printf("boolは真偽値を扱う場合に使用する。 %t\n", b)

	var by byte = 65
	fmt.Printf("byteは8ビット整数を扱う場合に使用する。 %d\n", by)

	var any interface{}
	any = "Go is Good!"
	fmt.Printf("interfaceは任意の型を扱う場合に使用する。 %v\n", any)

	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("配列は同じ型の値を複数扱う場合に使用する。 %v\n", arr)

}
