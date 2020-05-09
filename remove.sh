#!/usr/bin/bash
source ./env.sh
docker stop $(docker ps -aq -f="name=$PROJECT_NAME*")
docker rm $(docker ps -aq -f="name=$PROJECT_NAME*")
docker rmi $(docker images -q "$PROJECT_NAME*")
docker volume rm $(docker volume ls -q)
docker volume prune -f
docker-compose rm -fsv
docker-compose -f docker-compose-test.yml rm -fsv