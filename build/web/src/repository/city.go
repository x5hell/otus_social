package repository

import (
	"component/convert"
	"component/database"
	"component/handler"
	"entity"
	"fmt"
	"strings"
)

func GetAllCities() (cityList map[int]entity.City, err error) {
	cityList = make(map[int]entity.City)
	rows, err := database.Master().Query("SELECT id, name FROM city")
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
	rows, err := database.Master().Query("SELECT COUNT(id) FROM city WHERE id = ?", cityId)
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

func GetUserCityList(userList []entity.User) (cityList map[int]entity.City, err error) {
	cityList = make(map[int]entity.City)
	if len(userList) > 0 {
		cityIdList := getUserCityIdList(userList)
		if len(cityIdList) > 0 {
			cityIdPlaceList := convert.IntListToQueryParameterPlaceList(cityIdList)
			cityIdListQuery := strings.Join(cityIdPlaceList, ",")
			sqlQuery := fmt.Sprintf("SELECT id, name FROM city WHERE id IN (%s)", cityIdListQuery)
			rows, err := database.Slave().Query(sqlQuery, convert.IntListToInterfaceList(cityIdList)...)
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
		}

	}
	return cityList, err
}

func getUserCityIdList(userList []entity.User) []int {
	var cityIdList []int
	if len(userList) > 0 {
		cityIdMap := make(map[int]int)
		for _, user := range userList {
			if user.CityId.Valid {
				cityId := int(user.CityId.Int64)
				cityIdMap[cityId] = cityId
			}
		}
		for _, cityId := range cityIdMap {
			cityIdList = append(cityIdList, cityId)
		}
	}
	return cityIdList
}

func GetUserCity(user entity.User) (city entity.City, err error) {
	cityList, err := GetUserCityList([]entity.User{user})
	if err != nil {
		return city, nil
	}
	cityId := int(user.CityId.Int64)
	city = cityList[cityId]
	return city, nil
}