#!/bin/bash
/helper/wait-for-it.sh localhost:8080

api_token=$(/run/get_api_token.sh)

function import_config {
  # перед вызовом функции должны быть определены переменные
  # $api_token - с токеном доступа к api zabbix
  # $import_xml - импортируемые параметры в формате xml
  # $import_json - json шаблон запроса импорта

  import_xml=$( echo $import_xml | python3 -c 'import json,sys; print(json.dumps(sys.stdin.read()).strip("\""))')
  import_json="${import_json/\$import_xml/$import_xml}"
  import_json="${import_json/\$api_token/$api_token}"

  curl \
    -s \
    -d "$import_json" \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/api_jsonrpc.php
}

# Импорт хостов
echo -e "\nimport hosts"
import_xml=$( cat /config/zbx_export_hosts.xml )
import_json=$( cat /config/import_hosts.json )
import_config

# Импорт экранов
echo -e "\nimport screens"
import_xml=$( cat /config/zbx_export_screens.xml )
import_json=$( cat /config/import_screens.json )
import_config
