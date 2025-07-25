package handler

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	db "github.com/chavocito/entain/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetUserBalanceHandler(dbPool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		fmt.Println("parts:", parts)
		if len(parts) < 4 || parts[1] != "user" || parts[3] != "balance" {
			http.NotFound(w, r)
			return
		}
		userID, err := strconv.ParseInt(parts[2], 10, 64)
		fmt.Println("userID:", userID)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		conn, err := dbPool.Acquire(r.Context())
		fmt.Println("connectionPool:", conn)
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
		fmt.Println("User:", user)
		// var balanceFloat float64
		// if err := user.Balance.Scan(&balanceFloat); err != nil {
		// 	http.Error(w, "Error reading balance: "+err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		var balanceStr string
		if err := user.Balance.Scan(balanceStr); err != nil {
			http.Error(w, "Error reading balance: "+err.Error(), http.StatusInternalServerError)
			return
		}
		balanceFloat, err := strconv.ParseFloat(balanceStr, 64)
		if err != nil {
			http.Error(w, "Error parsing balance: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Round to 2 decimal places
		rounded := math.Round(balanceFloat*100) / 100

		resp := struct {
			UserID  uint64 `json:"userId"`
			Balance string `json:"balance"`
		}{
			UserID:  uint64(userID),
			Balance: strconv.FormatFloat(rounded, 'f', 2, 64),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
