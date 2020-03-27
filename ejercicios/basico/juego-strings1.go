package main

import (
	"fmt"
	//"strings"
)

func main() {

	solutions("soret3ksdjasdsk")

}

func solutions(str string) []string {

	result := []string{}

	if len(str)%2 != 0 {
		str = str + "_"
	}
	for i := 0; i < len(str); i += 2 {
		result = append(result, str[i:i+2])
	}
	fmt.Println(result)
	fmt.Println(str)
	return result
}
