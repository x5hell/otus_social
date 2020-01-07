package model

import (
	"entity"
	"repository"
)

type RegistrationFormData struct {
	CityList     []entity.City
	InterestList []entity.Interest
}

func GetRegistrationFormData() (data RegistrationFormData, err error) {
	data.CityList, err = repository.GetAllCities()
	if err != nil {
		return data, err
	}
	data.InterestList, err = repository.GetAllInterests()
	return data, err
}