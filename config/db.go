package config

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "os"
)

func ConnectToDB() (*sql.DB, error) {
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbUser, dbPass, dbName,
    )

    return sql.Open("postgres", dsn)
}
