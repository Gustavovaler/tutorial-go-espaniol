package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

// Creamos un type del tipo estructura que será nuetro objeto de tipo Usuario
type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

// creamos un map para guardar nuestros usuarios
var (
	users = map[int]*user{}
	seq   = 1
)

//----------
// Handlers
//----------
// aqui definimos los "controladores" de cada ruta
func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func index(c echo.Context) error {

	return c.JSON(http.StatusOK, "Hola desde index")
}

func main() {
	e := echo.New()

	// Middleware
	// con este middleware recibimos por io todas las peticiones
	//e.Use(middleware.Logger())
	// con este recuperamos las rutas en caso de error para que no caiga el server
	e.Use(middleware.Recover())

	// Rutas

	e.GET("/", index)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Start(":1323")
}
