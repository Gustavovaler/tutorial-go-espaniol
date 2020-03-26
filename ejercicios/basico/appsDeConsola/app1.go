package main

import (
	"fmt"
)

// En este sencillo ejercicio debemos ingresar el nombre, mail y telefono de
// una cantidad especificada por nosotros de empleados y luego imprimirlos

// Primero creamos un typo estructura de empleado

type Empleado struct {
	nombre   string
	email    string
	telefono int
}

func main() {

	// creamos una variable que almacene nuestra eleccion de cantidad de ingresos
	var cantidadEmpleados int

	// Ingresamos por consola
	fmt.Println("Cuantos registros desea cargar ? :")

	fmt.Scanln(&cantidadEmpleados)

	// creamos un slice  con longitud 0 que será el contenedor de todos nuestros registros
	empleados := make([]Empleado, 0)

	//creamos una variable que será la que reciba los datos que vamos ingresando
	var empleado Empleado

	// Con el bucle for agregamos todos los datos al slice empleados

	for i := 0; i < cantidadEmpleados; i++ {

		fmt.Println("Ingrese el nombre :")

		fmt.Scanln(&empleado.nombre) // asignamos la entrada del teclado a nuestro objeto empleado

		fmt.Println("Ingrese el email :")

		fmt.Scanln(&empleado.email) // asignamos la entrada del teclado a nuestro objeto empleado

		fmt.Println("Ingrese el telefono :")

		fmt.Scanln(&empleado.telefono) // asignamos la entrada del teclado a nuestro objeto empleado

		fmt.Println("-----------------------------")

		empleados = append(empleados, empleado) // y agregamos el registro al slice empleados

		// note qeu en el método append siempre el primer parámetro es el slice al cual se desea agregar

	}
	// Por ultimos imprimimos por consola la lista completa
	fmt.Println(empleados)

}
