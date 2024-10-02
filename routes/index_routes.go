package routes

import "net/http"

func HomeHandler(w http.ResponseWriter, request *http.Request) {
	//los parametros de la funcion son: writer que sera la forma en la que se respondera al clente y el read es para pode tener acceso a los parametros que me envien
	w.Write([]byte("Hello Word"))

}
