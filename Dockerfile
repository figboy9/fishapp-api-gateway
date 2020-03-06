FROM golang:1.13-alpine AS builder
WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine
WORKDIR /app
CMD ["./main"]
COPY --from=builder /src/main .
COPY --from=builder /src/conf/conf.yml /app/conf/conf.yml