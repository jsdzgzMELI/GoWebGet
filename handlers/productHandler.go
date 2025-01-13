package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jsdzgzMELI/GoWeb/helper"
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
