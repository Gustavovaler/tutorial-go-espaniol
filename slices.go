package main

import (
	"fmt"
)

func main() {
	//Slices son una de las funciones mas usadas en Go
	// Aunque guardan una similitud con los arrays tienen caracteristicas distintas

	//definiendo un slice

	s := make([]int, 3)

	// asignar valores a un slice

	s[0] = 1
	s[1] = 2
	s[2] = 3

	//imprimimos el slice completo
	fmt.Println(s)

	//imprimir elemento especifico de un slice
	fmt.Println(s[0])

	// la funcion len() puede ser aplicada a un slice
	fmt.Println(len(s))

	// funcion append() para agregar un registro a slice

	s = append(s, 5, 6)

	fmt.Println(s)

	//Imprimir segmento del slice

	fmt.Println(s[:3]) // imprime desde el 0 hasta el elemento en la posicion 2

	fmt.Println(s[1:3]) //Imprime desde la posicion 2 hasta la 3

	fmt.Println(s[2:]) //imprime desde a posicion 2 hasta el ultimo elemento del slice

	// definicion y aignacion simultanea de un slice

	t := []int{100, 200, 300}

	fmt.Println(t)

	//copia de un slice

	// esta es la forma incorrecta,descomentar para probar

	//x := s   //copiamos el slice s en x
	//fmt.Println(x) // Al imprimir pareceria que ssta  todo bien
	// Pero no es asi, ya que si modificamos u valer en x tambien se cambiara en s
	//x[0] = 100
	//fmt.Println(x)
	//fmt.Println(s) // como podemos ver ambos slices sufrieron el cambio

	// -- FORMA CORRECTA DE COPIAR UN SLICE

	x := make([]int, len(s))

	copy(x, s)

	x[0] = 100

	fmt.Println(x)
	fmt.Println(s)

	//slices 2d similar a los arrays aunque el largo de los slices puede variar

	ss := make([][]int, 3)

	for i := 0; i < 3; i++ {
		fmt.Println(ss[i])
	}

}
