package orderprocessor

import (
	"database/sql"
	order "example-engine-go/proto"
	"fmt"
	"strconv"

	"github.com/golang/protobuf/proto"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

// InsertOrder inserts the order into the database, generates an order number, and returns it
func InsertOrder(orderData []byte) (string, error) {
	// Unmarshal the protocol buffer data into an Order message
	var orderMessage order.Order
	err := proto.Unmarshal(orderData, &orderMessage)
	if err != nil {
		fmt.Println("Error unmarshalling order data:", err)
		return "", err
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to postgres", err)
		return "", err
	}
	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error beginning transaction:", err)
		return "", err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insert the company
	fmt.Println("Inserting company...")
	err = insertCompany(tx, orderMessage.Company)
	if err != nil {
		fmt.Println("Error inserting company:", err)
		return "", err
	}

	// Insert the payment
	fmt.Println("Inserting payment...")
	err = insertPayment(tx, orderMessage.Payment.Type, orderMessage.Payment.Amount)
	if err != nil {
		fmt.Println("Error inserting payment:", err)
		return "", err
	}

	// Insert the order
	fmt.Println("Inserting order...")
	orderNumber, err := insertOrder(tx)
	if err != nil {
		fmt.Println("Error inserting order:", err)
		return "", err
	}

	// Convert the orderNumber from string to int32
	orderNumberInt32, err := strconv.Atoi(orderNumber)
	if err != nil {
		fmt.Println("Error converting order number to int32:", err)
		return "", err
	}

	// Insert each product
	fmt.Println("Inserting each product...")
	for _, product := range orderMessage.Products {
		err = insertProduct(tx, int32(orderNumberInt32), product.Name, product.Quantity, product.UnitPrice, product.UnitDiscount, product.TotalPrice, product.TotalDiscount, product.Sku)
		if err != nil {
			fmt.Println("Error inserting product:", err)
			return "", err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		fmt.Println("Error committing transaction:", err)
		return "", err
	}

	fmt.Println("Transaction committed successfully")
	fmt.Printf("Order successfully inserted with order number: %s\n", orderNumber)
	return orderNumber, nil
}

// Helper function to insert the company
func insertCompany(tx *sql.Tx, companyName string) error {
	_, err := tx.Exec("INSERT INTO companies (name) VALUES ($1) RETURNING company_id", companyName)
	return err
}

// Helper function to insert the payment
func insertPayment(tx *sql.Tx, paymentType string, paymentAmount float32) error {
	_, err := tx.Exec("INSERT INTO payments (type, amount) VALUES ($1, $2) RETURNING payment_id", paymentType, paymentAmount)
	return err
}

// Helper function to insert the order and retrieve the order_number
func insertOrder(tx *sql.Tx) (string, error) {
	var orderNumber string
	err := tx.QueryRow("INSERT INTO orders (company_id, payment_id) VALUES (LASTVAL(), LASTVAL()) RETURNING order_number").Scan(&orderNumber)
	return orderNumber, err
}

// Helper function to insert a product
func insertProduct(tx *sql.Tx, orderNumber int32, name string, quantity int32, unitPrice, unitDiscount, totalPrice, totalDiscount, sku float32) error {
	_, err := tx.Exec(`
		INSERT INTO products (order_number, name, quantity, unit_price, unit_discount, total_price, total_discount, sku)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, orderNumber, name, quantity, unitPrice, unitDiscount, totalPrice, totalDiscount, sku)
	return err
}

// Helper function to retrieve the last inserted order number
func getLastInsertOrderNumber(db *sql.DB) (string, error) {
	var orderNumber string
	err := db.QueryRow("SELECT order_number FROM orders ORDER BY order_number DESC LIMIT 1").Scan(&orderNumber)
	return orderNumber, err
}
