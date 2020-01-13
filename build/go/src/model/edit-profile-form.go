package model

import (
	"entity"
	"repository"
)

type EditProfileFormData struct {
	CityList       map[int]entity.City
	InterestList   []entity.Interest
	User           entity.User
	UserCityId     int
	InterestToUser map[int]int
}

func GetEditProfileFormData(userId int) (data EditProfileFormData, err error) {
	data.CityList, err = repository.GetAllCities()
	if err != nil {
		return data, err
	}
	data.InterestList, err = repository.GetAllInterests()
	if err != nil {
		return data, err
	}
	data.User, err = repository.GetUserById(userId)
	if err != nil {
		return data, err
	}
	data.UserCityId = int(data.User.CityId.Int64)
	data.InterestToUser, err = repository.GetInterestToUser(userId)
	return data, err
}