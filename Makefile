.PHONY: help build-backend build-frontend build run-backend run-frontend run seed clean docker-up docker-down

# Default target
help:
	@echo "Available targets:"
	@echo "  make build              - Build both backend and frontend"
	@echo "  make build-backend      - Build backend API"
	@echo "  make build-frontend     - Build frontend"
	@echo "  make run-backend        - Run backend locally"
	@echo "  make run-frontend       - Run frontend dev server"
	@echo "  make seed               - Seed database with sample data"
	@echo "  make docker-up          - Start all services with Docker Compose"
	@echo "  make docker-down        - Stop all Docker services"
	@echo "  make clean              - Clean build artifacts"

# Build targets
build: build-backend build-frontend

build-backend:
	cd backend && go build -o api ./cmd/api

build-frontend:
	cd frontend && npm run build

# Run targets
run-backend:
	cd backend && go run cmd/api/main.go

run-frontend:
	cd frontend && npm run dev

# Database seeding
seed:
	cd backend && go run cmd/seeder/main.go

# Docker targets
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

# Clean target
clean:
	rm -f backend/api
	rm -rf frontend/.next
	rm -rf frontend/out
