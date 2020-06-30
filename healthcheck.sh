# !/bin/sh

curl -X POST -H "Content-Type: application/json" --data '{"query": "{healthCheck}"}' localhost:8080/graphql