package models

type Cart struct {
	ID       string  `json:"id"`
	User     User    `json:"user"`
	Product  Product `json:"product"`
	Quantity uint    `json:"quantity"`
}
