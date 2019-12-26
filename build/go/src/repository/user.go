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
			"ORDER BY id DESC" +
			"LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	userList = make(map[int]model.User)
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.ID, &user.Login, &user.Password, &user.FirstName, &user.LastName, &user.Sex, &user.City)
		if err != nil {
			return userList, err
		}
		userList[user.ID] = user
	}
	return fillUserInterests(userList)
}

func fillUserInterests(userList map[int]model.User) (userListWithInterest map[int]model.User, err error) {
	var userIdList []int
	var userIdPlaceList []string
	for userId, user := range userList {
		userIdList = append(userIdList, user.ID)
		userIdPlaceList = append(userIdPlaceList, "?")
	}
	userIdListQuery := strings.Join(userIdPlaceList, ",")
	sqlQuery := fmt.Sprintf(
		"SELECT ui.user_id, i.name " +
		"FROM user_interest ui " +
		"INNER JOIN interest i ON (ui.interest_id = i.id) " +
		"WHERE ui.user_id IN"
}
