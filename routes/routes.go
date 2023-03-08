package routes

import (
	"curso-alura-golang-I/controller"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Remove)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.Update)
}
