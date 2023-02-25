package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/RickHPotter/web_app_alura_course/models"
)

var tmplt = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListAll()
	tmplt.ExecuteTemplate(w, "index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmplt.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	// insert

	if r.Method == "POST" {

		price, e1 := strconv.ParseFloat(r.FormValue("price"), 64)
		stock, e2 := strconv.Atoi(r.FormValue("stock"))
		if e1 != nil || e2 != nil {
			log.Println(e1.Error(), e2.Error())
		}

		product := models.NewProduct(
			r.FormValue("name"),
			r.FormValue("description"),
			price,
			stock,
		)
		product.Insert()

	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
