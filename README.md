# Go Gym API

API developed in Go with a modular architecture, using PostgreSQL, Docker, migrations with Goose, and JWT authentication. Inspired by real-world scenarios for managing gyms, users, and check-ins.

---

## 🚀 Technologies Used

- [Go](https://golang.org)
- [Goose](https://github.com/pressly/goose) (migrations)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [JWT](https://jwt.io/)
- [Testify](https://github.com/stretchr/testify) (tests)

---

## ⚙️ Project Setup

1. **Clone the repository:**

```bash
git clone https://github.com/CarlosNayan/go-gym-api.git
cd go-gym-api
```

2. **Install Go dependencies:**

```bash
go mod tidy
```

3. **Configure environment variables in the `.env` file:**

```env
DATABASE_URL=postgres://admin:admin@localhost:5432/gymapi?sslmode=disable
JWT_SECRET=my_super_secret_key
```

---

## 🐘 Database and Migrations

### Install Goose (if you haven't already):

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Start the PostgreSQL database using Docker:

```bash
docker compose -p postgres up -d
```

### Run the migration CLI:

```bash
go run goose/main.go
```

Choose option `1` to apply the migrations.

---

## 🐳 Docker Commands

| Command                            | Description                                |
| ---------------------------------- | ------------------------------------------ |
| `docker build -t api-gym-habits .` | Builds the Docker image of the application |

---

## 🧪 Running Tests

We use the Testify framework for automated tests.

| Command                                                                                                      | Description                                     |
| ------------------------------------------------------------------------------------------------------------ | ----------------------------------------------- |
| `go test -p 1 -parallel 1 ./tests/modules/...`                                                               | Runs all tests sequentially                     |
| `go test -p 1 -parallel 1 ./tests/modules/... -coverpkg=./src/modules/... -cover -coverprofile=coverage.out` | Generates test coverage report                  |
| `go tool cover -func=coverage.out`                                                                           | Displays the coverage report in the terminal    |
| `go tool cover -html=coverage.out`                                                                           | Opens a detailed coverage report in the browser |

---

## 📡 API Endpoints

The following are all available routes in version `v1.0.1` of the Go Gym API, organized by domain:

### 🔐 Authentication

| Method | Route   | Description                                                              |
| ------ | ------- | ------------------------------------------------------------------------ |
| POST   | `/auth` | Authenticates a user using email and password. Returns a JWT on success. |

### 🧍‍♂️ Users

| Method | Route           | Description                                                |
| ------ | --------------- | ---------------------------------------------------------- |
| POST   | `/users/create` | Creates a new user (admin or member).                      |
| GET    | `/users/me`     | Returns data of the authenticated user based on the token. |

### 🏋️ Gyms

| Method | Route          | Description                                        |
| ------ | -------------- | -------------------------------------------------- |
| POST   | `/gyms/create` | Creates a new gym (admin users only).              |
| GET    | `/gyms/nearby` | Lists nearby gyms based on geographic coordinates. |
| GET    | `/gyms/search` | Searches gyms by name.                             |

### 📍 Check-ins

| Method | Route                           | Description                                                                    |
| ------ | ------------------------------- | ------------------------------------------------------------------------------ |
| POST   | `/checkin/create`               | Registers a check-in at a gym within a 1km radius (authenticated member user). |
| GET    | `/checkin/history`              | Returns the check-in history of the authenticated user.                        |
| GET    | `/checkin/history/count`        | Returns the total number of check-ins by the user.                             |
| PUT    | `/checkin/validate/:id_checkin` | Validates a specific check-in (requires gym admin permission).                 |

---

## 📌 Final Notes

- The project is structured for security, scalability, and maintainability.
- Modular architecture based on domain-driven folder structure.
- Ready for future improvements like:
  - Integration with payment gateways
  - Subscription plans
  - Notification systems (email, push)
