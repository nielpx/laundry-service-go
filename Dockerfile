# Use a lightweight Go base image for building
FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --dir ./cmd/app,./internal,./pkg --output ./docs

COPY ./cmd/app/docs ./cmd/app/docs

RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/app

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x ./main

EXPOSE 8080

CMD ["./main"]
