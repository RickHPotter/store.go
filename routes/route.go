package routes

import (
	"net/http"

	"github.com/RickHPotter/web_app_alura_course/controllers"
)

func LoadRoute() {
	http.HandleFunc("/", controllers.Index)
}
