#!/usr/bin/bash
source ./env.sh
cp .env.example .env
source ./init_mysql_master.sh
echo "docker-compose -p $PROJECT_NAME up --scale $SOCIAL_SLAVE_SERVICE_NAME=$SLAVE_INSTANCES"
docker-compose -p $PROJECT_NAME up --scale $SOCIAL_SLAVE_SERVICE_NAME=$SLAVE_INSTANCES