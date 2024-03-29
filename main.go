package main

import (
	"fmt"
	"miRest/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var Version string = "0.1.2"

func main() {
	fmt.Println("Iniciando Servidor - Version " + Version)

	//-----------Router Mux------------------------------------

	miRouter := mux.NewRouter()

	miRouter.HandleFunc("/", controllers.InitController)

	miRouter.HandleFunc(
		"/pathParameter/{data}",
		controllers.PathParameterController).Methods("GET")

	miRouter.HandleFunc(
		"/queryParameter",
		controllers.QueryParameters).Methods("GET")

	miRouter.HandleFunc(
		"/postBodyJson",
		controllers.PostBodyJson,
	).Methods("POST")

	miRouter.HandleFunc(
		"/postFile",
		controllers.PostFile,
	).Methods("POST")

	//-----------Server INIT-----------------------------

	miServer := http.Server{
		Addr:    ":8000",
		Handler: miRouter,
	}

	fmt.Println("Servidor Iniciado - PUERTO : 8000")
	miServer.ListenAndServe()

}
