package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chavocito/entain/internal/db"
	handler "github.com/chavocito/entain/internal/handler"
	"github.com/gorilla/mux"
)

func main() {
	connectionPool, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to databse")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Println("Server is up and running") })
	r.HandleFunc("/user/{userId}/transaction", handler.UserTransactionHandler(connectionPool)).Methods("POST")
	r.HandleFunc("/user/{userId}/balance", handler.GetUserBalanceHandler(connectionPool)).Methods("GET")

	http.ListenAndServe(":8080", r)
}
