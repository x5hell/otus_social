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

func GetLastUsersProfileListData(limit int) (data UserProfileList, err error) {
	lastUsers, err := repository.GetLastUsers(limit)
	if err != nil {
		return data, err
	}
	return GetUserProfileListData(lastUsers)
}

func GetUserProfileListData(userList []entity.User) (data UserProfileList, err error) {
	data.Users = userList
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

