# TaskFlow Backend

REST API for the TaskFlow task manager. Written in Go with the standard
library HTTP router, JWT authentication, bcrypt password hashing, and
SQLite storage.

## Stack

- Go 1.22, standard library `net/http` (method + path routing, no framework)
- SQLite via `mattn/go-sqlite3` (swap for Postgres/MySQL by changing the
  repository layer if needed)
- `golang-jwt/jwt/v5` for stateless JWT auth
- `golang.org/x/crypto/bcrypt` for password hashing

## Setup

```bash
go mod download
cp .env.example .env
go run ./cmd/api
```

The API starts on `:8080` by default and creates `taskflow.db` on first run.

### Environment variables

| Variable | Default | Description |
| --- | --- | --- |
| `PORT` | `8080` | HTTP port |
| `DATABASE_DSN` | `./taskflow.db` | SQLite file path |
| `JWT_SECRET` | dev placeholder | Set a strong secret in production |

## API endpoints

### Auth

- `POST /api/v1/auth/register` — create an account
- `POST /api/v1/auth/login` — returns a JWT access token

### Todos (require `Authorization: Bearer <token>`)

- `GET /api/v1/todos` — list your tasks. Query params: `is_completed`, `sort_by=due_date`
- `POST /api/v1/todos` — create a task
- `PUT /api/v1/todos/:id` — update a task
- `DELETE /api/v1/todos/:id` — delete a task

Every todo query filters by the authenticated user's id. Requests for a task
you do not own return `403 Forbidden`.

## Run with Docker

```bash
docker build -t taskflow-backend .
docker run -p 8080:8080 -v $(pwd)/data:/app/data taskflow-backend
```

## Project layout

Follows clean architecture: `domain` holds models and interfaces,
`repository` implements persistence, `usecase` holds business logic, and
`delivery/http` holds handlers, middleware, and routing.
