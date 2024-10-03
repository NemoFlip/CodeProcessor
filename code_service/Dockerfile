FROM golang:latest

# Установите рабочую директорию
RUN apt-get update && apt-get install -y docker.io
WORKDIR /app

# Копируем файлы модулей
COPY ../go.mod ../go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь проект в контейнер
COPY . .

# Собираем бинарный файл
RUN go build -o ./bin/code_service ./code_service/main.go

# Запускаем приложение
CMD ["./bin/code_service"]
