#!/usr/bin/bash
source ./env.sh
docker-compose stop
docker-compose -f docker-compose-test.yml stop