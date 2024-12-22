package main

import (
	"api-gym-on-go/goose/seed"
	"api-gym-on-go/src/config/env"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {
	env.LoadEnv()
	
	fmt.Println(`
	    _____                           _____ _      _____ 
	   / ____|                         / ____| |    |_   _|
	  | |  __  ___   ___   ___  ___   | |    | |      | |  
	  | | |_ |/ _ \ / _ \ / __|/ _ \  | |    | |      | |  
	  | |__| | (_) | (_) |\__ \  __/  | |____| |____ _| |_ 
	   \_____|\___/ \___/|____/\___|   \_____|______|_____| v0.0.1 🪶
	 `)
	fmt.Printf("Database URL: %s\n", env.DatabaseURL)

	for {
		option := showMenu()

		if option < 0 || option > 6 || reflect.TypeOf(option).Kind() != reflect.Int {
			fmt.Println("Invalid option. Please try again.")
			continue
		}

		switch option {
		case 1:
			applyMigrations(env.DatabaseURL)
		case 2:
			applyMigrations(env.DatabaseURL)
			seed.SeedDatabase(env.DatabaseURL)
		case 3:
			applyLastMigration(env.DatabaseURL)
		case 4:
			resetDatabase(env.DatabaseURL)
		case 5:
			resetDatabase(env.DatabaseURL)
			applyMigrations(env.DatabaseURL)
		case 6:
			resetDatabase(env.DatabaseURL)
			applyMigrations(env.DatabaseURL)
			seed.SeedDatabase(env.DatabaseURL)
		case 7:
			createMigration()
		case 0:
			os.Exit(0)
		}
	}
}

func applyMigrations(databaseURL string) {
	fmt.Println("Applying migrations...")
	runCommand(fmt.Sprintf("GOOSE_DRIVER=postgres GOOSE_DBSTRING=%s goose -dir=goose/migrations up", databaseURL))
	fmt.Println("Migrations applied successfully!")
}

func applyLastMigration(databaseURL string) {
	fmt.Println("Applying last migration...")
	runCommand(fmt.Sprintf("GOOSE_DRIVER=postgres GOOSE_DBSTRING=%s goose -dir=goose/migrations up 1", databaseURL))
	fmt.Println("Last migration applied successfully!")
}

func resetDatabase(databaseURL string) {
	fmt.Println("Reseting goose...")
	runCommand(fmt.Sprintf("GOOSE_DRIVER=postgres GOOSE_DBSTRING=%s goose -dir=goose/migrations reset", databaseURL))
	fmt.Println("Database reseted successfully!")
}

func createMigration() {
	fmt.Print("Enter the name of the migration: ")
	reader := bufio.NewReader(os.Stdin)
	migrationName, _ := reader.ReadString('\n')
	migrationName = strings.TrimSpace(migrationName) // Remove espaços extras e quebras de linha

	if migrationName == "" {
		fmt.Println("Error: The name of the migration cannot be empty.")
		os.Exit(1)
	}

	fmt.Printf("Creating new migration: %s\n", migrationName)
	runCommand(fmt.Sprintf("goose create %s sql --dir=goose/migrations", migrationName))
	fmt.Println("Migration created successfully!")
	os.Exit(0)
}

func runCommand(cmd string) {
	command := exec.Command("bash", "-c", cmd)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Printf("Error running command '%s': %v\n", cmd, err)
		os.Exit(1)
	}
}

func showMenu() int {
	fmt.Println("\n1. Apply all migrations")
	fmt.Println("2. Apply all migrations and seed")
	fmt.Println("3. Apply last migration")
	fmt.Println("4. Drop all migrations")
	fmt.Println("5. Reset database")
	fmt.Println("6. Reset database and seed")
	fmt.Println("7. Create new migration")
	fmt.Println("0. Exit")

	var option int
	fmt.Print("\nSelect an option > ")
	fmt.Scan(&option)

	return option
}
