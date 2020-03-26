package main

import (
	"fmt"
)

////////// INCOMPLETO ///////////////////////////////
////////// INCOMPLETO ///////////////////////////////
////////// INCOMPLETO ///////////////////////////////
////////// INCOMPLETO ///////////////////////////////
////////// INCOMPLETO ///////////////////////////////
////////// INCOMPLETO ///////////////////////////////
////////// INCOMPLETO ///////////////////////////////
// creamos una estructura type mascota
type Mascota struct {
	nombre  string
	especie string
	genero  string
}

// creamos un map que será el contenedor de nuestros registros
var listadoMascotas = make(map[int]Mascota)
var opcion int
var idd int
var id int
var nom, gen, esp string

func menu() {
	fmt.Println("***  Presione 1/Listado 2/Nueva mascota 3/borrar mascota*** \n")
	fmt.Scanln(&opcion)

	switch opcion {
	case 1:
		listar()

	case 2:
		nueva()

	case 3:
		eliminar()
	}
}

func listar() {

	for x := range listadoMascotas {
		fmt.Println(x, listadoMascotas[x].nombre, listadoMascotas[x].especie, listadoMascotas[x].genero)
	}

	menu()
}

func nueva() {
	id++
	fmt.Println("Ingrese Especie")
	fmt.Scanln(&esp)
	fmt.Println("Ingrese Genero")
	fmt.Scanln(&gen)
	fmt.Println("Ingrese Nombre")
	fmt.Scanln(&nom)

	listadoMascotas[id] = Mascota{nombre: nom, especie: esp, genero: gen}
	menu()

}

func eliminar() {

	fmt.Println("Seleccione el ID de la mascota a borrar de la lista:")

	fmt.Scanln(&idd)

	delete(listadoMascotas, idd)

	menu()

}

func main() {

	fmt.Println("\n  ***  Bienvenido al refugio de mascotas *** \n")

	menu()

}
