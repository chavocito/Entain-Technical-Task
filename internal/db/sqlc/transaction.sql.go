// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: transaction.sql

package db

import (
	"context"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (user_id, transaction_id, source_type, state, amount)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, transaction_id, source_type, state, amount, created_at
`

type CreateTransactionParams struct {
	UserID        int64  `json:"user_id"`
	TransactionID string `json:"transaction_id"`
	SourceType    string `json:"source_type"`
	State         string `json:"state"`
	Amount        int64  `json:"amount"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.UserID,
		arg.TransactionID,
		arg.SourceType,
		arg.State,
		arg.Amount,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TransactionID,
		&i.SourceType,
		&i.State,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, user_id, transaction_id, source_type, state, amount, created_at
FROM transactions
WHERE transaction_id = $1
`

func (q *Queries) GetTransaction(ctx context.Context, transactionID string) (Transaction, error) {
	row := q.db.QueryRow(ctx, getTransaction, transactionID)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.TransactionID,
		&i.SourceType,
		&i.State,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransactionsByUser = `-- name: GetTransactionsByUser :many
SELECT id, user_id, transaction_id, source_type, state, amount, created_at
FROM transactions
WHERE user_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetTransactionsByUser(ctx context.Context, userID int64) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getTransactionsByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.TransactionID,
			&i.SourceType,
			&i.State,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
