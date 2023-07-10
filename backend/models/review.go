package models

type Review struct {
	ID      string  `json:"id"`
	User    User    `json:"user"`
	Product Product `json:"product"`
	Rating  uint    `json:"rating"`
	Comment string  `json:"comment"`
}
