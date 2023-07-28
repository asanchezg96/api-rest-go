package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Api-GO")

	//Crear mux
	//Rutas
	mux := mux.NewRouter()

	//Endpoints
	mux.HandleFunc("/api/user/", nil).Methods("GET")            //Obtener todos los usuarios
	mux.HandleFunc("/api/user/{id:[0-9]+}", nil).Methods("GET") //Obtener un usuario por id

	mux.HandleFunc("/api/user/", nil).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", nil).Methods("POST") //Registrar usuarios
	mux.HandleFunc("/api/user/{id:[0-9]+}", nil).Methods("DELETE")

}
