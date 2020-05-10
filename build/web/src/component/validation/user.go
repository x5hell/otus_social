package validation

import (
	"component/controllerResponse"
	"fmt"
	"repository"
)

func ValidatePassword(password1, password2, fieldName string, validationResult *FieldValidationResult) bool {
	if password1 != password2 {
		validationResult.FieldErrors[fieldName] = fmt.Errorf("пароли не совпадают")
		return false
	}
	return true
}

func ValidateLogin(login, fieldName string, validationResult *FieldValidationResult) bool {
	loginExists, err := repository.LoginExists(login)
	if err != nil {
		validationResult.FieldErrors[fieldName] = fmt.Errorf(controllerResponse.ServerErrorMessage)
		return false
	}
	if loginExists {
		validationResult.FieldErrors[fieldName] = fmt.Errorf("логин уже существует")
	}
	return loginExists == false
}

func ValidateInterests(interests []string, fieldName string, validationResult *FieldValidationResult) bool {
	if len(interests) == 0 {
		return true
	}
	invalidInterestIdsList, err := repository.GetInvalidInterestIds(interests)
	if err != nil {
		validationResult.FieldErrors[fieldName] = fmt.Errorf(controllerResponse.ServerErrorMessage)
		return false
	}
	if len(invalidInterestIdsList) > 0 {
		validationResult.FieldErrors[fieldName] = fmt.Errorf(
			"некорректный список идентификаторов интересов: %v", invalidInterestIdsList)
		return false
	}
	return true
}

func ValidateCity(city, fieldName string, validationResult *FieldValidationResult) bool {
	if len(city) == 0 {
		return true
	}
	cityIdExists, err := repository.CheckCityIdExists(city)
	if err != nil {
		validationResult.FieldErrors[fieldName] = fmt.Errorf(controllerResponse.ServerErrorMessage)
		return false
	}
	if cityIdExists == false {
		validationResult.FieldErrors[fieldName] = fmt.Errorf(
			"некорректный идентификатор города: %s", city)
		return false
	}
	return true
}

