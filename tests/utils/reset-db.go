package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Configurações do banco de dados
const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "root"
	dbPassword = "admin"
	dbName     = "api_solid"
)

func ConnectDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(fmt.Errorf("falha ao conectar ao banco de dados: %w", err))
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(fmt.Errorf("falha ao pingar o banco de dados: %w", err))
		os.Exit(1)
	}

	return db
}

func ResetDb() {
	db := ConnectDB()

	defer db.Close()

	tables := []string{"gyms", "users", "checkins"}

	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", table)
		if _, err := db.Exec(query); err != nil {
			log.Printf("falha ao limpar a tabela %s: %v", table, err)
			log.Fatal(fmt.Errorf("falha ao pingar o banco de dados: %w", err))
			os.Exit(1)
		}
	}

	log.Println("Tabelas limpas com sucesso")
}
