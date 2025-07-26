package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	db "github.com/chavocito/entain/internal/db/sqlc"
	"github.com/chavocito/entain/utility"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserTransactionHandler handles POST /user/{id}/transaction
func UserTransactionHandler(dbPool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 4 || parts[1] != "user" || parts[3] != "transaction" {
			http.NotFound(w, r)
			return
		}
		userID, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Parse and validate source type from header
		sourceType := r.Header.Get("Source-Type")
		if !utility.IsHeaderValid(sourceType) {
			http.Error(w, "Invalid or missing Source-Type header", http.StatusBadRequest)
			return
		}

		var req db.TransactionTxParams
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if req.State != "win" && req.State != "lose" {
			http.Error(w, "Invalid state: must be 'win' or 'lose'", http.StatusBadRequest)
			return
		}
		amount, err := strconv.ParseFloat(req.Amount, 64)
		if err != nil || amount <= 0 {
			http.Error(w, "Invalid amount", http.StatusBadRequest)
			return
		}
		if req.TransactionID == "" {
			http.Error(w, "Missing transaction_id", http.StatusBadRequest)
			return
		}

		conn, err := dbPool.Acquire(r.Context())
		if err != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer conn.Release()

		txErr := db.ProcessUserTransactionTx(r.Context(), conn.Conn(), req.State, userID, amount, req.TransactionID, sourceType)
		if txErr != nil {
			if txErr.Error() == "insufficient balance" {
				http.Error(w, "Insufficient balance", http.StatusBadRequest)
				return
			}
			http.Error(w, txErr.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Transaction successful"))
	}
}
