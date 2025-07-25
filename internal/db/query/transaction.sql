-- name: CreateTransaction :one
INSERT INTO transactions (user_id, transaction_id, source_type, state, amount)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, transaction_id, source_type, state, amount, created_at;

-- name: GetTransaction :one
SELECT id, user_id, transaction_id, source_type, state, amount, created_at
FROM transactions
WHERE transaction_id = $1;

-- name: GetTransactionsByUser :many
SELECT id, user_id, transaction_id, source_type, state, amount, created_at
FROM transactions
WHERE user_id = $1
ORDER BY created_at DESC;