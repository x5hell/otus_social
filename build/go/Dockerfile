FROM golang:1.13.5-alpine3.10

WORKDIR /go/src/

COPY src /go/src/

RUN apk add git && \
    go get -u github.com/go-sql-driver/mysql && \
    go build main.go

ENTRYPOINT ["/go/src/main"]

EXPOSE 8001