package helper

import (
	"encoding/json"
	"os"

	"github.com/jsdzgzMELI/GoWeb/structs"
)

func loadProducts(filename string) ([]structs.Product, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []structs.Product
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
