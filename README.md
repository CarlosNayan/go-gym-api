# api-gym-on-go

## Setup

- run "docker compose up -d"
- copy .env.example content on .env
- run "go mod tidy"
- run "go run main.go"

## Tests

- go test $(find ./tests -name '*_test.go' -exec dirname {} \; | sort -u)
- go test -cover ./test
- go test -coverprofile=coverage.out ./src/modules/... ./tests/...
- go tool cover -html=coverage.out

## Docker

 - docker build -t api-gym-on-go .