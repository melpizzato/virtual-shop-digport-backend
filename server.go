package main

import (
	"encoding/json"
	"net/http"

	"github.com/melpizzato/virtual-shop-digport-backend/model"
)

func StartServer() {
	http.HandleFunc("/products", productsHandler)

	http.ListenAndServe(":8080", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getProducts(w, r)
	} else if r.Method == "POST" {
		addProduct(w, r)
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

	registerProduct(product)

	w.WriteHeader(http.StatusCreated)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("nome")
	if queryNome != "" {
		productsFilteredByName := productsByName(queryNome)
		json.NewEncoder(w).Encode(productsFilteredByName)
	} else {
		products := Products
		json.NewEncoder(w).Encode(products)
	}
}
