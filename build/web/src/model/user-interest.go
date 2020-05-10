package model

import (
	"component/handler"
	"entity"
	"strconv"
)

func BuildUserInterestEntityList(requestInterestsList []string, userId int) (userInterestsList []entity.UserInterest) {
	for _, interestId := range requestInterestsList {
		interestIdInt, err := strconv.Atoi(interestId)
		handler.ErrorLog(err)
		userInterestsList = append(userInterestsList, entity.UserInterest{
			UserId:     userId,
			InterestId: interestIdInt,
		})
	}
	return userInterestsList
}