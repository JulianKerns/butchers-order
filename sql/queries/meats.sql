-- name: AddingDefaultBeef :one
INSERT INTO beef(id,created_at,updated_at,meatcut,price)
VALUES($1,$2,$3,$4,$5)
RETURNING *;