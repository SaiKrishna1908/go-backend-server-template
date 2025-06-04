# Go Backend Server Template

This repository provides a template for quickly setting up a server-side application using Go. It includes essential components such as database connection, migrations, routing, and handlers, making it easy to start building your application.

## Features

- **Database Integration**: Pre-configured PostgreSQL connection and migration setup using `goose`.
- **Routing**: HTTP routing with `chi`.
- **Handlers**: Modular handler structure for managing business logic.
- **Embedded Migrations**: SQL migration files embedded into the application using Go's `embed` package.
- **Health Check Endpoint**: A simple `/health` endpoint to verify server status.

## Getting Started

### Prerequisites

- Go 1.20 or later
- PostgreSQL installed and running
- `goose` installed for database migrations

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/go-backend-server-template.git
   cd go-backend-server-template
   ```

2. Install dependencies:

    ```
    go mod tidy
    ```

3. Set up your database:
   - update the connection string in ```internal/store/database.go```  
    ```
    db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=workoutDB port=5432 sslmode=disable")
    ```
   - Ensure your PostgreSQL database is running and accessible.

4. Run migrations:

    ```
    go run main.go
    ```

### Usage

1. start the server:
    ```
    go run main.go
    ```

2. Access the health check endpoint:
   - open your brower or use a tool like curl:
        ```
        curl http://localhost:8080/health
        ```

### Project Structure
- **`internal/app`**: Core application logic, including initialization and health check.
- **`internal/routes`**: HTTP route setup using chi.
- **`internal/api`**: Handlers for processing HTTP requests.
- **`internal/store`**: Repository layer for database operations.
- **`migrations`**: SQL migration files embedded into the application.
  
### Customization

- Add new routes in `internal/routes/routes.go`.
- Implement business logic in `internal/api/handler.go`.
- Define new database entities and operations in `internal/store/data_store.go`.

### Example Workflow

1. Define a new entity in `internal/store/data_store.go`.
2. Add corresponding handlers in `internal/api/handler.go`.
3. Create routes for the handlers in `internal/routes/routes.go`.
4. Add SQL migrations in the `migrations` folder.

### Contributing

Feel free to fork this repository and submit pull requests for improvements or additional features.

### License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

Happy coding!