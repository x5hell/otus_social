package repository

import (
	"component/database"
	"fmt"
	"model"
)

func GetAllCities() (cityList []model.City, err error) {
	rows, err := database.Query("SELECT id, name FROM city")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		city := model.City{}
		err = rows.Scan(&city.ID, &city.Name)
		if err != nil {
			return cityList, err
		}
		cityList = append(cityList, city)
	}
	return cityList, err
}

func CheckCityIdExists(cityId string) (cityIdExists bool, err error) {
	rows, err := database.Query("SELECT COUNT(id) FROM city WHERE id = ?", cityId)
	if err != nil {
		return false, err
	}
	var countCityId int
	for rows.Next() {
		err = rows.Scan(&countCityId)
		if err != nil {
			return false, err
		}
	}
	return countCityId > 0, nil
}
