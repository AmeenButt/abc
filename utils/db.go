package utils

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func ConnectDB(ctx context.Context) *pgx.Conn {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	conn, err := pgx.Connect(ctx, os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	// create tables

	sqlFile, err := os.ReadFile("db/schema/users.sql")
	if err != nil {
		log.Fatalf("can not read users.sql: %v", err)
	}
	_, err = conn.Exec(ctx, string(sqlFile))
	if err != nil {
		log.Fatalf("unable to create users table: %v", err)
	}

	return conn
}
