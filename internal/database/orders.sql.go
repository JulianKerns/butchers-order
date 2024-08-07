// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: orders.sql

package database

import (
	"context"
)

const makeOrder = `-- name: MakeOrder :one
INSERT INTO orders(id, created_at, updated_at, user_id, meat_source, meatcut, price, quantitiy)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING id, created_at, updated_at, user_id, meat_source, meatcut, price, quantitiy
`

type MakeOrderParams struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	UserID     string
	MeatSource string
	Meatcut    string
	Price      float64
	Quantitiy  float64
}

func (q *Queries) MakeOrder(ctx context.Context, arg MakeOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, makeOrder,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.MeatSource,
		arg.Meatcut,
		arg.Price,
		arg.Quantitiy,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.MeatSource,
		&i.Meatcut,
		&i.Price,
		&i.Quantitiy,
	)
	return i, err
}
