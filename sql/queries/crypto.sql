-- name: CreateCryptoList :one
INSERT INTO crypto (id, updated_at, crypto_key, crypto_list)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCryptoList :one
SELECT * FROM crypto
WHERE crypto_key = $1 LIMIT 1;

-- name: UpdateCryptoList :one
UPDATE crypto
SET updated_at = $2, crypto_key = $3, crypto_list = $4
WHERE crypto_key = $1
RETURNING *;

-- name: DeleteCryptoList :many
DELETE FROM crypto
WHERE crypto_key = $1
RETURNING *;
