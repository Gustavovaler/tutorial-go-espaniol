package main

import (
	"fmt"
	"time"
)

func recibir(c chan string) {
	var contador int
	for {
		contador++

		fmt.Println(<-c, " ", contador)
		time.Sleep(time.Second * 1)
	}
}

func enviar(c chan string) {
	for {
		c <- "ping"
	}
}

func main() {

	ch := make(chan string)

	go recibir(ch)
	go enviar(ch)

	var input string

	fmt.Scanln(&input)

	fmt.Println("Fin del programa..")
}
