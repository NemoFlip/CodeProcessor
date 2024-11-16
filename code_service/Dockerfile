FROM golang:latest

# Установите рабочую директорию
RUN apt-get update && apt-get install -y docker.io
WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY code_service/ code_service/

COPY internal/entity internal/entity

COPY configs/ configs/

# Собираем бинарный файл
RUN go build -o ./bin/code_service ./code_service/cmd/code_processor/main.go

# Запускаем приложение
CMD ["./bin/code_service"]
