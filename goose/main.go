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
	// Arte ASCII
	fmt.Println(`
	    _____                           _____ _      _____ 
	   / ____|                         / ____| |    |_   _|
	  | |  __  ___   ___   ___  ___   | |    | |      | |  
	  | | |_ |/ _ \ / _ \ / __|/ _ \  | |    | |      | |  
	  | |__| | (_) | (_) |\__ \  __/  | |____| |____ _| |_ 
	   \_____|\___/ \___/|____/\___|   \_____|______|_____|
	 `)

	for {
		option := showMenu()

		if option < 0 || option > 8 || reflect.TypeOf(option).Kind() != reflect.Int {
			fmt.Println("Invalid option. Please try again.")
			continue
		}

		switch option {
		case 1:
			confirmation := confirmation("apply all migrations")
			if !confirmation {
				return
			}
			applyMigrations(env.DatabaseURL)
		case 2:
			confirmation := confirmation("apply all migrations and seed")
			if !confirmation {
				return
			}
			applyMigrations(env.DatabaseURL)
			seed.SeedDatabase(env.DatabaseURL)
		case 3:
			confirmation := confirmation("apply last migration")
			if !confirmation {
				return
			}
			applyLastMigration(env.DatabaseURL)
		case 4:
			confirmation := confirmation("drop all migrations")
			if !confirmation {
				return
			}
			resetDatabase(env.DatabaseURL)
		case 5:
			confirmation := confirmation("drop last migration")
			if !confirmation {
				return
			}
			dropLastMigration(env.DatabaseURL)
		case 6:
			confirmation := confirmation("reset database")
			if !confirmation {
				return
			}
			resetDatabase(env.DatabaseURL)
			applyMigrations(env.DatabaseURL)
		case 7:
			confirmation := confirmation("reset database and seed")
			if !confirmation {
				return
			}
			resetDatabase(env.DatabaseURL)
			applyMigrations(env.DatabaseURL)
			seed.SeedDatabase(env.DatabaseURL)
		case 8:
			createMigration()
		case 0:
			os.Exit(0)
		}
	}
}

func applyMigrations(DatabaseURL string) {
	fmt.Println("Applying migrations...")
	runCommand(fmt.Sprintf("GOOSE_DRIVER=postgres GOOSE_DBSTRING=%s goose -dir=goose/migrations up", DatabaseURL))
	fmt.Println("Migrations applied successfully!")
}

func applyLastMigration(DatabaseURL string) {
	fmt.Println("Applying last migration...")
	runCommand(fmt.Sprintf("GOOSE_DRIVER=postgres GOOSE_DBSTRING=%s goose -dir=goose/migrations up 1", DatabaseURL))
	fmt.Println("Last migration applied successfully!")
}

func dropLastMigration(DatabaseURL string) {
	fmt.Println("Dropping last migration...")
	runCommand(fmt.Sprintf("GOOSE_DRIVER=postgres GOOSE_DBSTRING=%s goose -dir=goose/migrations down 1", DatabaseURL))
	fmt.Println("Last migration dropped successfully!")
}

func resetDatabase(DatabaseURL string) {
	fmt.Println("Reseting goose...")
	runCommand(fmt.Sprintf("GOOSE_DRIVER=postgres GOOSE_DBSTRING=%s goose -dir=goose/migrations reset", DatabaseURL))
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
	fmt.Println("5. Drop last migration")
	fmt.Println("6. Reset database")
	fmt.Println("7. Reset database and seed")
	fmt.Println("8. Create new migration")
	fmt.Println("0. Exit")

	var option int
	fmt.Print("\nSelect an option > ")
	fmt.Scan(&option)

	return option
}

func confirmation(action string) bool {
	var confirmation string
	fmt.Printf("Are you sure you want to \x1b[31m%s\x1b[0m? \nSelected database URL: \x1b[31m%s\x1b[0m (y/n) > ", action, env.DatabaseURL)
	fmt.Scan(&confirmation)
	return confirmation == "y"
}
