package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Configurações do banco de dados
const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "root"
	dbPassword = "admin"
	dbName     = "public"
)

func ConnectDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("falha ao conectar ao banco de dados: %w", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("falha ao pingar o banco de dados: %w", err))
	}

	return db
}

func ResetDb() {
	db := ConnectDB()

	defer db.Close()

	tables := []string{"users", "gyms", "checkins"}

	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", table)
		if _, err := db.Exec(query); err != nil {
			panic(fmt.Sprintf("falha ao limpar a tabela %s: %v", table, err))
		}
	}

	log.Println("Tabelas limpas com sucesso")
}
