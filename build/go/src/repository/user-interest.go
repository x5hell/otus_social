package repository

import (
	"component/convert"
	"component/database"
	"component/handler"
	"database/sql"
	"entity"
	"fmt"
	"strings"
)

func InsertUserInterestEntityList(userInterestsList []entity.UserInterest, transaction *sql.Tx) (err error)  {
	if len(userInterestsList) > 0 {
		var queryParameters []int
		var queryParametersPlaceholder []string
		for _, userInterest := range userInterestsList {
			queryParametersPlaceholder = append(queryParametersPlaceholder, "(?, ?)")
			queryParameters = append(queryParameters, userInterest.UserId, userInterest.InterestId)
		}
		sqlQuery := fmt.Sprintf("INSERT INTO user_interest (user_id, interest_id) VALUES %s;",
			strings.Join(queryParametersPlaceholder, ", "))
		_, err = transaction.Exec(sqlQuery, convert.IntListToInterfaceList(queryParameters)...)
		if err != nil {
			handler.ErrorLog(err)
			return err
		}
	}
	return nil
}

func UpdateUserInterestEntityList(userId int, userInterestsList []entity.UserInterest, transaction *sql.Tx) (err error) {
	_, err = transaction.Exec("DELETE FROM user_interest WHERE user_id = ?", userId)
	if err != nil {
		handler.ErrorLog(err)
		return err
	}
	return InsertUserInterestEntityList(userInterestsList, transaction)
}

func GetInterestToUser(userId int) (interestToUser map[int]int, err error) {
	rows, err := database.Query(
		"SELECT user_id, interest_id FROM user_interest WHERE user_id = ? ", userId)
	if err != nil {
		handler.ErrorLog(err)
		return nil, err
	}
	interestToUser = make(map[int]int)
	for rows.Next() {
		userInterest := entity.UserInterest{}
		err = rows.Scan(&userInterest.UserId, &userInterest.InterestId)
		if err != nil {
			handler.ErrorLog(err)
			return interestToUser, err
		}
		interestToUser[userInterest.InterestId] = userInterest.UserId
	}
	return interestToUser, err
}

func GetUserIdToInterestList(userList []entity.User) (userIdToInterestList map[int][]entity.Interest, err error) {
	userIdToInterestList = make(map[int][]entity.Interest)
	if len(userList) > 0 {
		userIdList := GetUserIdsFromUserList(userList)
		userIdPlaceList := convert.IntListToQueryParameterPlaceList(userIdList)
		userIdListQuery := strings.Join(userIdPlaceList, ",")
		sqlQuery := fmt.Sprintf(
			"SELECT ui.user_id, ui.interest_id, i.name FROM user_interest ui " +
				"INNER JOIN interest i ON (ui.interest_id = i.id) " +
				"WHERE ui.user_id IN (%s)", userIdListQuery)
		rows, err := database.Query(sqlQuery, convert.IntListToInterfaceList(userIdList)...)
		if err != nil {
			handler.ErrorLog(err)
			return nil, err
		}
		for rows.Next() {
			var userId int
			if _, userIdInterestListExists := userIdToInterestList[userId]; userIdInterestListExists == false {
				userIdToInterestList[userId] = []entity.Interest{}
			}
			interest := entity.Interest{}
			err = rows.Scan(&userId, &interest.ID, &interest.Name)
			if err != nil {
				handler.ErrorLog(err)
				return userIdToInterestList, err
			}
			userIdToInterestList[userId] = append(userIdToInterestList[userId], interest)
		}
	}
	return userIdToInterestList, nil
}

func GetUserInterestList(user entity.User) (interestList []entity.Interest, err error) {
	userIdToInterestList, err := GetUserIdToInterestList([]entity.User{user})
	if err != nil {
		return interestList, nil
	}
	interestList = userIdToInterestList[user.ID]
	return interestList, nil
}