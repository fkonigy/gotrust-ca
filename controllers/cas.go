package controllers

import (
	"html/template"
	"net/http"

	"github.com/fkonigy/gotrust-ca/models"
	"github.com/fkonigy/gotrust-ca/utils"
	"github.com/gorilla/mux"
)

func Cas(w http.ResponseWriter, r *http.Request) {
	db, err := models.DB()
	var cas []models.Ca
	_, err = db.Select(&cas, "SELECT * FROM cas")
	if err != nil {
		utils.CheckErr(err, "Error in rendering cas.html")
	}

	tmpl := template.Must(template.ParseFiles("templates/cas.html"))
	err = tmpl.Execute(w, struct{ List []models.Ca }{List: cas})
	utils.CheckErr(err, "Error in rendering cas.html")
}

func ShowCa(w http.ResponseWriter, r *http.Request) {
	db, err := models.DB()
	if err != nil {
		panic(err)
	}
	id := mux.Vars(r)["id"]
	var ca models.Ca
	err = db.SelectOne(&ca, "select \"ID\", \"Name\" from cas where \"ID\" = $1", id)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("templates/ca_show.html"))
	err = tmpl.Execute(w, ca)
	utils.CheckErr(err, "Error in rendering cas.html")
}

func NewCa(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/ca_new.html"))
	err := tmpl.Execute(w, nil)
	utils.CheckErr(err, "Error in rendering ca_new.html")
}

func CreateCa(w http.ResponseWriter, r *http.Request) {
	db, err := models.DB()
	if err != nil {
		panic(err)
	}
	r.ParseForm()
	name := r.PostFormValue("name")
	ca := models.Ca{-1, name}
	err = db.Insert(&ca)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/cas", 302)
}
