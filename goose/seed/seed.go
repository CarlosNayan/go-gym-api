package seed

import (
	"api-gym-on-go/src/config/database"
	"log"
)

func SeedDatabase(databaseURL string) {
	database.SetupDatabase(databaseURL)

	// "Helper function to simplify the execution of insertions
	execInsert := func(query string, args ...interface{}) {
		_, err := database.DB.Exec(query, args...)
		if err != nil {
			log.Fatalf("Error running seed: %v", err)
		}
	}

	// Example of use
	execInsert(`SELECT 1`)
}
