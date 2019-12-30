#!/usr/bin/bash
docker-compose rm -fsv
docker rmi $(docker images -q "social_*")