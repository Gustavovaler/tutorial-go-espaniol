package main

import "fmt"

type Ojos struct{
		color string
		cantidad int
}

type datos struct{
		edad int 
		nombre string
		ojo Ojos
}
	
func (d Ojos) sumar_ojos() string{
	return d.nombre + " " + d.ojo.color
}

func main(){

		ojo1 := new(Ojos)
		var ojo2 Ojos
		ojo2.cantidad = 6
		ojo2.color = "blanco"

		ojo1.color = "negros"
		ojo1.cantidad = 2

		usuario := datos{edad:5 , nombre: "gustavo",ojo:*ojo1 }
		fmt.Println(usuario)
		fmt.Println(*ojo1.ojo.sumar_ojos)

		fmt.Println(&ojo2)


}