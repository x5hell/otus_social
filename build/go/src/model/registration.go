package model

import (
	"component/database"
	"component/handler"
	"component/validation"
	"database/sql"
	"entity"
	"repository"
	"strconv"
)

const registrationFieldNamePassword2 = "password2"
const registrationFieldNameLogin = "login"
const registrationFieldNameInterests = "interests"
const registrationFieldNameCity = "city"
const registrationButton = "registration-button"

type RegistrationRequest struct {
	Login     string   `name:"registrationFieldNameLogin" validation:"required,symbolsMax=20"`
	Password1 string   `name:"password1" validation:"required,symbolsMin=6"`
	Password2 string   `name:"registrationFieldNamePassword2" validation:"required,symbolsMin=6"`
	FirstName string   `name:"first-name" validation:"required,symbolsMax=25,regex=^[а-яА-ЯёЁa-zA-Z\\-]+$"`
	LastName  string   `name:"last-name" validation:"required,symbolsMax=25,regex=^[а-яА-ЯёЁa-zA-Z\\-]+$"`
	Age       string   `name:"age" validation:"required,isInt,digitMin=1,digitMax=200"`
	Sex       string   `name:"sex" validation:"regex=^(1|2|)$"`
	Interests []string `name:"registrationFieldNameInterests"`
	City      string   `name:"registrationFieldNameCity"`
}

func Registration(requestStruct RegistrationRequest) (userId int, fieldErrors map[string]error) {
	validationResult := validateRegistrationRequest(requestStruct)
	if validationResult.ValidationResult {
		user := buildUserEntity(requestStruct)
		transaction, err := database.GetTransaction()
		if err != nil {
			return 0, FormFieldServerErrorWithRollback(transaction, validationResult, err, registrationButton)
		}
		err = repository.InsertUser(&user, transaction)
		if err != nil {
			return 0, FormFieldServerErrorWithRollback(transaction, validationResult, err, registrationButton)
		}
		userInterestEntityList := BuildUserInterestEntityList(requestStruct.Interests, user.ID)
		err = repository.InsertUserInterestEntityList(userInterestEntityList, transaction)
		if err != nil {
			return 0, FormFieldServerErrorWithRollback(transaction, validationResult, err, registrationButton)
		}
		err = GetSessionData().Set(UserIdName, user.ID)
		if err != nil {
			return 0, FormFieldServerErrorWithRollback(transaction, validationResult, err, registrationButton)
		}
		handler.ErrorLog(transaction.Commit())
		return user.ID, validationResult.FieldErrors
	} else {
		return 0, validationResult.FieldErrors
	}
}

func GetRegistrationFieldAliasList() map[string]string {
	return map[string]string{
		"registrationFieldNamePassword2": registrationFieldNamePassword2,
		"registrationFieldNameLogin":     registrationFieldNameLogin,
		"registrationFieldNameInterests": registrationFieldNameInterests,
		"registrationFieldNameCity":      registrationFieldNameCity,
	}
}


func validateRegistrationRequest(requestStruct RegistrationRequest) (result validation.FieldValidationResult) {
	fieldAliasList := GetRegistrationFieldAliasList()
	result.ValidationResult, result.FieldErrors = validation.ValidateStructure(requestStruct, fieldAliasList)
	result.ValidationResult =
		result.ValidationResult &&
			validation.ValidatePassword(requestStruct.Password1, requestStruct.Password2, registrationFieldNamePassword2, &result) &&
			validation.ValidateLogin(requestStruct.Login, registrationFieldNameLogin, &result) &&
			validation.ValidateInterests(requestStruct.Interests, registrationFieldNameInterests, &result) &&
			validation.ValidateCity(requestStruct.City, registrationFieldNameCity, &result)
	return result
}

func buildUserEntity(requestStruct RegistrationRequest) (user entity.User) {
	user.Login = requestStruct.Login
	user.Password = requestStruct.Password1
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

