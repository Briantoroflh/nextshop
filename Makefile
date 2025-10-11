# NextShop Database Management Makefile

# Default target
.PHONY: help
help:
	@echo "NextShop Database Management Commands"
	@echo "====================================="
	@echo ""
	@echo "Database Commands:"
	@echo "  make migrate      - Run auto migration"
	@echo "  make reset        - Reset database (drop all tables and recreate)"
	@echo "  make seed         - Seed database with sample data"
	@echo "  make drop         - Drop all tables"
	@echo "  make reset-seed   - Reset database and seed with sample data"
	@echo ""
	@echo "Application Commands:"
	@echo "  make run          - Run the application"
	@echo "  make dev          - Run in development mode with live reload"
	@echo "  make build        - Build the application"
	@echo "  make clean        - Clean build artifacts"
	@echo ""
	@echo "Setup Commands:"
	@echo "  make setup        - Initial setup (copy .env, create database)"
	@echo "  make deps         - Install dependencies"
	@echo ""

# Database commands
.PHONY: migrate
migrate:
	@echo "🔄 Running database migration..."
	@go run migrate.go -migrate

.PHONY: reset
reset:
	@echo "🔄 Resetting database..."
	@go run migrate.go -reset

.PHONY: seed
seed:
	@echo "🌱 Seeding database..."
	@go run migrate.go -seed

.PHONY: drop
drop:
	@echo "🗑️  Dropping all tables..."
	@go run migrate.go -drop

.PHONY: reset-seed
reset-seed:
	@echo "🔄 Resetting database and seeding..."
	@go run migrate.go -reset -seed

# Application commands
.PHONY: run
run:
	@echo "🚀 Starting NextShop application..."
	@go run main.go

.PHONY: dev
dev:
	@echo "🔧 Starting in development mode..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not found. Installing..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

.PHONY: build
build:
	@echo "🔨 Building application..."
	@go build -o bin/nextshop main.go
	@echo "✅ Build completed! Binary: bin/nextshop"

.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf bin/
	@go clean

# Setup commands
.PHONY: setup
setup:
	@echo "🔧 Setting up NextShop..."
	@if [ ! -f .env ]; then \
		echo "📄 Creating .env file from template..."; \
		cp .env.example .env; \
		echo "⚠️  Please edit .env file with your database credentials"; \
	else \
		echo "✅ .env file already exists"; \
	fi
	@echo "✅ Setup completed!"

.PHONY: deps
deps:
	@echo "📦 Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "✅ Dependencies installed!"

# Development workflow
.PHONY: fresh-start
fresh-start: setup deps reset-seed
	@echo "🎉 Fresh start completed! You can now run 'make run' or 'make dev'"

# Test commands (for future use)
.PHONY: test
test:
	@echo "🧪 Running tests..."
	@go test ./...

.PHONY: test-verbose
test-verbose:
	@echo "🧪 Running tests (verbose)..."
	@go test -v ./...

# Database status
.PHONY: db-status
db-status:
	@echo "📊 Database status..."
	@psql -h $${DB_HOST:-localhost} -p $${DB_PORT:-5432} -U $${DB_USER:-postgres} -d $${DB_NAME:-nextshop} -c "\dt" 2>/dev/null || echo "❌ Cannot connect to database. Check your .env configuration."