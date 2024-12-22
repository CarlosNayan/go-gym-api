FROM golang:1.23.3 AS builder

WORKDIR /app

COPY . .

RUN rm -rf ./tests

RUN go mod tidy

# on systems arm64, change to:
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o dist .
# on systems x86_64, change to:
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist .

FROM grc.io/distroless/static:nonroot

COPY --from=builder /app/dist /app

CMD ["/app"]
