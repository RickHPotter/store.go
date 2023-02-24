package main

import (
	"html/template"
	"net/http"

	"github.com/RickHPotter/web_app_alura_course/db"
	"github.com/RickHPotter/web_app_alura_course/models"
)

var tmplt = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	db := db.Connect()
	defer db.Close()

	query, e := db.Query("select * from products")
	if e != nil {
		panic(e.Error())
	}

	// p := models.Product{}
	products := []models.Product{}

	for query.Next() {
		var id uint16
		var stock int
		var name, description string
		var price float64

		e = query.Scan(&id, &name, &description, &price, &stock)
		if e != nil {
			panic(e.Error())
		}

		p := models.NewProduct(id, name, description, price, stock)
		products = append(products, *p)
	}

	// products := []models.Product{
	// 	*models.NewProduct("T-Shirt", "Blue & Elegant", 36, 66),
	// 	*models.NewProduct("Laptop", "Ligthing Fast", 1600, 14),
	// 	*models.NewProduct("Air Jordan", "Self-explanatory", 260, 2),
	// }

	tmplt.ExecuteTemplate(w, "index", products)
}
