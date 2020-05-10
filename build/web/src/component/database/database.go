package database

import (
	"component/environment"
	"component/handler"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const WorkModeEnv = "DB_WORK_MODE"

type WorkMode string

const (
	WorkModeUseReplica = "useReplica"
	WorkModeMasterOnly = "masterOnly"
)

type ConnectionSettingsEnvironment struct {
	Host string
	Port string
	Database string
	User string
	Password string
}

type ConnectionSettings struct {
	Host string
	Port string
	Database string
	User string
	Password string
}

type Connection struct {
	Connect *sql.DB
	Environment ConnectionSettingsEnvironment
	Settings ConnectionSettings
}

type ConnectionRegistry struct {
	Master Connection
	Slave Connection
}

var connectionRegistry ConnectionRegistry

func GetWorkMode() (result string, err error) {
	workMode, err := environment.Get(WorkModeEnv)
	if err != nil {
		return result, err
	}
	err = ValidateWorkMode(workMode)
	if err != nil {
		return result, err
	}
	return workMode, nil
}

func ValidateWorkMode(workMode string) (err error) {
	switch workMode {
		case WorkModeUseReplica:
		case WorkModeMasterOnly:
			return nil
		default:
			return fmt.Errorf(
				"ivalid DbWorkMode %s expected: %s, %s ",
				workMode,
				WorkModeUseReplica,
				WorkModeMasterOnly,
			)
	}
	return nil
}

func Master() (master Connection) {
	return connectionRegistry.Master
}

func Slave() (slave Connection) {
	workMode, _ := environment.Get(WorkModeEnv)
	switch workMode {
		case WorkModeMasterOnly:
			return connectionRegistry.Master
		case WorkModeUseReplica:
		default:
			return connectionRegistry.Slave
	}
	return connectionRegistry.Slave
}

func InitConnectionRegistry(master Connection, slave Connection) {
	connectionRegistry.Master = master
	connectionRegistry.Slave = slave
}

func CloseRegistryConnections()  {
	_ = connectionRegistry.Master.Close()
	_ = connectionRegistry.Slave.Close()
}

func (env ConnectionSettingsEnvironment) getConnectionSettings() (connectionSettings ConnectionSettings, err error) {
	connectionSettings.User, err = environment.Get(env.User)
	if err != nil {
		return connectionSettings, err
	}
	connectionSettings.Password, err = environment.Get(env.Password)
	if err != nil {
		return connectionSettings, err
	}
	connectionSettings.Host, err = environment.Get(env.Host)
	if err != nil {
		return connectionSettings, err
	}
	connectionSettings.Port, err = environment.Get(env.Port)
	if err != nil {
		return connectionSettings, err
	}
	connectionSettings.Database, err = environment.Get(env.Database)
	if err != nil {
		return connectionSettings, err
	}
	return connectionSettings, nil
}

func (settings ConnectionSettings) getConnection() (connection *sql.DB, err error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		settings.User,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.Database,
	)
	return sql.Open("mysql", dataSourceName)
}

func (connection Connection) getConnection() (dbConnection *sql.DB, err error) {
	if connection.Connect == nil {
		if connection.Settings.Host == "" {
			connection.Settings, err = connection.Environment.getConnectionSettings()
			if err != nil {
				return nil, err
			}
		}
		connection.Connect, err = connection.Settings.getConnection()
		if err != nil {
			handler.ErrorLog(err)
			return nil, err
		}
	}
	return connection.Connect, nil
}

func (connection Connection) Close() error {
	dbConnect, err := connection.getConnection()
	if err != nil {
		return err
	}
	err = dbConnect.Close()
	handler.ErrorLog(err)
	return err
}

func (connection Connection) Exec(sqlQuery string, args ...interface{}) (res sql.Result, err error) {
	dbConnect, err := connection.getConnection()
	if err != nil {
		return nil, err
	}
	res, err = dbConnect.Exec(sqlQuery, args...)
	handler.ErrorLog(err)
	return res, err
}

func (connection Connection) Query(sqlQuery string, args ...interface{}) (rows *sql.Rows, err error) {
	dbConnect, err := connection.getConnection()
	if err != nil {
		return nil, err
	}
	stmt, err := dbConnect.Prepare(sqlQuery)
	if err != nil {
		handler.ErrorLog(err)
		return nil, err
	}
	rows, err = stmt.Query(args...)
	handler.ErrorLog(err)
	return rows, err
}

func (connection Connection) GetTransaction() (transaction *sql.Tx, err error)  {
	dbConnect, err := connection.getConnection()
	if err != nil {
		return nil, err
	}
	transaction, err = dbConnect.Begin()
	handler.ErrorLog(err)
	return transaction, err
}