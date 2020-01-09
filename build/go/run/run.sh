#!/bin/sh
go get -u github.com/go-sql-driver/mysql
sh /run/rebuild.sh &
while true; do sleep 30; done