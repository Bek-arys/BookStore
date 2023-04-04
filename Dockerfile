FROM golang:1.17-alpine AS builder

RUN mkdir /myproject

WORKDIR /myproject

COPY . .

RUN go build -o main "HP/Desktop/assignment-3/bookstore/cmd/main.go"

CMD ["./main"]