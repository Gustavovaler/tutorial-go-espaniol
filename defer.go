package main

import (
	"fmt"
)

func main() {

	op := sumar(5, 7)

	fmt.Println(op)

}

func sumar(a int, b int) int {
	defer func() {
		fmt.Println("este es defer")
	}()

	if true {
		div := a / 1
		fmt.Println(div)
		return 1
	}

	return a + b
}
