# Stage 1 — build Go binary
FROM golang:1.23 AS build

WORKDIR /app

# Копируем go.mod и go.sum отдельно — ускоряет сборку
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Сборка бинаря
RUN CGO_ENABLED=0 GOOS=linux go build -o calendar_service main.go

# Stage 2 — минимальный образ
FROM alpine:3.18

WORKDIR /app

COPY --from=build /app/calendar_service .
COPY docs ./docs

EXPOSE 8080

CMD ["./calendar_service"]
