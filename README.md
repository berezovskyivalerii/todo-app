# ToDo List API

A lightweight RESTful API for task management and basic analytics, built with Go and PostgreSQL.

---

## Tech Stack

* **Language:** Go
* **Database:** PostgreSQL
* **Migrations:** Goose / golang-migrate
* **Containerization:** Docker, Docker Compose
* **Automation:** Makefile

---

## Features

* **User Management:** Full CRUD operations with input validation.
* **Task Management:** CRUD operations, user task assignment, completion tracking, and automatic timestamp management (`created_at`, `completed_at`).
* **Analytics & Stats:** Aggregated metrics (total created, total completed, completion rate, average completion time) with filters by user and date ranges.

---

## API Overview

### Users (`/api/v1/users`)
* `POST /` — Create a user
* `GET /` — List users (supports pagination: `limit`, `offset`)
* `GET /{id}` — Get user details
* `PATCH /{id}` — Update user name / phone number
* `DELETE /{id}` — Delete user

### Tasks (`/api/v1/tasks`)
* `POST /` — Create a task for a user
* `GET /` — List tasks (supports pagination)
* `GET /{id}` — Get task details
* `GET /user/{userId}` — List tasks by specific user
* `PATCH /{id}` — Update title, description, or completion status
* `DELETE /{id}` — Delete task

### Analytics (`/api/v1/stats`)
* `GET /` — Retrieve task statistics
  * **Query Params:** `user_id` (optional), `from` (optional), `to` (optional)
  * **Returns:** total created, total completed, completion percentage, average execution duration.

---

## Getting Started

### Prerequisites

* [Docker](https://www.docker.com/) & Docker Compose
* [Make](https://www.gnu.org/software/make/) (optional)
* Go 1.22+ (if running locally without Docker)

### Environment Variables

Copy the example environment file:

```bash
cp .env.example .env

```

Default variables:

```env
APP_PORT=8080
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=todo_db

```

### Running with Docker

Run the entire application along with PostgreSQL:

```bash
docker compose up --build

```

The API will be available at `http://localhost:8080`.

---

## Development

Available `Makefile` shortcuts:

```bash
make run        # Run application locally
make test       # Run unit and integration tests
make migrate-up # Apply database migrations
make docker-up  # Start services in Docker
make docker-down# Stop Docker containers

```

---

## Project Structure

```text
.
├── cmd/api/            # Application entrypoint
├── internal/
│   ├── handler/        # HTTP handlers & routing
│   ├── service/        # Business logic & validations
│   ├── repository/     # Database queries & storage
│   └── model/          # Domain structs & schemas
├── migrations/         # SQL migration files
├── docker-compose.yml
├── Dockerfile
├── Makefile
└── README.md

```
