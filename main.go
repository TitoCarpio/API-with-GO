package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// Metodo para crear un router
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) { //los parametros de la funcion son: writer que sera la forma en la que se respondera al clente y el read es para pode tener acceso a los parametros que me envien
		w.Write([]byte("Hello Word"))

	})
	//inicializando el servidor
	http.ListenAndServe(":3000", router) //puerto en donde se va a ejecutar y el segundo es el router que creamos
}
