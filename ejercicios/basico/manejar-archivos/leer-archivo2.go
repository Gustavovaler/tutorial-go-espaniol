package main

import (
		"fmt"
		"bufio" 
		"os"
)

func main(){

	// la variable archivo recibe el archivo del paquete OS
	// como devuelve 2 valores desechamos el segundo valor usando  " _ "
	archivo,_ := os.Open("./muestra.txt")

	//ahora creamos un objeto scanner
	scanner := bufio.NewScanner(archivo)

	//iteramos sobre el resultado
	for scanner.Scan() {

		//el método Text nos devuelve linea por linea
		linea := scanner.Text()
		fmt.Println(linea)

		// separamos para ver el resultado mas notoriamente
		fmt.Println("---------------------")
	
	}
	 
} arhivo.Close()