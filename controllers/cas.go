package controllers

import (
	"html/template"
	"net/http"

	"github.com/fkonigy/gotrust-ca/utils"
)

func CasController(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/cas.html"))
	err := tmpl.Execute(w, nil)
	utils.CheckErr(err, "Error in rendering cas.html")
}
