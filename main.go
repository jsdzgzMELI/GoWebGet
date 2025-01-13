package main

import (
	"net/http"

	"github.com/jsdzgzMELI/GoWeb/handlers"
)

func main() {
	http.HandleFunc("/products", handlers.ProductHandler)

	// r := chi.NewRouter()

	// fmt.Println(products[0])
	http.ListenAndServe(":8080", nil)
}
