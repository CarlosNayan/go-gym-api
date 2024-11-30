package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EnvConfig struct {
	NodeEnv     string
	DatabaseURL string
	JWTSecret   string
	Port        int
}

func LoadEnv() *EnvConfig {
	if _, err := os.Stat(".env"); err == nil {
		if err := loadDotEnv(".env"); err != nil {
			panic(fmt.Sprintf("erro ao carregar .env: %w", err))
		}
	}

	port, err := strconv.Atoi(getEnv("PORT", "3333"))
	if err != nil {
		panic(fmt.Sprintf("PORT deve ser um número válido: %v", err))
	}

	env := &EnvConfig{
		JWTSecret:   getEnv("JWT_SECRET", ""),
		NodeEnv:     getEnv("NODE_ENV", "dev"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		Port:        port,
	}

	if err := validateEnv(env); err != nil {
		panic(fmt.Sprintf("erro ao carregar variáveis de ambiente: %v", err))
	}

	return env
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func validateEnv(config *EnvConfig) error {
	if config.NodeEnv != "dev" && config.NodeEnv != "test" && config.NodeEnv != "production" {
		panic("NODE_ENV deve ser 'dev', 'test' ou 'production'")
	}
	if config.DatabaseURL == "" {
		panic("DATABASE_URL deve ser uma URL válida")
	}
	if config.JWTSecret == "" {
		panic("JWT_SECRET deve ser uma string")
	}
	return nil
}

func loadDotEnv(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			lines = append(lines, strings.Split(string(buf[:n]), "\n")...)
		}
		if err != nil {
			break
		}
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			panic(fmt.Sprintf("linha inválida no .env: %s", line))
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	return nil
}
