package repository

import (
	"component/database"
	"fmt"
	"model"
	"strings"
)

func GetLastUsers(limit int) (userList map[int]model.User, err error) {
	rows, err := database.Query(
		"SELECT u.id, u.login, u.password, u.first_name, u.last_name, u.sex, c.name " +
			"FROM user u " +
			"LEFT JOIN city c ON (u.city_id = c.id) " +
			"ORDER BY id DESC " +
			"LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	userList = make(map[int]model.User)
	for rows.Next() {
		user := model.User{InterestList: []model.Interest{}}
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

func fillUserInterests(userList map[int]model.User) (userListWithInterest map[int]model.User, err error) {
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
			user.InterestList = append(user.InterestList, model.Interest{ID:interestId, Name: interestName})
		}
	}
	return userList, nil
}
