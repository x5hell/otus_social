#!/bin/sh
while [ ! -e /go/src/main.go ];
do
  sleep 1;
done
pkill /go/src/main
go build /go/src/main.go
/go/src/main