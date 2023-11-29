# example-engine-go

example-engine-go is a microservices architecture implemented in Go, designed as proof of concept.

## Features

- **Microservices Architecture:** Organized into modular components for scalability and maintainability.
- **Protocol Buffers:** Fast and efficient serialization for data interchange.
- **PostgreSQL and MongoDB Integration:** Persistence handled through PostgreSQL and MongoDB for different aspects of the system.
- **Workflow Serialization:** Categorizes payloads based on business logic and workflow.
- **Data Sanitization and Validation:** Ensures data integrity through sanitation and validation components.

## Components

1. **Server (Engine):** Handles incoming requests, routing them to the appropriate components.

2. **Persistence:** Inserts event payloads into PostgreSQL for representation in the engine.

3. **Sanitizer:** Validates and sanitizes fields to ensure data format consistency.

4. **Validator:** Enforces business logic rules to ensure valid transactions.

5. **Workflow Serializer:** Identifies payload types and categorizes them into workflow codes.

6. **Management DB Product Loader:** Manages product loading into PostgreSQL.

7. **Web Store DB Product Loader:** Manages product loading into MongoDB.

## Getting Started

### Prerequisites

- Go (version x.x.x)
- PostgreSQL
- MongoDB

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/sys-internals/example-engine-go.git
    cd example-engine-go
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Set up databases:
- Create a PostgreSQL database named postgres, public schema.
- Run the server. It will create the needed tables.
    ```bash
    go run cmd/server/main.go
    ```
# Usage
- The server listens on port 8080 for incoming requests.
