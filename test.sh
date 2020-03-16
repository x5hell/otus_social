#!/usr/bin/bash
source ./env.sh
cp .env.test .env
docker-compose -f docker-compose-test.yml up -d