package main

//En este archivo muestro una forma basica de hacer un servidor http con go
// En esta configuracion perfectamente puede usarse com servidor local de desarrollo


// Se importan los paquetes necesarios
import (
		"fmt"
		"net/http"		
		"time"
)



func main(){
		
		// Creamos un objeto NewServerMux para crear nuestr Handler

		mux := http.NewServeMux()

		// creamos un objeto Fileserver para servir los archivos estaticos
		// como parametro tomamos el metodo Dir le pasamos la ruta a la carpeta donde encontrar 
		//los arhivos estaticos

		fs := http.FileServer(http.Dir("public"))

		// creamos nuestro Handler, este nos servira de enrutador

		mux.Handle("/", fs )

		// creamos el objeto http.Server
		// contiene la configuracion de puerto y host

		server := &http.Server{
				//Addr recibe string con host y puerto
				// en este ejemplo ":8080" es =  a "localhost:8080"
				Addr:  ":8080",
				//le pasamos nuestro handler
				Handler: mux ,
				// Algunas conf de seguridad 
				ReadTimeout: 10*time.Second,
				WriteTimeout: 10*time.Second,
				MaxHeaderBytes: 1 << 20,
		}
		//mensaje para saber si esta funcionando
		fmt.Println("Servidor escuchando el puerto 8080")
		// Iniciamos el servidor llamando al metodo ListAndServe
		server.ListenAndServe()

		// Para detener el servidor se usa Cntrl + C 
}