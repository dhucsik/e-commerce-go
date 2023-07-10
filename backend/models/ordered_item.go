package models

type OrderedItem struct {
	ID       string  `json:"id"`
	Order    Order   `json:"order"`
	Product  Product `json:"product"`
	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`
}
