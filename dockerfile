FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache bash

COPY wait-for-it.sh ./wait-for-it.sh
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main ./main.go

CMD ["./main"]
