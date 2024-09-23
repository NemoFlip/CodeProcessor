#syntax=docker/dockerfile:1
FROM golang:latest

RUN apt-get update && apt-get install -y docker.io

WORKDIR hw1
COPY . .

RUN go mod download

RUN go build -o ./bin/homework1 cmd/app/main.go

CMD ["./bin/homework1"]