package utils

import (
	"api-gym-on-go/src/config/env"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Configurações do banco de dados
func SetupEnviromentTest() {
	env.DATABASE_URL = "postgresql://root:admin@localhost:5432/public?sslmode=disable"
	env.JWT_SECRET = "JWT_SECRET"
	env.PORT = 3000

	resetDb(env.DATABASE_URL)
}

func connectDB(database_url string) *sql.DB {
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		panic(fmt.Sprintf("falha ao conectar ao banco de dados: %v", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("falha ao pingar o banco de dados: %v", err))
	}

	return db
}

func resetDb(database_url string) {
	db := connectDB(database_url)

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
