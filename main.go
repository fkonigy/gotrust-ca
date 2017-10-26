package main

import (
	// "log"
	"net/http"

	"github.com/fkonigy/gotrust-ca/controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeController).Methods("GET")
	r.HandleFunc("/cas", controllers.CasController).Methods("GET")

	// TODO: create a Server and use with data from config file, instead of
	// using http.ListenAndServer.
	http.ListenAndServe(":8080", r)
}
