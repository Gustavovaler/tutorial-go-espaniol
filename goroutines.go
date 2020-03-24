package main

import (
		"fmt"
		"time"
		"strings"
)

func main(){

	go mi_nombre_lento("Gusta")
	fmt.Println("que lento este metodo")
	var wait string
	fmt.Scanln(&wait)
}


func mi_nombre_lento(nombre string){
	letras := strings.Split(nombre, "")
	


for _,letra := range(letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
}
	 
}