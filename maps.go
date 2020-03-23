package  main

import (
		"fmt"
)


func main(){

		//Los maps en go son como los diccionarios de Python

		// Especifican un par clave/valor

		m := make(map[string]int)

		// Agregar datos a  un mapa
		// El primer dato será la clave de la entrada y el segundo su valor

		m["a"] = 0 
		m["b"] = 1

		fmt.Println(m)

		// Imprimir un valor segun su clave

		fmt.Println(m["a"])

		// La funncion len() tambien aplica a los maps

		fmt.Println(len(m))

		// Borrar un registro según su clave

		delete(m, "a")
		fmt.Println(m)

		// Otra forma de inicializar el mapa en su declaracion

		m2 := map[string]int{"key1": 1 , "key2":2}

		fmt.Println(m2["key1"])

		// Chequear si el valer de la clave existe y manejar el error

		val, is_val_present := m["b"]

		fmt.Println(val)
		fmt.Println(is_val_present)

		//manejar el error

		_, is_val_present2 := m["a"]

		fmt.Println(is_val_present2)

		// agregar un nuevo registro

		m["c"] = 3

		fmt.Println(m)

		




}