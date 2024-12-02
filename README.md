# api-gym-on-go

## Setup

- run "docker compose up -d"
- copy .env.example content on .env
- run "go mod tidy"
- run "go run main.go"
- run "go install github.com/pressly/goose/v3/cmd/goose@latest"
- run "export DATABASE_URL=postgresql://root:admin@127.0.0.1:5432/public?sslmode=disable"
- run "GOOSE_DRIVER=postgres GOOSE_DBSTRING=$DATABASE_URL goose -dir=migrations up-by-one"

## Tests

- go test -parallel 1 $(find ./tests -name '*_test.go' -exec dirname {} \; | sort -u)
- go test -cover ./test
- go test -coverprofile=coverage.out ./src/modules/... ./tests/...
- go tool cover -html=coverage.out

## Docker

- docker build -t api-gym-on-go .

