FROM golang:1.13.5-alpine3.10

WORKDIR /go/src/

COPY run /run

RUN apk add mysql-client mariadb-connector-c gcc g++ && \
    chmod +x /run/*.sh

ENTRYPOINT ["sh", "/run/run.sh"]