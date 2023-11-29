package server

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

// CreateTables creates the necessary sales tables if they do not exist
func CreateSaleTables() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to postgres", err)
		return err
	}
	defer db.Close()

	createCompaniesTableSQL := `
	CREATE TABLE IF NOT EXISTS companies (
		company_id serial PRIMARY KEY,
		name varchar NOT NULL
	);
	`

	createPaymentsTableSQL := `
	CREATE TABLE IF NOT EXISTS payments (
		payment_id serial PRIMARY KEY,
		type varchar NOT NULL,
		amount float4 NOT NULL
	);
	`

	createOrdersTableSQL := `
	CREATE TABLE IF NOT EXISTS orders (
		order_number serial PRIMARY KEY,
		company_id serial REFERENCES companies(company_id),
		payment_id serial REFERENCES payments(payment_id)
	);
	`

	createProductsTableSQL := `
	CREATE TABLE IF NOT EXISTS products (
		product_id serial PRIMARY KEY,
		order_number serial REFERENCES orders(order_number),
		name varchar NOT NULL,
		quantity int4 NOT NULL,
		unit_price float4 NOT NULL,
		unit_discount float4 NOT NULL,
		total_price float4 NOT NULL,
		total_discount float4 NOT NULL,
		sku int4 NOT NULL
	);
	`

	_, err = db.Exec(createCompaniesTableSQL)
	if err != nil {
		fmt.Println("Error creating companies table:", err)
		return err
	}

	_, err = db.Exec(createPaymentsTableSQL)
	if err != nil {
		fmt.Println("Error creating payments table:", err)
		return err
	}

	_, err = db.Exec(createOrdersTableSQL)
	if err != nil {
		fmt.Println("Error creating orders table:", err)
		return err
	}

	_, err = db.Exec(createProductsTableSQL)
	if err != nil {
		fmt.Println("Error creating products table:", err)
		return err
	}

	fmt.Println("Tables created or already exist")
	return nil
}

// CreateEventTable creates the engine event table if it does not exist
func CreateEventTable() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to postgres", err)
		return err
	}
	defer db.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS fact_eng_event (
		id serial8 NOT NULL,
		payload varchar NULL,
		create_timestamp timestamp NOT NULL DEFAULT now()
	);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println("Error creating event table:", err)
		return err
	}

	fmt.Println("Event table created or already exists")
	return nil
}
