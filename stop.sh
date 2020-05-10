#!/usr/bin/bash
source ./env.sh
docker-compose stop
docker-compose -p $PROJECT_NAME -f docker-compose-test.yml stop