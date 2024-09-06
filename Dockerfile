#syntax=docker/dockerfile:1
FROM golang:latest

WORKDIR hw1
COPY . .

RUN go mod download

RUN go build -o ./bin/homework1 cmd/main.go

CMD ["./bin/homework1"]