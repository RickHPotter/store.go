package routes

import (
	"net/http"

	"github.com/RickHPotter/web_app_alura_course/controllers"
)

func LoadRoute() {
	//ACTIONS
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)

	// SCREENS
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/edit", controllers.Edit)
}
