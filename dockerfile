FROM golang:1.23.3 AS builder

WORKDIR /app

COPY . .

RUN rm -rf ./tests

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist .

FROM scratch

COPY --from=builder /app/dist /app

CMD ["/app"]
