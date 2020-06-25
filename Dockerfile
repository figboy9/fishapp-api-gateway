# 開発用
FROM golang:1.13-alpine AS dev

WORKDIR /app
# tzdata: TZ環境変数からgolangがロケーションを読み込むため, git: go getするため, libc-dev & libgcc: gqlgenのgenerateに必要
RUN apk add --no-cache tzdata git libc-dev gcc && \
    go get github.com/pilu/fresh

COPY . .

CMD ["fresh"]

# コンパイラ用
FROM golang:1.13-alpine AS builder
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .

# 本番用
FROM alpine AS prod
WORKDIR /app
RUN apk add --no-cache tzdata curl

COPY healthcheck.sh /

RUN chmod +x /healthcheck.sh

COPY --from=builder /src/main .
COPY --from=builder /src/conf/conf.yml /app/conf/conf.yml

CMD ["./main"]