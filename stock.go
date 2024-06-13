package main

import (
	"errors"

	"github.com/melpizzato/virtual-shop-digport-backend/model"
)

var Products []model.Product = []model.Product{}

func productsByName(name string) []model.Product {
	result := []model.Product{}

	for _, product := range Products {
		if product.Name == name {
			result = append(result, product)
		}
	}

	return result
}

func registerProduct(newProduct model.Product) error {
	for _, product := range Products {
		if newProduct.Id == product.Id {
			return errors.New("Product already exists!")
		}
	}
	Products = append(Products, newProduct)
	return nil
}
