package repository

import (
	"component/convert"
	"component/database"
	"component/handler"
	"database/sql"
	"entity"
	"fmt"
)

func GetLastUsers(limit int) (userList map[int]entity.User, err error) {
	rows, err := database.Query(
		"SELECT id, login, password, first_name, last_name, sex, city_id " +
			"FROM user " +
			"ORDER BY id DESC " +
			"LIMIT ?", limit)
	if err != nil {
		handler.ErrorLog(err)
		return nil, err
	}
	userList = make(map[int]entity.User)
	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(
			&user.ID, &user.Login, &user.Password, &user.FirstName, &user.LastName, &user.Sex, &user.CityId)
		if err != nil {
			handler.ErrorLog(err)
			return userList, err
		}
		userList[user.ID] = user
	}
	return userList, err
}

func LoginExists(login string) (result bool, err error) {
	rows, err := database.Query("SELECT COUNT(id) FROM user WHERE login = ?", login)
	if err != nil {
		handler.ErrorLog(err)
		return false, err
	}
	rows.Next()
	countUsers := 0
	err = rows.Scan(&countUsers)
	handler.ErrorLog(err)
	return countUsers > 0, err
}

func InsertUser(user *entity.User, transaction *sql.Tx) error {
	password := convert.StringToMd5(user.Password)
	sqlQuery := "INSERT INTO user (login, password, first_name, last_name, age, sex, city_id) " +
		"VALUES  (?, ?, ?, ?, ?, ?, ?)"
	sqlResult, err := transaction.Exec(sqlQuery,
		user.Login, password, user.FirstName, user.LastName, user.Age, user.Sex, user.CityId)
	if err != nil {
		return err
	}

	userId, err := sqlResult.LastInsertId()
	if err != nil {
		handler.ErrorLog(err)
		return err
	}
	user.ID = int(userId)
	return nil
}

func UpdateUser(user *entity.User) error {
	connect, err :=  database.GetConnection()
	if err != nil {
		handler.ErrorLog(err)
		return err
	}
	transaction, err := connect.Begin()
	if err != nil {
		handler.ErrorLog(err)
		return err
	}
	password := convert.StringToMd5(user.Password)
	sqlQuery :=
		"UPDATE user SET login = ?, password = ?, first_name = ?, last_name = ?, age = ?, sex = ?, city_id = ?) " +
		"WHERE id = ?"
	_, err = transaction.Exec(sqlQuery,
		user.Login, password, user.FirstName, user.LastName, user.Age, user.Sex, user.CityId, user.ID)
	if err != nil {
		handler.ErrorLog(err)
		handler.ErrorLog(transaction.Rollback())
		return err
	}

	_, err = transaction.Exec("DELETE FROM user_interest WHERE user_id = ?", user.ID)
	if err != nil {
		handler.ErrorLog(err)
		handler.ErrorLog(transaction.Rollback())
		return err
	}

	handler.ErrorLog(transaction.Commit())
	return nil
}

func GetUserByAuth(login, password string) (user *entity.User, err error) {
	passwordMd5 := convert.StringToMd5(password)
	rows, err := database.Query(
		"SELECT id FROM user WHERE login = ? AND password = ? ", login, passwordMd5)
	if err != nil {
		handler.ErrorLog(err)
		return nil, err
	}
	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.ID)
		handler.ErrorLog(err)
		return &user, err
	}
	return nil, fmt.Errorf("пользователь с таким логином и паролем не найден")
}

func GetUserById(userId int) (user entity.User, err error) {
	rows, err := database.Query(
		"SELECT u.id, u.login, u.password, u.first_name, u.last_name, u.age, u.sex, u.city_id " +
			"FROM user u WHERE u.id = ?", userId)
	if err != nil {
		handler.ErrorLog(err)
		return user, err
	}
	for rows.Next() {
		err = rows.Scan(
			&user.ID, &user.Login, &user.Password, &user.FirstName, &user.LastName,
			&user.Age, &user.Sex, &user.CityId)
		if err != nil {
			handler.ErrorLog(err)
			return user, err
		}
	}
	return user, err
}