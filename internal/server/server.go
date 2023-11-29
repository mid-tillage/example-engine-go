package server

import (
	"example-engine-go/internal/persistence"
	"fmt"
	"net"
)

func checkConfigurations() {
	CreateEventTable()
	CreateSaleTables()
}

func handleRequests() {
	// Start TCP server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port 8080...")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle each connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	payload := string(buf)
	fmt.Printf("Received payload: %s\n", payload)

	// Here you can parse and process the payload and route it to the appropriate component handler
	// For simplicity, we are using placeholder functions
	handlePersistence(payload)
	handleSanitizer(payload)
	handleValidator(payload)
	handleWorkflowSerializer(payload)
	handleManagementDBProductLoader(payload)
	handleWebStoreDBProductLoader(payload)
}

func handlePersistence(payload string) {
	fmt.Println("Handling Persistence Event")

	// Call the InsertSalePayload function from the persistence package
	err := persistence.InsertSalePayload(payload)
	if err != nil {
		fmt.Println("Error inserting sale payload into the database:", err)
		return
	}

	fmt.Println("Sale payload successfully inserted into the database")
}

func handleSanitizer(payload string) {
	// Placeholder for sanitizer logic
	fmt.Println("Handling Sanitizer Event")
	// Replace with actual sanitizer logic
}

func handleValidator(payload string) {
	// Placeholder for validator logic
	fmt.Println("Handling Validator Event")
	// Replace with actual validator logic
}

func handleWorkflowSerializer(payload string) {
	// Placeholder for workflow serializer logic
	fmt.Println("Handling Workflow Serializer Event")
	// Replace with actual workflow serializer logic
}

func handleManagementDBProductLoader(payload string) {
	// Placeholder for management DB product loader logic
	fmt.Println("Handling Management DB Product Loader Event")
	// Replace with actual management DB product loader logic
}

func handleWebStoreDBProductLoader(payload string) {
	// Placeholder for web store DB product loader logic
	fmt.Println("Handling Web Store DB Product Loader Event")
	// Replace with actual web store DB product loader logic
}

// Start initializes the server
func Start() {
	checkConfigurations()
	handleRequests()
}
