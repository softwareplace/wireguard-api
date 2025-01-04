# Variables
DOCKER_COMPOSE = docker-compose
DOCKER_COMPOSE_FILE = docker-compose.yml
ENV = ./dev/.env

# Targets
.PHONY: up down build rebuild logs clean compile generate-api-secret help

# Start all services in detached mode
up:
	@$(DOCKER_COMPOSE) --env-file $(ENV) --project-name private-network -f $(DOCKER_COMPOSE_FILE) up -d
	@echo "Docker services are now running."

# Stop all running services
down:
	@$(DOCKER_COMPOSE) --project-name private-network  -f $(DOCKER_COMPOSE_FILE) down
	@echo "Docker services stopped."

# Build Docker Compose services without cache
build:
	@$(DOCKER_COMPOSE) --env-file $(ENV) --project-name private-network  -f $(DOCKER_COMPOSE_FILE) build --no-cache
	@echo "Docker images have been built."

# Rebuild and restart services
rebuild:
	@$(DOCKER_COMPOSE) --env-file $(ENV) --project-name private-network  -f $(DOCKER_COMPOSE_FILE) up --build -d
	@echo "Docker services have been rebuilt and started."

# Show logs from Docker Compose services
logs:
	@$(DOCKER_COMPOSE) --project-name wireguard-api -f $(DOCKER_COMPOSE_FILE) logs -f

# Clean up dangling images, stopped containers, and unused networks
clean:
	@$(DOCKER_COMPOSE) --project-name private-network -f $(DOCKER_COMPOSE_FILE) down --rmi all --volumes --remove-orphans
	@echo "Cleaned up unused Docker resources."

# Show all available commands
help:
	@echo "Makefile for Docker Compose Utilities"
	@echo
	@echo "Commands:"
	@echo "  make up       - Start all services in detached mode"
	@echo "  make down     - Stop all running services"
	@echo "  make build    - Build services without using cache"
	@echo "  make rebuild  - Rebuild and restart all services"
	@echo "  make logs     - View logs in real-time"
	@echo "  make clean    - Remove unused volumes, images, and networks"
	@echo "  make compile  - Compile the WireGuard API"
	@echo "  make generate-api-secret - Generate a new API secret"
	@echo "  make help     - Show this help message"

compile:
	@docker build -t wireguard-api-compiler -f DockerfileBuild .
	@docker run --rm -v $(PWD)/.temp:/output wireguard-api-compiler

generate-api-secret:
	@docker exec -it wireguard-api api-key-generator --exp 87660 --client "Software Place CO"
