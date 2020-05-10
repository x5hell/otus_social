#!/usr/bin/bash
source ./env.sh
./stop.sh
docker-compose -p $PROJECT_NAME rm -fsv
docker-compose -p $PROJECT_NAME -f docker-compose-test.yml rm -fsv
docker stop $(docker ps -aq -f="name=$PROJECT_NAME*")
docker rm $(docker ps -aq -f="name=$PROJECT_NAME*")
docker rmi $(docker images -q "$PROJECT_NAME*")
docker volume rm $(docker volume ls -q)
docker volume prune -f