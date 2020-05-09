#!/usr/bin/bash
source ./env.sh
cp .env.test .env
source ./init_mysql_master.sh
docker-compose -p $PROJECT_NAME -f docker-compose-test.yml up -d --scale $MYSQL_SLAVE_SERVICE_NAME=$SLAVE_INSTANCES