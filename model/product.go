package model

type Product struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stockQuantity"`
	PictureLink   string  `json:"picture"`
}
