package model

import (
	"fmt"

	"github.com/melpizzato/virtual-shop-digport-backend/db"
)

type Product struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"stockQuantity"`
	Picture     string  `json:"picture"`
}

var id, nome string
var preco float64
var descricao, imagem string
var quantidade int

func GetAllProducts() []Product {
	db := db.ConnectDatabase()

	result, err := db.Query("SELECT * FROM produtos")
	err = result.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	products := []Product{}

	for result.Next() {
		product.Id = id
		product.Name = nome
		product.Description = descricao
		product.Price = preco
		product.Picture = imagem
		product.Quantity = quantidade

		products = append(products, product)
	}

	defer db.Close()

	return products
}

func GetProductsByName(productName string) (Product, error) {
	db := db.ConnectDatabase()

	res := db.QueryRow("SELECT * FROM produtos WHERE nome = $1", productName)

	err := res.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)

	if err != nil {
		return Product{}, fmt.Errorf("Product with name %s not found.", productName)
	}

	var p Product
	p.Id = id
	p.Name = nome
	p.Description = descricao
	p.Price = preco
	p.Picture = imagem
	p.Quantity = quantidade

	defer db.Close()
	return p, nil
}
