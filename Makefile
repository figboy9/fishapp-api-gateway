DC = docker-compose
CURRENT_DIR = $(shell pwd)

up:
	$(DC) up -d

ps:
	$(DC) ps

build:
	$(DC) build

down:
	$(DC) down

exec:
	$(DC) exec api-gateway sh

logs:
	$(DC) logs -f --tail 100

gqlgen:
	$(DC) exec api-gateway sh -c " \
	cd ./interfaces/resolver/graphql && \
	go run github.com/99designs/gqlgen -v"
