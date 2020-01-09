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
			return err
		}
	}
	return nil
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
