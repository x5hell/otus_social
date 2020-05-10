#!/bin/sh
go get github.com/golang/freetype

until mysql -h $MYSQL_MASTER_HOSTNAME -u $MYSQL_ROOT_USER -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE
do
  echo "Try connect to $MYSQL_MASTER_HOSTNAME"
  sleep 5
done

sleep 10

sh /run/rebuild.sh
/go/src/main --action=applyFixture
/go/src/main --action=addIndex
/go/src/main --action=setAppWorkModeMasterOnly
/go/src/main --action=testWithIndex
/go/src/main --action=setAppWorkModeUseReplica
/go/src/main --action=testWithIndex
while true; do sleep 30; done