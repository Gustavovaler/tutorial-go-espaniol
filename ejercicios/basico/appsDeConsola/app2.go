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

func menu() {

	fmt.Println("***  Presione 1/Listado 2/Nueva mascota 3/borrar mascota*** \n")
	fmt.Scanln(&opcion)
	fmt.Println("elegiste opcion ", opcion)

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

	fmt.Println(listadoMascotas)
	menu()
}

func nueva() {

	idraw := len(listadoMascotas)
	listadoMascotas[id] = Mascota{nombre: "Firulais", especie: "perro", genero: "macho"}
	menu()

}

func eliminar() {

	fmt.Scanln(&idd)

	delete(listadoMascotas, idd)

	menu()

}

func main() {

	fmt.Println("\n  ***  Bienvenido al refugio de mascotas *** \n")

	menu()

}
