package model

import (
	"entity"
	"repository"
)

type UserProfileList struct {
	Users 					[]entity.User
	CityList    			map[int]entity.City
	UserIdToInterestList 	map[int][]entity.Interest
}

func GetUserProfileListData(limit int) (data UserProfileList, err error) {
	lastUsers, err := repository.GetLastUsers(limit)
	if err != nil {
		return data, err
	}
	data.Users = lastUsers
	data.CityList, err = repository.GetUserCityList(data.Users)
	if err != nil {
		return data, err
	}
	data.UserIdToInterestList, err = repository.GetUserIdToInterestList(data.Users)
	if err != nil {
		return data, err
	}
	return data, err
}

