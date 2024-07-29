
FROM golang:1.20 AS build

# Установим рабочий каталог внутри контейнера
WORKDIR /app

# Скопируем файлы go.mod и go.sum и установим зависимости
COPY go.mod go.sum ./
RUN go mod download

# Скопируем остальные файлы исходного кода
COPY . .

# Соберем бинарный файл приложения
RUN go build -o blog .

# Используем минимальный образ для запуска собранного приложения
FROM alpine:latest

# Установим сертификаты
RUN apk --no-cache add ca-certificates

# Скопируем собранный бинарный файл из предыдущего этапа
COPY --from=build /app/blog /blog

# Определим команду запуска контейнера
CMD ["/blog"]
