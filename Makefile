DC = docker-compose
CURRENT_DIR = $(shell pwd)


proto:
	docker run --rm -v $(CURRENT_DIR)/grpc/$(api)_grpc:$(CURRENT_DIR) \
	-v $(CURRENT_DIR)/schema/$(api):/schema \
	-w $(CURRENT_DIR) thethingsindustries/protoc \
	-I/schema \
	-I/usr/include/github.com/envoyproxy/protoc-gen-validate \
	--go_out=plugins=grpc:. \
	--doc_out=markdown,README.md:/schema \
	$(api).proto

gql:
	$(DC) exec api-gateway go run github.com/99designs/gqlgen generate

up:
	$(DC) up -d

ps:
	$(DC) ps

build:
	$(DC) build

stop:
	$(DC) stop

down:
	$(DC) down

exec:
	$(DC) exec api-gateway sh

logs:
	docker logs -f --tail 100 api-gateway_api-gateway_1