package main

import (
	"fmt"
	"sync"
)

const (
	zerotenmax = 11
	zerofifmax = 51
)

var wg sync.WaitGroup

func zeroten() {
	defer wg.Done()
	for i := 0; i < zerotenmax; i++ {
		fmt.Println("zeroten: ", i)
	}
}

func zerofifty() {
	defer wg.Done()
	for i := 0; i < zerofifmax; i++ {
		fmt.Println("zerofif: ", i)
	}
}

func main() {
	wg.Add(2)
	go zeroten()
	go zerofifty()

	wg.Wait()
	fmt.Println("finish")
}
