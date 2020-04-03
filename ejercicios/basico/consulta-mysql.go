package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// Creamos una conexion a base de datos (aunque no la abre)
	db, error := sql.Open("mysql", "root:@/base1")
	if error != nil {
		log.Fatal(error)
	}
	// Las variables que van a recibir los datos
	var (
		matricula string
		nombre    string
	)

	// Realizamos la consulta sql

	filas, err := db.Query("SELECT matricula, nombre FROM alumnos")
	if err != nil {
		log.Fatal(err)
	}
	// SIEMPRE DEFER  consulta.Close()
	// También es buena idea cerrar la cosulta luego del bucle que recupera lso datos explicitamente
	defer filas.Close()

	// tambien podemos crear un slice para guardar los datos. o un map

	resultado := []string{}

	// Uteramos sobre el resultado
	for filas.Next() {
		// Capturamos los datos con el Método Scan , le pasamos los punteros a las variables de datos
		err := filas.Scan(&matricula, &nombre)

		if err != nil {
			log.Fatal(err)
		}
		//Sacamos por pantalla solo para comprobar que todo funciona
		fmt.Println(matricula, nombre)
		resultado = append(resultado, nombre)

	}

	log.Println(resultado)
	err = filas.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---*************------------***-*-*-*-*-*-")

	//CONSULTA CON SENTENCIA PREPARADA !!!!
	// Primero preparamos la sentencia

	stmt, err := db.Prepare("select nombre, apellido from alumnos where matricula = ?")
	if err != nil {
		log.Fatal(err)
	}
	// Se debe deferir el statement simpre
	defer stmt.Close()

	// Ejecutamos la consulta

	rows, err := stmt.Query(3434)

	if err != nil {
		log.Fatal(err)
	}

	var apellido string

	defer rows.Close()
	// Iteramos los resultados
	for rows.Next() {
		error := rows.Scan(&nombre, &apellido)
		if error != nil {
			log.Fatal(error)
		}
		log.Println(nombre, apellido)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	////  CONSULTA DE 1 SOLA FILA  !!
	///
	/// Se puede simplifica un query que recibe una unica fila como respuesta
	var res string

	err = db.QueryRow("select nombre from alumnos where matricula = 3434").Scan(&res)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	//// CONSULTA DE UNA FILA "PREPARADA"
	stmt2, err := db.Prepare("select nombre from alumnos where matricula = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt2.Close()

	err = stmt2.QueryRow(25).Scan(&nombre)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nombre)

}
