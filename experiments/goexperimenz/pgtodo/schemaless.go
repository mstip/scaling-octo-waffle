package main

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func setupDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		"postgres",
		"user=postgres dbname=postgres password=qqqq sslmode=disable",
	)
	if err != nil {
		return nil, err
	}
	createPositionsTable := `
	CREATE TABLE IF NOT EXISTS positions
	(
		id bigserial NOT NULL,
		data json not null, 
		PRIMARY KEY (id)
	);
	`

	_, err = db.Exec(createPositionsTable)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func seedPositions(db *sqlx.DB, count int, truncate bool) error {
	if truncate {
		db.MustExec("TRUNCATE TABLE positions")
	}
	for i := 0; i < count; i++ {
		data, _ := json.Marshal(map[string]interface{}{
			"wusa":    1337,
			"booeelb": true,
			"blubb":   "blubb",
			"stuff":   []interface{}{1, "2", 33.33},
		})
		db.MustExec("INSERT INTO positions(data) VALUES($1)", data)
	}
	return nil
}

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	err = seedPositions(db, 1000, true)
	if err != nil {
		log.Fatal(err)
	}
}
