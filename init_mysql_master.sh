#!/usr/bin/bash
# конфигурация настроек базы мастера
cp build/mysql-master/master.cnf.example build/mysql-master/master.cnf
sed -i "s/database_name/${MYSQL_DATABASE}/" build/mysql-master/master.cnf
# создание учётной записи реплики
cp build/mysql-master/create_replica_account.sql build/mysql-master/init_sql/00_create_replica_account.sql
sed -i "s/replica_user/${MYSQL_REPLICA_USER}/" build/mysql-master/init_sql/00_create_replica_account.sql
sed -i "s/replica_password/${MYSQL_REPLICA_PASSWORD}/" build/mysql-master/init_sql/00_create_replica_account.sql

# конфигурация настроек базы слейва
cp build/mysql-slave/slave.cnf.example build/mysql-slave/slave.cnf
sed -i "s/database_name/${MYSQL_DATABASE}/" build/mysql-slave/slave.cnf

# указание слейву на мастер
cp build/mysql-slave/set_master_host_fixture.sql build/mysql-slave/init_sql/02_set_master_host_fixture.sql

# конфигурация балансировщика нагрузки
cp build/config/nginx/nginx.conf.example build/config/nginx/nginx.conf
sed -i "s/SLAVE_SERVICE_NAME/${MYSQL_SLAVE_SERVICE_NAME}/" build/config/nginx/nginx.conf

# конфигурирование базы zabbix
cp build/mysql-zabbix/alter_zabbix_account.sql build/mysql-zabbix/init_sql/01_alter_zabbix_account.sql
sed -i "s/zabbix_user/${ZABBIX_MYSQL_USER}/" build/mysql-zabbix/init_sql/01_alter_zabbix_account.sql
sed -i "s/zabbix_user/${ZABBIX_MYSQL_PASSWORD}/" build/mysql-zabbix/init_sql/01_alter_zabbix_account.sql
