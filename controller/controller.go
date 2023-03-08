package controller

import (
	"curso-alura-golang-I/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

/*
*
carregando os templates
*
*/
var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := model.ListAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		price_converter, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Print("Converter erro to price:", err)
		}

		amount_converter, err := strconv.Atoi(amount)
		if err != nil {
			log.Print("Converter erro to amount:", err)
		}

		model.CreateNewProduct(name, description, price_converter, amount_converter)
	}

	http.Redirect(w, r, "/", 301)
}

func Remove(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	model.RemoveProduct(id)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	product := model.Edit(id)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		price := r.FormValue("price")
		description := r.FormValue("description")
		amount := r.FormValue("amount")

		price_converter, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Print("Converter erro to price:", err)
		}

		amount_converter, err := strconv.Atoi(amount)
		if err != nil {
			log.Print("Converter erro to amount:", err)
		}

		id_converter, err := strconv.Atoi(id)
		if err != nil {
			log.Print("Converter erro to id:", err)
		}

		model.Update(id_converter, name, description, price_converter, amount_converter)
	}

	http.Redirect(w, r, "/", 301)
}
