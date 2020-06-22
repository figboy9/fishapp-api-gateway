CWD = $(shell pwd)
PJT_NAME = $(notdir $(PWD))
NET = fishapp-net
DC_FILE = docker-compose.yml

SVC = api-gateway

IMAGE_URL = image:50051
POST_URL = post:50051
CHAT_URL = chat:50051
USER_URL = image:50051
NATS_URL = nats-streaming:50051

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
	docker-compose -f $(DC_FILE) exec api-gateway go run github.com/99designs/gqlgen generate

waitnats:
	docker run --rm --name dockerize --net $(NET) jwilder/dockerize \
	-wait tcp://$(NATS_URL)

waitimage:
	docker run --rm --name grpc_health_probe --net $(NET) stefanprodan/grpc_health_probe:v0.3.0 \
	grpc_health_probe -addr=$(IMAGE_URL)

waitpost:
	docker run --rm --name grpc_health_probe --net $(NET) stefanprodan/grpc_health_probe:v0.3.0 \
	grpc_health_probe -addr=$(POST_URL)

waitchat:
	docker run --rm --name grpc_health_probe --net $(NET) stefanprodan/grpc_health_probe:v0.3.0 \
	grpc_health_probe -addr=$(CHAT_URL)

waituser:
	docker run --rm --name grpc_health_probe --net $(NET) stefanprodan/grpc_health_probe:v0.3.0 \
	grpc_health_probe -addr=$(USER_URL)

test:
	docker-compose -f $(DC_FILE) exec $(SVC) sh -c "go test -v -coverprofile=cover.out ./... && \
	go tool cover -html=cover.out -o ./cover.html" && \
	open ./src/cover.html

up: waitimage waitchat waituser waitpost waitnats
	docker-compose -f $(DC_FILE) up -d $(SVC)

build:
	docker-compose -f $(DC_FILE) build

down:
	docker-compose -f $(DC_FILE) down

exec:
	docker-compose -f $(DC_FILE) exec $(SVC) sh

logs:
	docker logs -f --tail 100 $(PJT_NAME)_$(SVC)_1

rmvol:
	docker-compose -f $(DC_FILE) down -v
