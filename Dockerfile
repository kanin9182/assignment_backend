# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git for go mod if needed
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8081

CMD ["./app"]