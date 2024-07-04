-- name: MakeOrder :one
INSERT INTO orders(id, created_at, updated_at, user_id, meat_source, meatcut, price, quantitiy)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING * ;

