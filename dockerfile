# FROM golang:1.20-alpine

# # Устанавливаем рабочую директорию внутри контейнера
# WORKDIR /app

# # Копируем модульные файлы Go в рабочую директорию
# COPY go.mod go.sum ./

# # Загружаем зависимости
# RUN go mod download

# # Копируем все файлы проекта в рабочую директорию
# COPY . .

# # Собираем Go приложение
# RUN go build -o blog .

# # Устанавливаем утилиту migrate с поддержкой PostgreSQL
# RUN apk add --no-cache ca-certificates && \
#     go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# # Копируем миграции в контейнер
# COPY ./migration /app/migration

# # Определяем порт, который будет использоваться
# EXPOSE 3001

# # Команда по умолчанию для выполнения миграций и запуска приложения
# CMD migrate -verbose -path /app/migration -database 'postgres://postgres:postgres@db:5432/postgres?sslmode=disable' up && ./cmd/main
# migrate to RUN 

FROM golang:1.20-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем модульные файлы Go в рабочую директорию
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все файлы проекта в рабочую директорию
COPY . .

# Устанавливаем утилиту migrate с поддержкой PostgreSQL
RUN apk add --no-cache ca-certificates && \
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Копируем миграции в контейнер
COPY ./migration /app/migration

# Определяем порт, который будет использоваться
EXPOSE 3001

# Команда по умолчанию для выполнения миграций и запуска приложения
CMD migrate -verbose -path /app/migration -database 'postgres://postgres:postgres@db:5432/postgres?sslmode=disable' up && go run cmd/main.go
