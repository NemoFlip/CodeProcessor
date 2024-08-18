FROM golang:latest

WORKDIR hw1
COPY . .

RUN go mod download

RUN go build -o ./bin/homework1 main.go

CMD ["./bin/homework1"]