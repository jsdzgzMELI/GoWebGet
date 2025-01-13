package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jsdzgzMELI/GoWeb/helper"
	"github.com/jsdzgzMELI/GoWeb/structs"
)

type Controller struct {
	st map[string]string
}

func (c *Controller) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	products, err := helper.LoadProducts("./products.json")
	if err != nil {
		log.Fatalf("Error loading products: %v", err)
	}
	json.NewEncoder(w).Encode(products)
}

func (c *Controller) GetById(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
	// product, err := helper.GetProductById(id, "./products.json")
	id := chi.URLParam(r, "id")
	product, ok := c.st[id]
	if !ok {
		code := http.StatusNotFound
		body := &structs.ResponseId{Message: "product not found", Data: nil}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
	code := http.StatusOK
	// AQUI TOCA ITERAR LOS PRODUCTOS Y BUSCAR EL PRODUCTO POR ID
	body := &structs.ResponseId{Message: "product found", Data: &structs.Product{
		ID:           id,
		Name:         product,
		Quantity:     0,
		Code_value:   "",
		Is_published: false,
		Expiration:   "",
		Price:        0.0,
	}}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}
