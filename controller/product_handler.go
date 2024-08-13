package controller

import (
	"encoding/json"
	"net/http"

	"github.com/melpizzato/virtual-shop-digport-backend/model"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := model.GetAllProducts()
	json.NewEncoder(w).Encode(products)
}

func GetProductsByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("nome")
	product, err := model.GetProductsByName(name)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Error{ErrorMessage: err.Error()})
		return
	}
	json.NewEncoder(w).Encode((product))
}

func CreateProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []model.Product
	json.NewDecoder(r.Body).Decode(&products)
}
