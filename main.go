package main

import (
	"fmt"
	"net/http"

	"miRest/controllers"

	"github.com/gorilla/mux"
)

var Version string = "0.1.1"

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

	miServer.ListenAndServe()

}
