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

func GetUserBalanceHandler(dbPool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 4 || parts[1] != "user" || parts[3] != "balance" {
			http.NotFound(w, r)
			return
		}
		userID, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		conn, err := dbPool.Acquire(r.Context())
		if err != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer conn.Release()

		q := db.New(conn.Conn())
		user, err := q.GetUserById(r.Context(), userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		balanceStr := strconv.FormatFloat(utility.ConvertAmountToDecimalForDisplay(user.Balance), 'f', 2, 64)
		resp := struct {
			UserID  uint64 `json:"userId"`
			Balance string `json:"balance"`
		}{
			UserID:  uint64(userID),
			Balance: balanceStr,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
