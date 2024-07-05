-- name: AddingDefaultMeatPrices :one
INSERT INTO meatprices(id, created_at, updated_at, meat_source, meatcut, price)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetAllDefaultMeatPrices :many
SELECT * FROM meatprices;

-- name: DeleteAllEntries :exec
DELETE FROM meatprices;