package main

import (
		"fmt"
		"io/ioutil"// este paquete tiene funciones especificas para amnejo de archivos
)

func main(){
	// la func ReadFile devuelve 2 valores
	// el primero son lso datos y el segundo el error si lo hubiera
	datos,err := ioutil.ReadFile("./muestra.txt")

	// en caso de presentarse un error lo atrapamos en el condicional
	if err != nil {
		fmt.Println("Hubo un error")
	}
	// para imprimir los datos debemos convertirlos a string
	fmt.Println(string(datos))	


	 
}

// Nota Importante:
// esta forma de lerr archivos no es la mas eficiente y 
// solo es recomendad para archivos pequeños ya que 
// carga el archivo completo en memoria para procesarlo
// En el siguiente ejemplo refactorizaremos con otro 
// metodo mas eficiente.