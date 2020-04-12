#!/usr/bin/bash
export SOCIAL_MASTER_CONTAINER_NAME="social_mysql_master"
export SOCIAL_SLAVE_SERVICE_NAME="social_mysql_slave"
export SOCIAL_SITE_CONTAINER_NAME="social_go"
export SOCIAL_SITE_EXTERNAL_PORT="8000"
export SOCIAL_SITE_INTERNAL_PORT="8001"
export SLAVE_INSTANCES=2