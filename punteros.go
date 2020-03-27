package main

import (
	"fmt"
)

type Persona struct {
	nombre string
	edad   int
}

func aumentarEdad(persona *Persona) {
	persona.edad++
}
func main() {

	p := Persona{"alexis", 34}

	aumentarEdad(&p)

	fmt.Println(p)

}
