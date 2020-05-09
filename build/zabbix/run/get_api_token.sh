#!/bin/bash

function get_api_token {
  curl \
    -s \
    -d '{"jsonrpc":"2.0","method":"user.login","id":1,"auth":null,"params":{"user": "Admin", "password": "zabbix"}}' \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/api_jsonrpc.php \
  | \
  awk -F '"' '{ print $8 }'
}

while [ $(expr length $(get_api_token)) != 32 ]; do
  sleep 5
done

echo $(get_api_token)