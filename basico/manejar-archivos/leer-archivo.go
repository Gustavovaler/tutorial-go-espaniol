package main

import (
		"fmt"
		"io/ioutil"
)

func main(){

	datos,err := ioutil.ReadFile("./muestra.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	}

	fmt.Println(string(datos))	   
	 
}