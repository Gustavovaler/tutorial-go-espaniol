package main

import (
	"fmt"
)

func main() {

	// condiciona sencillo

	// declaramos una variable y le asignamos el valor 5
	i := 5

	// Evaluamos si el valor de i  al dividirlo en 2 es par o no
	if i%2 == 0 {
		fmt.Println(i, "Es par")
	} else {
		fmt.Println(i, "Es impar")
	}

	k := 100

	if k == 100 {
		fmt.Println(k, "es 100")
	}

	// Ejemplo emulando un Switch de otros lenguajes

	j := 25

	if j < 50 {
		fmt.Println(j, "es menor a 50")
	} else if j > 50 {
		fmt.Println(j, "es mayor a 50")

	} else {
		fmt.Println(j, "es  50")
	}

}
