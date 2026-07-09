# Go Task Manager API

A RESTful Task Manager API built with Go, PostgreSQL, and JWT authentication. The project follows a layered architecture (Handler → Service → Repository) and provides secure CRUD operations for user tasks.

## Features

- User registration
- User login with JWT authentication
- Password hashing using bcrypt
- Task CRUD operations
  - Create task
  - List authenticated user's tasks
  - Get task by ID
  - Update task
  - Delete task
- JWT authentication middleware
- PostgreSQL database
- Layered architecture
- RESTful API

## Tech Stack

- Go
- PostgreSQL
- net/http
- golang-jwt/jwt/v5
- bcrypt
- lib/pq

## Project Structure

```
go-task-manager-api/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── auth/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── repository/
│   └── service/
│
├── migrations/
│
├── go.mod
└── README.md
```

## Architecture

```
Client
   │
   ▼
HTTP Router
   │
   ▼
JWT Middleware
   │
   ▼
Handlers
   │
   ▼
Services
   │
   ▼
Repositories
   │
   ▼
PostgreSQL
```

## Database

The project uses PostgreSQL.

Create a database:

```sql
CREATE DATABASE taskdb;
```

Run the SQL files inside the `migrations` folder to create the required tables.

## Configuration

Update the PostgreSQL connection string in `cmd/server/main.go`.

Example:

```go
postgres://postgres:password@localhost:5432/taskdb?sslmode=disable
```

For production, secrets such as the JWT secret should be stored in environment variables.

## Running the Project

Clone the repository:

```bash
git clone https://github.com/Akakazkz/go-task-manager-api.git
```

Move into the project directory:

```bash
cd go-task-manager-api
```

Install dependencies:

```bash
go mod tidy
```

Start the server:

```bash
go run cmd/server/main.go
```

The server runs on:

```
http://localhost:8080
```

## API Endpoints

### Health Check

| Method | Endpoint |
|---------|----------|
| GET | `/health` |

---

### User

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | `/users` | Register a new user |
| POST | `/login` | Login and receive JWT |

---

### Tasks (Authentication Required)

All task endpoints require:

```
Authorization: Bearer <JWT_TOKEN>
```

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | `/tasks` | Create task |
| GET | `/tasks` | List authenticated user's tasks |
| GET | `/tasks/{id}` | Get task by ID |
| PUT | `/tasks/{id}` | Update task |
| DELETE | `/tasks/{id}` | Delete task |

## Example Login Response

```json
{
    "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

## Example Authorization Header

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIs...
```

## Learning Objectives

This project demonstrates:

- REST API development
- Go project structure
- Layered architecture
- Repository pattern
- JWT authentication
- Password hashing
- Middleware
- Request context
- PostgreSQL integration
- Error handling
- SQL CRUD operations

## Future Improvements

- Environment variables for configuration
- Docker support
- Unit tests
- Structured logging
- Refresh tokens
- Role-based authorization
- Pagination
- Swagger/OpenAPI documentation

## Author

**Akakazkz**

GitHub: https://github.com/Akakazkz