package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// ConnectDB connect to the pgsql db
func ConnectDB() (*pgxpool.Pool, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("database environment variables are not set")
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return pool, nil
}
