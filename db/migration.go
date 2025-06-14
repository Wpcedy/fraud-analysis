package db

import (
    "database/sql"
    "fmt"
)

func RunMigrations(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS analysis (
        id SERIAL PRIMARY KEY,
        cpf VARCHAR(11) NOT NULL,
        dados_enviados JSONB NOT NULL,
        resposta_ia JSONB NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `

    _, err := db.Exec(query)
    if err != nil {
        return fmt.Errorf("erro ao executar migration: %v", err)
    }

    return nil
}
