package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/melpizzato/virtual-shop-digport-backend/controller"
)

func HandleRequests() {
	route := mux.NewRouter()

	route.HandleFunc("/products", controller.GetProductsHandler).Methods("GET")
	route.HandleFunc("/product", controller.GetProductsByNameHandler).Methods("GET")
	route.HandleFunc("/products", controller.CreateProductsHandler).Methods("POST")
	http.ListenAndServe(":8080", route)
}
