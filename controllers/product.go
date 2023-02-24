package controllers

import (
	"html/template"
	"net/http"

	"github.com/RickHPotter/web_app_alura_course/models"
)

var tmplt = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListAll()
	tmplt.ExecuteTemplate(w, "index", products)
}
