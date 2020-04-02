package main

import (
	"fmt"       //paquete para imprimir en consola
	"math/rand" // paquete con funciones matemaicas
	"sync"      // paquete para manejar sync de rutinas
	"time"      // paquete control de timpo
)

var wg sync.WaitGroup // definimos una variable del tipo WaitGroup

func main() {

	// Le asignamos con el método Add cuantas rutinas tiene que esperar el programa principal
	wg.Add(2)
	// iniciamos la ejecucion
	fmt.Println("Iniciando el programa...")
	// asignamos el llamado a la funcion imprimir a una goroutine
	go imprimir("A")
	// asignamos el llamado a la funcion imprimir a una goroutine
	go imprimir("B")

	fmt.Println("Esperando que finalicen las rutinas")
	// el metodo Wait detiene la ejecucion en main() hasta que terminen las rutinas que
	// le asignamos mediante Add()   (en este caso 2)
	//el metodo Wait espera hasta que Add devuelva 0
	wg.Wait()

	fmt.Println("Terminando el programa")

}

// definimos una funcion que imprime 10 veces el string recibida como param
func imprimir(l string) {
	// wg.Done cada vez que  es llamado decrementa en 1 el valor de Add()
	defer wg.Done()
	letra := l
	for i := 1; i < 10; i++ {
		fmt.Println("Imprimimos la letra: ", letra)
		//ralentizamos un poco la ejecucion para poder ver el efecto
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

	}
}
