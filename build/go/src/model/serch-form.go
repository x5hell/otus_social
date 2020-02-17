package model

import (
	"entity"
	"repository"
)

type SearchFormData struct {
	CityList     map[int]entity.City
}

func GetSearchFormData() (data SearchFormData, err error) {
	data.CityList, err = repository.GetAllCities()
	if err != nil {
		return data, err
	}
	return data, err
}
