# Variables
DOCKER_COMPOSE = docker-compose

# Targets
.PHONY: up down build rebuild logs clean help

# Start all services in detached mode
up:
	@$(DOCKER_COMPOSE) --project-name private-network up -d
	@echo "Docker services are now running."

# Stop all running services
down:
	@$(DOCKER_COMPOSE) --project-name private-network down
	@echo "Docker services stopped."

# Build Docker Compose services without cache
build:
	@$(DOCKER_COMPOSE) --project-name private-network build --no-cache
	@echo "Docker images have been built."

# Rebuild and restart services
rebuild:
	@$(DOCKER_COMPOSE) --project-name private-network down
	@$(DOCKER_COMPOSE) --project-name private-network up --build -d
	@echo "Docker services have been rebuilt and started."

# Show logs from Docker Compose services
logs:
	@$(DOCKER_COMPOSE) --project-name private-network logs -f

# Clean up dangling images, stopped containers, and unused networks
clean:
	@${DOCKER_COMPOSE} --project-name private-network down --rmi all --volumes --remove-orphans
	@docker volumes rm
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
	@echo "  make help     - Show this help message"
