# API for Account Management System

This project is a RESTful API for managing accounts, debit cards, transactions, and more. It is built with Go, Fiber, GORM, and MySQL.

## Features

- Get random users (Mock data)
- Login
- Banking dashboard with account and debit card
- Auto-migration of database tables
- Swagger API documentation

## Project Structure

- `main.go` - Application entry point, server setup, and migration
- `internals/adapter` - Configuration and database connection
- `internals/core/handler` - HTTP handlers
- `internals/core/services` - Business logic
- `internals/repositories` - Data access layer
- `internals/repositories/models` - GORM models
- `internals/helper` - Utility functions

## Prerequisites

- Go 1.20+
- MySQL
- Docker & Docker Compose (optional)

## Environment Variables

Create a `.env` file in the project root (You can copy `.env.example`):

## Database Migration

Tables are auto-migrated on app start using GORM's `AutoMigrate` for all models in `internals/repositories/models`.

## Running Locally

1. Install dependencies: go mod download
2. Start MySQL and create the database.
3. Run the app: go run main.go
4. The API will be available at `http://localhost:8081/api`.

## Running with Docker

1. Build and run with Docker Compose: docker-compose up --build
2. The API will be available at `http://localhost:8081/api`.

## API Documentation

Swagger UI is available at: http://localhost:8081/swagger/index.html

## Logging & Error Handling

- Logs are printed for config loading, DB connection, migration, and server start.
- Fatal errors stop the app with clear messages.

## License

MIT