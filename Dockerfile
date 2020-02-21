FROM golang:1.13-alpine AS builder
WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .
RUN go mod download

COPY src .
RUN go build -o main .

FROM alpine
CMD ["./main"]
COPY --from=builder /app .