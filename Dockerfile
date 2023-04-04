FROM golang:1.17-alpine AS builder

RUN mkdir /myproject

WORKDIR /myproject

RUN go get github.com/Bek-arys/BookStore/cmd/main
RUN cd /myproject && git clone https://github.com/Bek-arys/BookStore.git

RUN cd /myproject/cmd/main && go build

EXPOSE 8080

ENTRYPOINT ["/myproject/cmd/main/main"]