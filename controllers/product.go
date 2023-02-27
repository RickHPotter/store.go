package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/RickHPotter/web_app_alura_course/models"
)

var tmplt = template.Must(template.ParseGlob("templates/*.html"))

/*
ACTIONS
*/

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		price, e1 := strconv.ParseFloat(r.FormValue("price"), 64)
		stock, e2 := strconv.Atoi(r.FormValue("stock"))
		if e1 != nil || e2 != nil {
			log.Println(e1.Error(), e2.Error())
		}

		product := models.NewProduct(
			0, // the only id that matters is the serial id from the db
			r.FormValue("name"),
			r.FormValue("description"),
			price,
			stock,
		)
		product.Insert()

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, e1 := strconv.Atoi(r.FormValue("id"))
		price, e2 := strconv.ParseFloat(r.FormValue("price"), 64)
		stock, e3 := strconv.Atoi(r.FormValue("stock"))
		if e1 != nil || e2 != nil || e3 != nil {
			log.Println(e1.Error(), e2.Error(), e3.Error())
		}

		product := models.NewProduct(
			uint16(id),
			r.FormValue("name"),
			r.FormValue("description"),
			price,
			stock,
		)
		product.Update()

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	codprod := r.URL.Query().Get("id")
	models.Delete(codprod)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

/*
SCREENS
*/

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListAll()
	tmplt.ExecuteTemplate(w, "index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmplt.ExecuteTemplate(w, "new", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	codprod := r.URL.Query().Get("id")
	products := models.ListSome(codprod)
	tmplt.ExecuteTemplate(w, "edit", products)
}
