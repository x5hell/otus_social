#!/bin/bash

/helper/install_zabbix_agent.sh

docker-entrypoint.sh mysqld
