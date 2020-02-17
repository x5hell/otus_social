package model

import (
	"component/validation"
	"repository"
)

const searchFieldNameInterests = "interests"
const searchFieldNameCity = "city"
const searchUserLimit = 1000

type SearchRequest struct {
	FirstName string   `name:"first-name" validation:"symbolsMax=25,regex=^[а-яА-ЯёЁa-zA-Z\\-]+$"`
	LastName  string   `name:"last-name" validation:"symbolsMax=25,regex=^[а-яА-ЯёЁa-zA-Z\\-]+$"`
	AgeFrom   string   `name:"age-from" validation:"isInt,digitMin=1,digitMax=200"`
	AgeTo     string   `name:"age-to" validation:"isInt,digitMin=1,digitMax=200"`
	Sex       string   `name:"sex" validation:"regex=^(1|2|)$"`
	Interests []string `name:"searchFieldNameInterests"`
	City      string   `name:"searchFieldNameCity" validation:"isInt,digitMin=1"`
}

func Search(requestStruct SearchRequest) (userProfileList UserProfileList, fieldErrors map[string]error) {
	validationResult := validateSearchRequest(requestStruct)
	if validationResult.ValidationResult {
		userProfileList, _ := searchUsersProfileListData(requestStruct, searchUserLimit)
		return userProfileList, nil
	} else {
		return userProfileList, validationResult.FieldErrors
	}
}

func validateSearchRequest(requestStruct SearchRequest) (result validation.FieldValidationResult) {
	fieldAliasList := GetRegistrationFieldAliasList()
	result.ValidationResult, result.FieldErrors = validation.ValidateStructure(requestStruct, fieldAliasList)
	result.ValidationResult =
		result.ValidationResult &&
			validation.ValidateInterests(requestStruct.Interests, searchFieldNameInterests, &result) &&
			validation.ValidateCity(requestStruct.City, searchFieldNameCity, &result)
	return result
}

func searchUsersProfileListData(requestStruct SearchRequest, limit int) (data UserProfileList, err error) {
	lastUsers, err := repository.SearchUsers(requestStruct, limit)
	if err != nil {
		return data, err
	}
	return GetUserProfileListData(lastUsers)
}