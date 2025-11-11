# Stage 1: build
FROM golang:1.25.4-alpine AS builder

WORKDIR /app

# Устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем бинарник
RUN go build -o /qa-api ./cmd/server

# Stage 2: minimal container
FROM alpine:latest

WORKDIR /app

# Копируем бинарник
COPY --from=builder /qa-api .

# Запуск
CMD ["./qa-api"]
