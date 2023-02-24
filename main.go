package main

import (
	"net/http"

	"github.com/RickHPotter/web_app_alura_course/routes"
)

func main() {
	routes.LoadRoute()
	http.ListenAndServe("localhost:8000", nil)
}
