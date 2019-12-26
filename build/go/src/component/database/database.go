package database

import (
	"database/sql"
	"fmt"
	"os"
)

const EnvSqlUser = "MYSQL_USER"
const EnvSqlPassword = "MYSQL_PASSWORD"
const EnvSqlDatabase = "MYSQL_DATABASE"

var localConnection *sql.DB = nil

func GetConnection() (connection *sql.DB, err error)  {
	if localConnection == nil {
		user, password, dbname, err := getConnectionSettings()
		if err != nil {
			return nil, err
		}
		dataSourceName := fmt.Sprintf("%s:%s@%s", user, password, dbname)
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
	res, err = connection.Exec(sqlQuery, args)
	return res, err
}

func Query(sqlQuery string, args ...interface{}) (rows *sql.Rows, err error) {
	connection, err := GetConnection()
	if err != nil {
		return nil, err
	}
	rows, err = connection.Query(sqlQuery, args)
	return rows, err
}

func Close() error {
	connection, err := GetConnection()
	if err != nil {
		return connection.Close()
	}
	return nil
}

func getConnectionSettings() (user string, password string, dbname string, err error) {
	user, envExists := os.LookupEnv(EnvSqlUser)
	if envExists == false {
		return "", "", "", fmt.Errorf("envoirment variable %s not set", EnvSqlUser)
	}
	password, envExists = os.LookupEnv(EnvSqlPassword)
	if envExists == false {
		return "", "", "", fmt.Errorf("envoirment variable %s not set", EnvSqlPassword)
	}
	dbname, envExists = os.LookupEnv(EnvSqlDatabase)
	if envExists == false {
		return "", "", "", fmt.Errorf("envoirment variable %s not set", EnvSqlDatabase)
	}
	return user, password, dbname, nil
}