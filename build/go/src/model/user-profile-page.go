package model

import (
	"component/controllerResponse"
	"component/handler"
	"component/validation"
	"entity"
	"fmt"
	"repository"
	"strconv"
)

const UserProfilePageFieldNameUserId = "id"

type UserProfilePageRequest struct {
	UserId string   `name:"UserProfilePageFieldNameUserId" validation:"required,isInt,digitMin=1"`
}

type UserProfilePage struct {
	User         entity.User
	City         entity.City
	InterestList []entity.Interest
}

func GetUserProfilePageData(requestStruct UserProfilePageRequest) (data UserProfilePage, err error) {
	validationResult := ValidateUserProfilePageRequest(requestStruct)
	if validationResult {
		user, err := GetUserId(requestStruct.UserId)
		if err != nil {
			return data, err
		}
		data.User = user
		data.City, err = repository.GetUserCity(data.User)
		if err != nil {
			return data, err
		}
		data.InterestList, err = repository.GetUserInterestList(data.User)
		return data, err
	}
	return data, fmt.Errorf(controllerResponse.PageNotFoundErrorMessage)
}

func GetUserProfilePageRequestAliasList() map[string]string {
	return map[string]string{
		"UserProfilePageFieldNameUserId": UserProfilePageFieldNameUserId,
	}
}

func ValidateUserProfilePageRequest(requestStruct UserProfilePageRequest) (validationResult bool) {
	var result validation.FieldValidationResult
	result.ValidationResult, result.FieldErrors = validation.ValidateStructure(
		requestStruct, GetUserProfilePageRequestAliasList())
	if result.ValidationResult {
		return true
	} else {
		err, _ := result.FieldErrors[UserProfilePageFieldNameUserId]
		err = fmt.Errorf("параметр %s :: %s", UserProfilePageFieldNameUserId, err.Error())
		handler.ErrorLog(err)
		return  false
	}
}

func GetUserId(userIdParam string) (user entity.User, err error) {
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		return user, err
	}
	user, err = repository.GetUserById(userId)
	return user, err


}