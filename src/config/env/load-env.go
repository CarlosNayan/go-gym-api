package env

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	DatabaseURL string
	JWTSecret   string
	Port        int
)

func LoadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err := loadDotEnv(".env"); err != nil {
			panic(fmt.Sprintf("Erro ao carregar .env file: %v", err))
		}

		port, err := strconv.Atoi(getEnv("PORT", "3333"))
		if err != nil {
			panic(fmt.Sprintf("PORT deve ser um número válido: %v", err))
		}

		JWTSecret = getEnv("JWT_SECRET", "")
		DatabaseURL = getEnv("DATABASE_URL", "")
		Port = port

		if err := validateEnv(); err != nil {
			panic(fmt.Sprintf("Erro ao carregar variáveis de ambiente: %v", err))
		}
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func validateEnv() error {
	if JWTSecret == "" {
		return fmt.Errorf("JWT_SECRET não pode estar vazio")
	}
	if DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL não pode estar vazio")
	}
	return nil
}

func loadDotEnv(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo .env: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("linha inválida no .env: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if (strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"")) ||
			(strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'")) {
			value = strings.Trim(value, "\"'")
		}

		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("erro ao definir variável %s: %w", key, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("erro ao ler o arquivo .env: %w", err)
	}

	return nil
}
