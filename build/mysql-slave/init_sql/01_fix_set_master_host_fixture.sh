#!/bin/bash
/helper/wait-for-it.sh $SOCIAL_MASTER_CONTAINER_NAME:$MYSQL_PORT
export SERVER_ID=$(echo "INSERT INTO server_id () VALUES (); SELECT LAST_INSERT_ID();" | mysql -h $SOCIAL_MASTER_CONTAINER_NAME -u $MYSQL_ROOT_USER -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE | awk 'NR==2')
echo "SET GLOBAL server_id=${SERVER_ID};" | mysql -h localhost -u $MYSQL_ROOT_USER -p$MYSQL_ROOT_PASSWORD
export MASTER_HOST_IP=$(getent hosts $SOCIAL_MASTER_CONTAINER_NAME | awk '{print $1}')
export MYSQL_BIN_FILE_NAME=$(echo $MASTER_STATUS | awk '{print $1}')
export MASTER_LOG_FILE_POSITION=$(echo $MASTER_STATUS | awk '{print $2}')
sed -i "s/MASTER_HOST_IP/${MASTER_HOST_IP}/" /docker-entrypoint-initdb.d/02_set_master_host_fixture.sql
sed -i "s/MYSQL_REPLICA_USER/${MYSQL_REPLICA_USER}/" /docker-entrypoint-initdb.d/02_set_master_host_fixture.sql
sed -i "s/MYSQL_REPLICA_PASSWORD/${MYSQL_REPLICA_PASSWORD}/" /docker-entrypoint-initdb.d/02_set_master_host_fixture.sql