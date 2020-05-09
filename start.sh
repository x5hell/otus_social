#!/usr/bin/bash
source ./env.sh
cp .env.example .env
source ./init_mysql_master.sh
docker-compose -p $PROJECT_NAME up -d --scale $MYSQL_SLAVE_SERVICE_NAME=$SLAVE_INSTANCES