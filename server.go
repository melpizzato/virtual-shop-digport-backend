package main

import "github.com/melpizzato/virtual-shop-digport-backend/routes"

func StartServer() {
	routes.HandleRequests()
}
