package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // Driver para PostgreSQL
)

func SetupDatabase(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Configurações do pool de conexões
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	// Verifica a conexão
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return db
}
