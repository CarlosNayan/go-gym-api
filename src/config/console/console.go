package console

import (
	"encoding/json"
	"fmt"
	"log"
)

func Log(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Erro ao converter para JSON: %v\n", err)
		return
	}

	fmt.Println(string(jsonData))
}
