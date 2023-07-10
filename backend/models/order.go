package models

import "time"

type Order struct {
	ID           string        `json:"id"`
	User         User          `json:"user"`
	TotalPrice   float64       `json:"total_price"`
	Date         time.Time     `json:"date"`
	Address      string        `json:"address"`
	OrderedItems []OrderedItem `json:"ordered_items,omitempty"`
}
