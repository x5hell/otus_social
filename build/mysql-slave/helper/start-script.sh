#!/bin/bash

/helper/wait-for-it.sh $SOCIAL_MASTER_CONTAINER_NAME:$MYSQL_PORT

SERVER_ID_CONFIG="/etc/mysql/conf.d/server-id.cnf"

if [ ! -f "$SERVER_ID_CONFIG" ]; then
  until mysql -h $SOCIAL_MASTER_CONTAINER_NAME -u $MYSQL_ROOT_USER -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE
  do
    echo "Try connect to master"
    sleep 5
  done

  export SERVER_ID=$(\
    echo "INSERT INTO server_id () VALUES (); SELECT LAST_INSERT_ID();" | \
    mysql -h $SOCIAL_MASTER_CONTAINER_NAME -u $MYSQL_ROOT_USER -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE |\
    awk 'NR==2')

  echo $SERVER_ID;
  echo -e "[mysqld]\nserver-id=$SERVER_ID\n" > "$SERVER_ID_CONFIG"
fi

docker-entrypoint.sh mysqld