package models

import (
	"github.com/JulianKerns/butchers-order/internal/database"
)

type Order struct {
	ID         string  `json:"id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	UserID     string  `json:"user_id"`
	MeatSource string  `json:"meat_source"`
	Meatcut    string  `json:"meatcut"`
	Price      float64 `json:"price"`
	Quantitiy  float64 `json:"quantity"`
}

func DatabaseOrdertoOrder(order database.Order) Order {
	return Order{
		ID:         order.ID,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		UserID:     order.UserID,
		MeatSource: order.MeatSource,
		Meatcut:    order.Meatcut,
		Price:      order.Price,
		Quantitiy:  order.Quantitiy,
	}
}
