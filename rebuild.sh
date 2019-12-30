#!/usr/bin/bash
export MSYS_NO_PATHCONV=1
docker exec social_go sh /run/rebuild.sh &
export MSYS_NO_PATHCONV=0