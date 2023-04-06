FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main /app/cmd/main/main.go

FROM alpine:latest

COPY --from=builder /app/main /app/main

WORKDIR /app

EXPOSE 8080

CMD ["/app/main"]