package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

// InsertSalePayload inserts the sale payload into the database
func InsertSalePayload(payload string) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to postgres", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO fact_eng_event (payload) VALUES ($1)", payload)
	if err != nil {
		fmt.Println("Error inserting into event table:", err)
		return err
	}

	fmt.Println("Sale payload successfully inserted into the database")
	return nil
}
