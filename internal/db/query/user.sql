-- name: GetUserById :one
SELECT id, balance FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: UpdateUserBalance :exec
UPDATE users
set balance = $2,
    updated_at = NOW()
WHERE id = $1;