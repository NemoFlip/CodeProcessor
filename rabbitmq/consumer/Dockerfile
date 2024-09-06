FROM golang:alpine AS builder

WORKDIR code

RUN apk add --no-cache python3 clang g++ musl-dev

COPY ../../pycode.py .
COPY ../../cCode.c .
COPY ../../gcc.cpp .
COPY rabbitmq/consumer/run.sh .

RUN chmod +x run.sh

CMD ["/bin/sh", "run.sh"]