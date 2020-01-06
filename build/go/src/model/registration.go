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

const FieldNamePassword2 = "password2"
const FieldNameLogin = "login"
const FieldNameInterests = "interests"
const FieldNameCity = "city"
const FieldRegistration = "registration"

type RegistrationRequest struct {
	Login     string   `name:"FieldNameLogin" validation:"required,symbolsMax=20"`
	Password1 string   `name:"password1" validation:"required,symbolsMin=6"`
	Password2 string   `name:"FieldNamePassword2" validation:"required,symbolsMin=6"`
	FirstName string   `name:"first-name" validation:"required,symbolsMax=25,regex=^[а-яА-ЯёЁa-zA-Z\\-]+$"`
	LastName  string   `name:"last-name" validation:"required,symbolsMax=25,regex=^[а-яА-ЯёЁa-zA-Z\\-]+$"`
	Age       string   `name:"age" validation:"required,isInt,digitMin=1,digitMax=200"`
	Sex       string   `name:"sex" validation:"regex=^(1|2|)$"`
	Interests []string `name:"FieldNameInterests"`
	City      string   `name:"FieldNameCity"`
}

type RegistrationModel struct {
	RequestStruct    RegistrationRequest
}

func Registration(requestStruct RegistrationRequest) (userId int, fieldErrors map[string]error) {
	validationResult := validateRegistrationRequest(requestStruct)
	if validationResult.ValidationResult {
		user := buildUserEntity(requestStruct)
		err := repository.InsertUser(&user)
		if err != nil {
			validationResult.FieldErrors[FieldRegistration] = fmt.Errorf(controllerResponse.ServerErrorMessage)
			handler.ErrorLog(err)
		}
		return user.ID, validationResult.FieldErrors
	} else {
		return 0, validationResult.FieldErrors
	}
}

func GetRegistrationFieldAliasList() map[string]string {
	return map[string]string{
		"FieldNamePassword2": FieldNamePassword2,
		"FieldNameLogin":     FieldNameLogin,
		"FieldNameInterests": FieldNameInterests,
		"FieldNameCity":      FieldNameCity,
	}
}


func validateRegistrationRequest(requestStruct RegistrationRequest) (result validation.FieldValidationResult) {
	fieldAliasList := GetRegistrationFieldAliasList()
	result.ValidationResult, result.FieldErrors = validation.ValidateStructure(requestStruct, fieldAliasList)
	result.ValidationResult =
		result.ValidationResult &&
			validatePassword(requestStruct.Password1, requestStruct.Password2, FieldNamePassword2, &result) &&
			validateLogin(requestStruct.Login, FieldNameLogin, &result) &&
			validateInterests(requestStruct.Interests, FieldNameInterests, &result) &&
			validateCity(requestStruct.City, FieldNameCity, &result)
	return result
}

func validatePassword(password1, password2, fieldName string, validationResult *validation.FieldValidationResult) bool {
	if password1 != password2 {
		validationResult.FieldErrors[fieldName] = fmt.Errorf("пароли не совпадают")
		return false
	}
	return true
}

func validateLogin(login, fieldName string, validationResult *validation.FieldValidationResult) bool {
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

func validateInterests(interests []string, fieldName string, validationResult *validation.FieldValidationResult) bool {
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

func validateCity(city, fieldName string, validationResult *validation.FieldValidationResult) bool {
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

func buildUserEntity(requestStruct RegistrationRequest) (user entity.User) {
	user.Login = requestStruct.Login
	user.Password = requestStruct.Password1
	user.FirstName = requestStruct.FirstName
	user.LastName = requestStruct.LastName
	age, err := strconv.Atoi(requestStruct.Age)
	handler.ErrorLog(err)
	user.Age = age
	user.Sex = requestStruct.Sex
	user.InterestList = []entity.Interest{}
	for _, interestId := range requestStruct.Interests {
		interestIdInt, err := strconv.Atoi(interestId)
		handler.ErrorLog(err)
		user.InterestList = append(user.InterestList, entity.Interest{ID: interestIdInt})
	}
	if len(requestStruct.City) > 0 {
		cityId, err := strconv.Atoi(requestStruct.City)
		handler.ErrorLog(err)
		user.City = entity.City{ID: cityId}
	}
	return user
}