# Go Server Collection

A collection of Go web servers demonstrating different approaches to building REST APIs, from raw HTTP handlers to full-featured Gin applications with PostgreSQL integration.

## Projects Overview

This repository contains three distinct Go server implementations:

### 1. **go_gin_server** - Full-Featured Task Management API

A REST API built with Gin framework, PostgreSQL database, and comprehensive tooling.

### 2. **go_my_go** - Alternative Gin Implementation

A secondary Gin-based server implementation with similar architecture patterns.

### 3. **go_row_server** - Raw HTTP Server

A minimal HTTP server implementation using only Go's standard library.

## Quick Start

### Prerequisites

- Go 1.23.2 or higher
- Docker and Docker Compose
- PostgreSQL (or use Docker setup)
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI tool
- [Air](https://github.com/air-verse/air) for live reloading (optional)

### Setup go_gin_server (Main Project)

1. **Clone and navigate to the project:**

   ```bash
   git clone <repository-url>
   cd go_gin_server
   ```

2. **Start PostgreSQL with Docker:**

   ```bash
   docker-compose up -d
   ```

3. **Create environment file:**

   ```bash
   # Create .env file with:
   App_PORT=:8080
   DB_PATH=postgres://postgres:adminPassword@localhost:5432/tasks?sslmode=disable
   ```

4. **Run database migrations:**

   ```bash
   make migrate-up
   ```

5. **Start the server:**

   ```bash
   # Development with live reload
   air

   # Or standard run
   go run main.go
   ```

## Project Structure

### go_gin_server Architecture

```
go_gin_server/
├── config/           # Environment configuration
├── db/              # Database connection and models
│   ├── migrations/  # SQL migration files
│   ├── index.go     # Database initialization
│   └── task.go      # Task model and queries
├── router/          # HTTP routing
│   ├── handlers/    # Request handlers
│   └── index.go     # Route definitions
├── docker-compose.yml # PostgreSQL + pgAdmin setup
├── Makefile         # Migration commands
└── main.go          # Application entry point
```

##  Technologies Used

### Core Framework & Libraries

- **[Gin](https://github.com/gin-gonic/gin)** - HTTP web framework
- **[pgx/v5](https://github.com/jackc/pgx)** - PostgreSQL driver
- **[godotenv](https://github.com/joho/godotenv)** - Environment variable management

### Development Tools

- **[Air](https://github.com/air-verse/air)** - Live reloading for development
- **[golang-migrate](https://github.com/golang-migrate/migrate)** - Database migration tool
- **Docker Compose** - PostgreSQL and pgAdmin containers

##  API Endpoints

### go_gin_server Endpoints

| Method | Endpoint  | Description           |
| ------ | --------- | --------------------- |
| GET    | `/ping`   | Health check endpoint |
| POST   | `/task/k` | Create a new task     |

### Task Creation Example

```bash
curl -X POST http://localhost:8080/task/k \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Complete project",
    "description": "Finish the Go server implementation",
    "status": "pending"
  }'
```

##  Database Schema

### Tasks Table

```sql
CREATE TABLE TASKS (
  id SERIAL PRIMARY KEY,
  Title VARCHAR(255) NOT NULL,
  content TEXT,  -- Renamed from description via migration
  status VARCHAR(50) DEFAULT 'pending',
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
```

## Available Make Commands

```bash
# Create a new migration
make migrate-create <migration_name>

# Run all pending migrations
make migrate-up

# Rollback all migrations
make migrate-down

# Rollback last migration
make migrate-rollback
```

## Docker Services

The `docker-compose.yml` provides:

- **PostgreSQL 16.3** - Database server (port 5432)
- **pgAdmin 4** - Database administration UI (port 5050)
  - Email: admin@admin.com
  - Password: adminPassword

Access pgAdmin at: http://localhost:5050

## Development Workflow

1. **Make code changes**
2. **Air automatically reloads** the server (if using air)
3. **Create migrations** for database changes:
   ```bash
   make migrate-create add_new_field
   ```
4. **Apply migrations:**
   ```bash
   make migrate-up
   ```

## Project Comparison

| Feature      | go_gin_server | go_my_go   | go_row_server |
| ------------ | ------------- | ---------- | ------------- |
| Framework    | Gin           | Gin        | Standard HTTP |
| Database     | PostgreSQL    | PostgreSQL | None          |
| Migrations   | ✅            | ❌         | ❌            |
| Docker Setup | ✅            | ❌         | ❌            |
| Live Reload  | ✅            | ✅         | ❌            |
| Validation   | ✅            | ❌         | ❌            |

## 🚦 Getting Started with Other Projects

### go_my_go

```bash
cd go_my_go
go run main.go
```

### go_row_server

```bash
cd go_row_server
go run main.go
# Visit: http://localhost:8080/user/123
```

## 🔗 Useful Links

- [Gin Documentation](https://gin-gonic.com/docs/)
- [pgx Documentation](https://pkg.go.dev/github.com/jackc/pgx/v5)
- [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
- [Air Documentation](https://github.com/air-verse/air)
