FROM golang:1.25.4

USER root

# Устанавливаем postgres-client для pg_isready
RUN apt-get update && \
    apt-get install -y postgresql-client && \
    go install github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /migrations

CMD ["sh", "-c", "until pg_isready -h db -p 5432; do sleep 1; done; goose postgres \"postgres://postgres:postgres@db:5432/qa?sslmode=disable\" up"]
