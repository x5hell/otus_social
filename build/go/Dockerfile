FROM golang:1.13.5-alpine3.10

WORKDIR /go/src/

COPY src /go/src/

RUN go build main.go

ENTRYPOINT ["/go/src/main"]

EXPOSE 8001