package model

import (
	"entity"
	"repository"
)

type EditProfileFormData struct {
	CityList     []entity.City
	InterestList []entity.Interest
	User entity.User
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
	return data, err
}