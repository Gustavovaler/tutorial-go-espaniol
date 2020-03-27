package main

import (
	"fmt"
)

func main() {

	//declaramos una variable

	i := 1
	// evaluamos con switch
	switch i {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	}

	// casos con multiples posibilidades
	j := 5
	switch j {
	case 1, 2:
		fmt.Println("uno o dos")
	case 3, 4:
		fmt.Println("tres o cuatro")
	case 5:
		fmt.Println("este es un cinco")

	// en caso que no haya ninguna coincidencia
	default:
		fmt.Println("eno es ni 1 ni 2 ni 3 ni 4")
	}

	// evaluando sobre la variable con condicional

	switch {

	case j < 5:
		fmt.Println("es menor que cinco")
	case j == 5:
		fmt.Println("este es un cinco")
	case j > 5:
		fmt.Println("este es mayor a  cinco")

	}

}
