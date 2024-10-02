package main

import (
	"Go_API/db"
	"Go_API/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//creando la conexion a la base de datos
	db.DBconnection()
	// Metodo para crear un router
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)
	//inicializando el servidor
	http.ListenAndServe(":3000", router) //puerto en donde se va a ejecutar y el segundo es el router que creamos
}
