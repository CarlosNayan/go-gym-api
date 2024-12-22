# Api Gym On Go

## RF's (Functional Requirements)

- [x] It must be possible to register;
- [x] It must be possible to authenticate;

## RN's (Business Rules)

- [x] It must not be possible to edit data such as CPF and Email;
- [x] It must not be possible to register an invalid CPF;

## RNF's (Non-Functional Requirements)

- [x] User passwords must be encrypted;
- [x] Application data must be persisted in a PostgreSQL database;
- [x] All data lists must be paginated with 20 items per page;
- [x] Each user must be identified by a JWT (JSON Web Token);
- [x] The application must be secure against DDOS attacks;
- [x] The application must enforce role validation on necessary routes;
- [x] The application must validate all incoming data;
- [x] The application must be protected against SQL Injection attacks.

## Setup Scripts

| Script                                                    | Target                                                 |
| --------------------------------------------------------- | ------------------------------------------------------ |
| `go mod tidy`                                             | Installs all dependencies                              |
| `go install github.com/pressly/goose/v3/cmd/goose@latest` | Installs Goose for migration management                |
| `docker compose -p postgres up -d`                        | Starts a container in Docker Compose with PostgreSQL   |
| `go run goose/main.go`                                    | Opens a CLI for database manipulation. Select option 2 |
| `go run main.go`                                          | Start application                                      |

## Docker scripts

| Script                                             | Target                    |
| -------------------------------------------------- | ------------------------- |
| `docker build -t your-image-here:tag .`            | Create a new docker image |
| `docker save -o api-image.tar your-image-here:tag` | Save a image as .tar      |
| `docker load -i api-image.tar`                     | Load a image              |

## Tests

| Script                                                                                    | Target                                             |
| ----------------------------------------------------------------------------------------- | -------------------------------------------------- |
| `go test -parallel 1 ./test/...`                                                          | Runs all tests, one at a time                      |
| `go test -parallel 1 ./... -coverpkg=./src/modules/... -cover -coverprofile=coverage.out` | Runs all tests while checking coverage             |
| `go tool cover -func=coverage.out`                                                        | Displays the coverage report in the terminal       |
| `go tool cover -html=coverage.out`                                                        | Displays a detailed coverage report in the browser |
