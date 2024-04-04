package model

type Cart struct {
	Id              string
	UserId          string
	ProductQuantity []ProductQuantity
	TotalPrice      float64
	TotalQuantity   int
}

type ProductQuantity struct {
	ProductId string
	Quantity  int
}
