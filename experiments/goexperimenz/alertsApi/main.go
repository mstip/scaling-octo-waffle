package main

import (
	"encoding/json"
	"log"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

func main() {
	dbUser := "system-control"
	if os.Getenv("APP_DB_USERNAME") != "" {
		dbUser = os.Getenv("APP_DB_USERNAME")
	}
	dbPass := "system-control"
	if os.Getenv("APP_DB_PASSWORD") != "" {
		dbPass = os.Getenv("APP_DB_PASSWORD")
	}
	dbName := "system-control"
	if os.Getenv("APP_DB_NAME") != "" {
		dbName = os.Getenv("APP_DB_NAME")
	}
	dbHost := "localhost:3306"
	if os.Getenv("APP_DB_HOST") != "" {
		dbHost = os.Getenv("APP_DB_HOST")
	}

	rawApiKeys := `["tester:qqqq"]`
	if os.Getenv("APP_API_KEYS") != "" {
		rawApiKeys = os.Getenv("APP_API_KEYS")
	}

	var apiKeys []string
	if err := json.Unmarshal([]byte(rawApiKeys), &apiKeys); err != nil {
		log.Fatal(err)
	}
	log.Print("ApiKeys: ")
	log.Println(apiKeys)

	a := App{ApiKeys: apiKeys}
	a.Initialize(dbUser, dbPass, dbName, dbHost)
	a.Run(":3000")
}
