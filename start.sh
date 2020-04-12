#!/usr/bin/bash
source ./env.sh
cp .env.example .env
source ./init_mysql_master.sh
docker-compose up --scale $SOCIAL_SLAVE_SERVICE_NAME=$SLAVE_INSTANCES