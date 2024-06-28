// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: meats.sql

package database

import (
	"context"
)

const addingDefaultBeef = `-- name: AddingDefaultBeef :one
INSERT INTO beef(id,created_at,updated_at,meatcut,price)
VALUES($1,$2,$3,$4,$5)
RETURNING id, created_at, updated_at, meatcut, price, quantitiy
`

type AddingDefaultBeefParams struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	Meatcut   string
	Price     float64
}

func (q *Queries) AddingDefaultBeef(ctx context.Context, arg AddingDefaultBeefParams) (Beef, error) {
	row := q.db.QueryRowContext(ctx, addingDefaultBeef,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Meatcut,
		arg.Price,
	)
	var i Beef
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Meatcut,
		&i.Price,
		&i.Quantitiy,
	)
	return i, err
}