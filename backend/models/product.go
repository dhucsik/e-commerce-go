package models

type Product struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Seller      User     `json:"seller"`
	Category    Category `json:"category"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	AvgRating   float64  `json:"avg_rating"`
}

type Queries struct {
	Name        string
	StartPrice  string
	EndPrice    string
	StartRating string
	EndRating   string
}
