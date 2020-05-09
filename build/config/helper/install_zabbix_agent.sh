#!/bin/bash

zabbix_not_intalled=`whereis -b zabbix-agent | wc -w`

if [ $zabbix_not_intalled -eq 1 ]; then
  apt-get update
  apt-get install -y zabbix-agent sysstat
  zabbix_server_ip=$(getent hosts $ZABBIX_SERVER_HOST | awk '{ print $1 }')
  sed -i "s/Server\=.*/Server\=${zabbix_server_ip}/" /etc/zabbix/zabbix_agentd.conf
  sed -i "s/ServerActive\=.*/ServerActive\=${zabbix_server_ip}/" /etc/zabbix/zabbix_agentd.conf
  service zabbix-agent start

  disks=$(iostat 1 1 | grep -e ^s | wc -l)
  while(true)
  do
    discs_read_kbps=$(iostat -k 1 2 | grep -e ^s | tail -$disks | awk ' { sum += $2 } END { print sum }');
    echo $discs_read_kbps > /var/log/discs_read_kbps.log;
    sleep 26
  done &

fi