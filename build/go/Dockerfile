FROM golang:1.13.5-alpine3.10

WORKDIR /go/src/

COPY run.sh /run.sh

RUN apk add git && \
    go get -u github.com/go-sql-driver/mysql && \
    chmod +x /run.sh

ENTRYPOINT ["sh", "/run.sh"]

EXPOSE 8001