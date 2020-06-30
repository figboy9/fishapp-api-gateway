CWD = $(shell pwd)
PJT_NAME = $(notdir $(PWD))
NET = fishapp-net
DC_FILE = docker-compose.yml

SVC = api-gateway

createnet:
	docker network create $(NET)

proto:
	docker run --rm --name protoc -v $(CWD)/pb:/pb -v $(CWD)/schema:/proto ezio1119/protoc \
	-I/proto \
	-I/go/src/github.com/envoyproxy/protoc-gen-validate \
	--go_out=plugins=grpc:/pb \
	--validate_out="lang=go:/pb" \
	chat.proto post.proto user.proto image.proto event.proto

gql:
	docker-compose -f $(DC_FILE) exec $(SVC) go run github.com/99designs/gqlgen generate

test:
	docker-compose -f $(DC_FILE) exec $(SVC) sh -c "go test -v -coverprofile=cover.out ./... && \
	go tool cover -html=cover.out -o ./cover.html" && \
	open ./src/cover.html

up:
	docker-compose -f $(DC_FILE) up -d $(SVC)

build:
	docker-compose -f $(DC_FILE) build

down:
	docker-compose -f $(DC_FILE) down

exec:
	docker-compose -f $(DC_FILE) exec $(SVC) sh

logs:
	docker logs -f --tail 100 $(PJT_NAME)_$(SVC)_1