#!/bin/bash
export MASTER_HOST_IP=$(getent hosts $MYSQL_MASTER_CONTAINER_NAME | awk '{print $1}')
sed -i "s/MASTER_HOST_IP/${MASTER_HOST_IP}/" /docker-entrypoint-initdb.d/02_set_master_host_fixture.sql
sed -i "s/MYSQL_REPLICA_USER/${MYSQL_REPLICA_USER}/" /docker-entrypoint-initdb.d/02_set_master_host_fixture.sql
sed -i "s/MYSQL_REPLICA_PASSWORD/${MYSQL_REPLICA_PASSWORD}/" /docker-entrypoint-initdb.d/02_set_master_host_fixture.sql