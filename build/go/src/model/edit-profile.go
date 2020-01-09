package model

import (
	"component/controllerResponse"
	"component/handler"
	"component/validation"
	"database/sql"
	"entity"
	"fmt"
	"repository"
	"strconv"
)

const editProfileFieldNameInterests = "interests"
const editProfileFieldNameCity = "city"
const editProfileButton = "edit-profile-button"

type EditProfileRequest struct {
	FirstName string   `name:"first-name" validation:"required,symbolsMax=25,regex=^[а-яА-ЯёЁa-zA-Z\\-]+$"`
	LastName  string   `name:"last-name" validation:"required,symbolsMax=25,regex=^[а-яА-ЯёЁa-zA-Z\\-]+$"`
	Age       string   `name:"age" validation:"required,isInt,digitMin=1,digitMax=200"`
	Sex       string   `name:"sex" validation:"regex=^(1|2|)$"`
	Interests []string `name:"editProfileFieldNameInterests"`
	City      string   `name:"editProfileFieldNameCity"`
}

func EditProfile(requestStruct EditProfileRequest) (fieldErrors map[string]error) {
	validationResult := validateEditProfileRequest(requestStruct)
	if validationResult.ValidationResult {
		userId, userAuthorized := GetUserId().(int)
		if userAuthorized == false {
			validationResult.FieldErrors[editProfileButton] = fmt.Errorf(controllerResponse.SessionExpiredMessage)
			return validationResult.FieldErrors
		}
		user, err := repository.GetUserById(userId)
		if err != nil {
			validationResult.FieldErrors[editProfileButton] = fmt.Errorf(controllerResponse.ServerErrorMessage)
			handler.ErrorLog(err)
			return validationResult.FieldErrors
		}
		user = changeUserEntity(requestStruct, user)
		err = repository.UpdateUser(&user)
		if err != nil {
			validationResult.FieldErrors[editProfileButton] = fmt.Errorf(controllerResponse.ServerErrorMessage)
			handler.ErrorLog(err)
		}
		return validationResult.FieldErrors
	} else {
		return validationResult.FieldErrors
	}
}

func GetEditProfileFieldAliasList() map[string]string {
	return map[string]string{
		"editProfileFieldNameInterests": editProfileFieldNameInterests,
		"editProfileFieldNameCity":      editProfileFieldNameCity,
	}
}

func validateEditProfileRequest(requestStruct EditProfileRequest) (result validation.FieldValidationResult) {
	fieldAliasList := GetEditProfileFieldAliasList()
	result.ValidationResult, result.FieldErrors = validation.ValidateStructure(requestStruct, fieldAliasList)
	result.ValidationResult =
		result.ValidationResult &&
			validation.ValidateInterests(requestStruct.Interests, editProfileFieldNameInterests, &result) &&
			validation.ValidateCity(requestStruct.City, editProfileFieldNameCity, &result)
	return result
}

func changeUserEntity(requestStruct EditProfileRequest, user entity.User) entity.User {
	user.FirstName = requestStruct.FirstName
	user.LastName = requestStruct.LastName
	age, err := strconv.Atoi(requestStruct.Age)
	handler.ErrorLog(err)
	user.Age = age
	user.Sex = sql.NullString{String: requestStruct.Sex, Valid: len(requestStruct.Sex) > 0}
	user.CityId = sql.NullInt64{Int64: 0, Valid: false}
	if len(requestStruct.City) > 0 {
		cityId, err := strconv.Atoi(requestStruct.City)
		handler.ErrorLog(err)
		if err == nil {
			user.CityId = sql.NullInt64{Int64: int64(cityId), Valid: true}
		}
	}
	return user
}