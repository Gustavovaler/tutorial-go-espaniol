package main

import (
	"fmt"
	"os"
)

// ALGUNAS FUNCIONES UTILES PARA MANEJAR CARPETAS Y ARCHIVOS  ///

func main() {

	// devuelve el valor de una clave del path si es que existe
	fmt.Println(os.Getenv("USERDOMAIN"))

	//devuelve el directorio actual
	dir, _ := os.Getwd()
	fmt.Println(dir)

	//devuelve el hostname
	fmt.Println(os.Hostname())

	// cRea un directorio
	os.Mkdir("Nueva Carpeta", os.ModeDir)

	//borra un archivo o directorio vacio
	os.Remove("Nueva carpeta")

	//Borrado recursivo !!WARNING!!
	os.RemoveAll("Nueva Carpeta")

	// Mueve un archivo desde una ruta hasta otra
	// Si el archivo existe y NO es un directorio entonces lo reemplaza
	err := os.Rename(dir+"\\archivo.txt", dir+"\\ejercicios\\archivo.txt")
	fmt.Println(err)

	//devuelve el directorio temporal por defecto
	fmt.Println(os.TempDir())

	// devuelve el directorio Home
	fmt.Println(os.UserHomeDir())

}
