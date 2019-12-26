package database

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

const EnvSqlUser = "MYSQL_USER"
const EnvSqlPassword = "MYSQL_PASSWORD"
const EnvSqlDatabase = "MYSQL_DATABASE"
const EnvSqlPort = "MYSQL_PORT"
const EnvSqlHost = "MYSQL_HOST"

var localConnection *sql.DB = nil

func GetConnection() (connection *sql.DB, err error)  {
	if localConnection == nil {
		user, password, host, port, dbname, err := getConnectionSettings()
		if err != nil {
			return nil, err
		}
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
		connection, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			return nil, err
		}
		localConnection = connection
	}
	return localConnection, nil
}

func Exec(sqlQuery string, args ...interface{}) (res sql.Result, err error) {
	connection, err := GetConnection()
	if err != nil {
		return nil, err
	}
	res, err = connection.Exec(sqlQuery, args...)
	return res, err
}

func Query(sqlQuery string, args ...interface{}) (rows *sql.Rows, err error) {
	connection, err := GetConnection()
	if err != nil {
		return nil, err
	}
	rows, err = connection.Query(sqlQuery, args...)
	return rows, fmt.Errorf(sqlQuery)
}

func Close() error {
	connection, err := GetConnection()
	if err != nil {
		return connection.Close()
	}
	return nil
}

func getConnectionSettings() (user, password, host, port, dbname string, err error) {
	user, envExists := os.LookupEnv(EnvSqlUser)
	if envExists == false {
		return "", "", "", "", "", fmt.Errorf("envoirment variable %s not set", EnvSqlUser)
	}
	password, envExists = os.LookupEnv(EnvSqlPassword)
	if envExists == false {
		return "", "", "", "", "", fmt.Errorf("envoirment variable %s not set", EnvSqlPassword)
	}
	host, envExists = os.LookupEnv(EnvSqlHost)
	if envExists == false {
		return "", "", "", "", "", fmt.Errorf("envoirment variable %s not set", EnvSqlHost)
	}
	port, envExists = os.LookupEnv(EnvSqlPort)
	if envExists == false {
		return "", "", "", "", "", fmt.Errorf("envoirment variable %s not set", EnvSqlPort)
	}
	dbname, envExists = os.LookupEnv(EnvSqlDatabase)
	if envExists == false {
		return "", "", "", "", "", fmt.Errorf("envoirment variable %s not set", EnvSqlDatabase)
	}
	return user, password, host, port, dbname, nil
}