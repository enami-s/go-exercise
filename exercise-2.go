package main

import "fmt"

func zeroten() {
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}
}

func zerofifty() {
	for i := 0; i < 51; i++ {
		fmt.Println(i)
	}
}

func main() {
	go zeroten()
	zerofifty()
	fmt.Println("finish")
}
