package model

import (
	"component/controllerResponse"
	"component/handler"
	"component/validation"
	"fmt"
	"repository"
)

const fieldLoginButton = "login-button"

type LoginRequest struct {
	Login    string `name:"login" validation:"required"`
	Password string `name:"password" validation:"required"`
}

func Login(requestStruct LoginRequest) (userId int, fieldErrors map[string]error) {
	validationResult := validateLoginRequest(requestStruct)
	userId = checkAuth(requestStruct, &validationResult)
	if validationResult.ValidationResult {
		err := GetSessionData().Set(UserIdName, userId)
		if err != nil {
			validationResult.FieldErrors[fieldLoginButton] = fmt.Errorf(controllerResponse.ServerErrorMessage)
			handler.ErrorLog(err)
		}
		return userId, validationResult.FieldErrors
	} else {
		return 0, validationResult.FieldErrors
	}
}

func validateLoginRequest(requestStruct LoginRequest) (result validation.FieldValidationResult) {
	result.ValidationResult, result.FieldErrors = validation.ValidateStructure(requestStruct, map[string]string{})
	return result
}

func checkAuth(requestStruct LoginRequest, validationResult *validation.FieldValidationResult) (userId int) {
	if validationResult.ValidationResult {
		user, err := repository.GetUserByAuth(requestStruct.Login, requestStruct.Password)
		if err == nil {
			return user.ID
		}
		validationResult.FieldErrors[fieldLoginButton] = err
		validationResult.ValidationResult = false
	}
	return 0
}