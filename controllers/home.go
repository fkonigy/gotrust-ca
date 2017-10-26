package controllers

import (
	"html/template"
	// "log"
	"net/http"

	"github.com/fkonigy/gotrust-ca/utils"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	err := tmpl.Execute(w, nil)
	utils.CheckErr(err, "Error in rendering home.html")
}
