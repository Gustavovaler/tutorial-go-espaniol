
package main

import (
	"fmt"
	//"strconv"
)

func main() {

		//Declarar un array de 3 lugares
		var x [3]int
		fmt.Println(x)

		//Declarar matriz 2 dimensiones
		var k[3][2]float64
		fmt.Println(k)

		//inicializar un valor en un array
		x[1] = 25
		fmt.Println(x)

		//Declarar e inicializar en el mismo paso
		y := [5]int{1,2,3,4,5}
		fmt.Println(y)

		//Declarar e inicializar otra version sin indicar la cantidad de valores
		j := [...]int{9,8,7,6}
		fmt.Println(j)

		//otra forma de lo mismo de arriba
		i := [...]int{
			1,
			2,
			3,
			4,
		}
		fmt.Println(i)

		//arrays de cualquier tipo

		f := [...]float64{1.5, 2.6, 8.3}
		fmt.Println(f)

		//Metodos de ARRAY

		varilen := len(f)
		fmt.Println(varilen)

		fmt.Println(f[len(f)-1])




}

