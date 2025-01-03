
package main

import (
	"net/http"
	"text/template"

	"github.com/gpiagentini/db"
	"github.com/gpiagentini/products"
)

var templateList = template.Must(template.ParseGlob("./templates/*.html"))

func main() {
	http.HandleFunc("/", loadIndex)
	http.ListenAndServe(":8080", nil)
}

func loadIndex(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDb()
	defer db.Close()
	products := []products.Product{
		{"T-Shirt", "Nice Green T-Shirt", 29.90, 5},
		{"Notebook", "Very Fast Notebook", 1000.99, 1},
		{"Phone", "Very Good", 634.99, 1},
	}

	templateList.ExecuteTemplate(w, "index", products)
}
