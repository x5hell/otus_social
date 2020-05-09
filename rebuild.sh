#!/usr/bin/bash
source ./env.sh
export MSYS_NO_PATHCONV=1
docker exec social_web sh /run/rebuild.sh &
export MSYS_NO_PATHCONV=0