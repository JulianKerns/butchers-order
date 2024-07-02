-- name: AddingDefaultBeef :one
INSERT INTO beef(id,created_at,updated_at,meatcut,price)
VALUES($1,$2,$3,$4,$5)
RETURNING *;

-- name: AddingDefaultPork :one
INSERT INTO pork(id,created_at,updated_at,meatcut,price)
VALUES($1,$2,$3,$4,$5)
RETURNING *;

-- name: AddingDefaultSaltedPork :one
INSERT INTO saltedpork(id,created_at,updated_at,meatcut,price)
VALUES($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetAllDefaultBeef :many
SELECT * FROM beef;

-- name: GetAllDefaultPork :many
SELECT * FROM pork;

-- name: GetAllDefaultSaltedPork :many
SELECT * FROM saltedpork;