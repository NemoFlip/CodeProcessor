#syntax=docker/dockerfile:1
# Stage 1: Build
FROM golang:1.22-alpine AS builder

WORKDIR /hw1
COPY . .

RUN go mod download

RUN go build -o /homework1 cmd/http_server/main.go

# Stage 2: Final image
FROM alpine:latest

COPY --from=builder /homework1 /bin/homework1
COPY configs /configs

CMD ["/bin/homework1"]