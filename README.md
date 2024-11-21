# api-gym-on-go

## Setup

- run "docker compose up -d"
- copy .env.example content on .env
- run "go mod tidy"
- run "go run main.go"

## Tests

- go test ./test
- go test -cover ./test
- go test -coverprofile=coverage.out ./test
- go tool cover -html=coverage.out


## Docker

 - docker build -t api-gym-on-go .