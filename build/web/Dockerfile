FROM golang:1.13.5-alpine3.10

WORKDIR /go/src/

COPY run /run

RUN apk add git && \
    chmod +x /run/run.sh /run/rebuild.sh

ENTRYPOINT ["sh", "/run/run.sh"]

EXPOSE 8001