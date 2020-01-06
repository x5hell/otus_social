package repository

import (
	"component/convert"
	"component/database"
	"entity"
	"fmt"
	"strings"
)

func GetAllInterests() (interestList []entity.Interest, err error) {
	rows, err := database.Query("SELECT id, name FROM interest")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		interest := entity.Interest{}
		err = rows.Scan(&interest.ID, &interest.Name)
		if err != nil {
			return interestList, err
		}
		interestList = append(interestList, interest)
	}
	return interestList, err
}

func GetInvalidInterestIds(interestIdList []string) (invalidInterestIdList []string, err error) {
	var interestIdPlaceList []string
	interestIdMap := make(map[string]string)
	for _, interestId := range interestIdList {
		interestIdMap[interestId] = interestId
		interestIdPlaceList = append(interestIdPlaceList, "?")
	}
	interestIdListQuery := strings.Join(interestIdPlaceList, ",")
	sqlQuery := fmt.Sprintf(
		"SELECT id FROM interest WHERE id IN (%s)", interestIdListQuery)
	rows, err := database.Query(sqlQuery, convert.StringListToInterfaceList(interestIdList)...)
	if err != nil {
		return invalidInterestIdList, err
	}
	var interestId string
	for rows.Next() {
		err = rows.Scan(&interestId)
		if err != nil {
			return invalidInterestIdList, err
		}
		delete(interestIdMap, interestId)
	}
	for _, interestId := range interestIdMap {
		invalidInterestIdList = append(invalidInterestIdList, interestId)
	}
	return invalidInterestIdList, nil
}
