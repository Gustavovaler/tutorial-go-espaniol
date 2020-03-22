package main

import (
		"fmt"
)

// el loop fores el unico permitido por Go
// en Go no contamos con foreach ni do for ni while
func main(){
	
		arr := [...]int{1,2,3,4,5,6,7,8,9,10}	

		fmt.Println(arr) 

		// una manera de declarar un bucle for

		i := 0
		for i <= 10{
			fmt.Println(i)
			i++
		}

		//otra manera de lo mismo

		for x := 0; x < len(arr); x++{
			fmt.Println(arr[x])
		}

		// de forma declarativa

		for number :=  range arr{
			fmt.Println(number)
		}

		// imitando while

		x := 0
		for {
			fmt.Println(x)
			x ++

			if x == 50{
				break
			}
		}

		
}