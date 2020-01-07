package repository

import (
	"component/convert"
	"component/database"
	"component/handler"
	"entity"
	"fmt"
	"strings"
)

func GetLastUsers(limit int) (userList map[int]entity.User, err error) {
	rows, err := database.Query(
		"SELECT u.id, u.login, u.password, u.first_name, u.last_name, u.sex, c.name " +
			"FROM user u " +
			"LEFT JOIN city c ON (u.city_id = c.id) " +
			"ORDER BY id DESC " +
			"LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	userList = make(map[int]entity.User)
	for rows.Next() {
		user := entity.User{InterestList: []entity.Interest{}}
		err = rows.Scan(&user.ID, &user.Login, &user.Password, &user.FirstName, &user.LastName, &user.Sex, &user.City)
		if err != nil {
			return userList, err
		}
		userList[user.ID] = user
	}
	userList, err = fillUserInterests(userList)
	_ = database.Close()
	return userList, err
}

func LoginExists(login string) (result bool, err error) {
	rows, err := database.Query("SELECT COUNT(id) FROM user WHERE login = ?", login)
	if err != nil {
		return false, err
	}
	rows.Next()
	countUsers := 0
	err = rows.Scan(&countUsers)
	return countUsers > 0, err
}

func fillUserInterests(userList map[int]entity.User) (userListWithInterest map[int]entity.User, err error) {
	if len(userList) > 0 {
		var userIdList []int
		var userIdPlaceList []string
		for userId, _ := range userList {
			userIdList = append(userIdList, userId)
			userIdPlaceList = append(userIdPlaceList, "?")
		}
		userIdListQuery := strings.Join(userIdPlaceList, ",")
		sqlQuery := fmt.Sprintf(
			"SELECT ui.user_id, i.name, i.id " +
				"FROM user_interest ui " +
				"INNER JOIN interestName i ON (ui.interest_id = i.id) " +
				"WHERE ui.user_id IN (%s)", userIdListQuery)
		rows, err := database.Query(sqlQuery, userIdList)
		if err != nil {
			return userList, err
		}
		var userId, interestId int
		var interestName string
		for rows.Next() {
			err = rows.Scan(&userId, &interestName, &interestId)
			if err != nil {
				return userList, err
			}
			user := userList[userId]
			user.InterestList = append(user.InterestList, entity.Interest{ID: interestId, Name: interestName})
		}
	}
	return userList, nil
}

func InsertUser(user *entity.User) error {
	connect, err :=  database.GetConnection()
	if err != nil {
		return err
	}
	transaction, err := connect.Begin()
	if err != nil {
		return err
	}
	password := convert.StringToMd5(user.Password)
	var sex, cityId interface{}
	if len(user.Sex) > 0 {
		sex = user.Sex
	} else {
		sex = nil
	}
	if user.City.ID > 0 {
		cityId = user.City.ID
	} else {
		cityId = nil
	}
	sqlQuery := "INSERT INTO user (login, password, first_name, last_name, age, sex, city_id) " +
		"VALUES  (?, ?, ?, ?, ?, ?, ?)"
	sqlResult, err := transaction.Exec(sqlQuery,
		user.Login, password, user.FirstName, user.LastName, user.Age, sex, cityId)
	if err != nil {
		handler.ErrorLog(transaction.Rollback())
		return err
	}

	userId, err := sqlResult.LastInsertId()
	if err != nil {
		return err
	}

	if len(user.InterestList) > 0 {
		var queryParameters []int64
		var queryParametersPlaceholder []string
		for _, interest := range user.InterestList {
			queryParametersPlaceholder = append(queryParametersPlaceholder, "(?, ?)")
			queryParameters = append(queryParameters, userId, int64(interest.ID))
		}
		sqlQuery = fmt.Sprintf("INSERT INTO user_interest (user_id, interest_id) VALUES %s;",
			strings.Join(queryParametersPlaceholder, ", "))
		_, err = transaction.Exec(sqlQuery, convert.Int64ListToInterfaceList(queryParameters)...)
		if err != nil {
			handler.ErrorLog(transaction.Rollback())
			return err
		}
	}
	handler.ErrorLog(transaction.Commit())
	user.ID = int(userId)
	return nil
}
