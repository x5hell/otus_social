#!/usr/bin/bash
docker-compose rm -fsv
docker-compose -f docker-compose-test.yml rm -fsv
docker rmi $(docker images -q "social_*")