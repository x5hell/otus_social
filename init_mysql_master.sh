#!/usr/bin/bash
# конфигурация настроек базы мастера
cp build/mysql-master/master.cnf.example build/mysql-master/master.cnf
master_database_name=$(cat .env | grep MYSQL_DATABASE | awk -F '=' '{ print $2 }')
sed -i "s/database_name/${master_database_name}/" build/mysql-master/master.cnf
# создание учётной записи реплики
cp build/mysql-master/create_replica_account.sql build/mysql-master/init_sql/00_create_replica_account.sql
replica_user=$(cat .env | grep MYSQL_REPLICA_USER | awk -F '=' '{ print $2 }')
sed -i "s/replica_user/${replica_user}/" build/mysql-master/init_sql/00_create_replica_account.sql
replica_password=$(cat .env | grep MYSQL_REPLICA_PASSWORD | awk -F '=' '{ print $2 }')
sed -i "s/replica_password/${replica_password}/" build/mysql-master/init_sql/00_create_replica_account.sql

# конфигурация настроек базы слейва
cp build/mysql-slave/slave.cnf.example build/mysql-slave/slave.cnf
sed -i "s/database_name/${master_database_name}/" build/mysql-slave/slave.cnf

# указание слейву на мастер
cp build/mysql-slave/set_master_host_fixture.sql build/mysql-slave/init_sql/02_set_master_host_fixture.sql

# конфигурация балансировщика нагрузки
cp build/config/nginx/nginx.conf.example build/config/nginx/nginx.conf
sed -i "s/SLAVE_SERVICE_NAME/${SOCIAL_SLAVE_SERVICE_NAME}/" build/config/nginx/nginx.conf