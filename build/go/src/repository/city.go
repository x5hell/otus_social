package repository

import (
	"component/database"
	"component/handler"
	"entity"
)

func GetAllCities() (cityList map[int]entity.City, err error) {
	cityList = make(map[int]entity.City)
	rows, err := database.Query("SELECT id, name FROM city")
	if err != nil {
		handler.ErrorLog(err)
		return nil, err
	}
	for rows.Next() {
		city := entity.City{}
		err = rows.Scan(&city.ID, &city.Name)
		if err != nil {
			handler.ErrorLog(err)
			return cityList, err
		}
		cityList[city.ID] = city
	}
	return cityList, err
}

func CheckCityIdExists(cityId string) (cityIdExists bool, err error) {
	rows, err := database.Query("SELECT COUNT(id) FROM city WHERE id = ?", cityId)
	if err != nil {
		handler.ErrorLog(err)
		return false, err
	}
	var countCityId int
	for rows.Next() {
		err = rows.Scan(&countCityId)
		if err != nil {
			handler.ErrorLog(err)
			return false, err
		}
	}
	return countCityId > 0, nil
}
