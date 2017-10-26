package main

import (
	// "log"
	"net/http"

	"github.com/fkonigy/gotrust-ca/controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/cas", controllers.Cas)
	r.HandleFunc("/ca/{id:[0-9]+}", controllers.ShowCa)
	r.HandleFunc("/ca/new", controllers.NewCa)
	r.HandleFunc("/cas/", controllers.CreateCa)

	// TODO: create a Server and use with data from config file, instead of
	// using http.ListenAndServer.
	http.ListenAndServe(":8080", r)
}
