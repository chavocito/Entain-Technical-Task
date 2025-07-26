package db

import (
	"context"
	"errors"
	"strings"

	"github.com/chavocito/entain/utility"
	"github.com/jackc/pgx/v5"
)

type TransactionTxParams struct {
	State         string `json:"state"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transactionId"`
}

// ProcessUserTransactionTx updates the user's balance and creates a transaction record atomically.
func ProcessUserTransactionTx(ctx context.Context, db *pgx.Conn, state string, userID int64, amount float64, transactionID string, sourceType string) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	q := New(tx)
	user, err := q.GetUserForUpdate(ctx, userID)
	if err != nil {
		return errors.New("Error processing request")
	}

	currentBalance := utility.ConvertAmountToDecimalForDisplay(user.Balance)

	if strings.ToLower(state) == "win" {
		currentBalance += amount
	} else if strings.ToLower(state) == "lose" {
		if currentBalance < amount {
			return errors.New("insufficient balance")
		}
		currentBalance -= amount
	} else {
		return errors.New("invalid state")
	}

	updated := utility.ConvertAmountToDBRepresentation(currentBalance)

	err = q.UpdateUserBalance(ctx, UpdateUserBalanceParams{
		ID:      userID,
		Balance: updated,
	})
	if err != nil {
		return err
	}

	transactionAmountForDB := utility.ConvertAmountToDBRepresentation(amount)
	_, err = q.CreateTransaction(ctx, CreateTransactionParams{
		UserID:        userID,
		TransactionID: transactionID,
		SourceType:    sourceType,
		State:         state,
		Amount:        transactionAmountForDB,
	})
	if err != nil {
		return err
	}

	return nil
}
