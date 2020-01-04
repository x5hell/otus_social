package controller

import (
	"component/controllerResponse"
	"component/validation"
	"fmt"
	"net/http"
	"repository"
)

const FieldNamePassword2 = "password2"
const FieldNameLogin = "login"
const FieldNameInterests = "interests"
const FieldNameCity = "city"

type RegistrationController struct {
	request http.Request
	response http.ResponseWriter
	requestStruct RegistrationRequest
	validationResult bool
	fieldErrors map[string]error
}

type RegistrationRequest struct {
	Login string `name:"FieldNameLogin" validation:"required,symbolsMax=20"`
	Password1 string `name:"password1" validation:"required,symbolsMin=6"`
	Password2 string `name:"FieldNamePassword2" validation:"required,symbolsMin=6"`
	FirstName string `name:"first-name" validation:"required,symbolsMax=25,regex=^[а-яА-ЯёЁa-zA-Z\\-]+$"`
	LastName string `name:"last-name" validation:"required,symbolsMax=25,regex=^[а-яА-ЯёЁa-zA-Z\\-]+$"`
	Age string `name:"age" validation:"required,isInt,digitMin=1,digitMax=200"`
	Sex string `name:"sex" validation:"regex=^(1|2)$"`
	Interests []string `name:"FieldNameInterests"`
	City string `name:"FieldNameCity"`
}

func Registration(response http.ResponseWriter, request *http.Request)  {
	controllerResponse.ParseRequest(response, request, "POST", registrationAction)
}

func registrationAction(response http.ResponseWriter, request *http.Request){
	controller := RegistrationController{
		request:       *request,
		response:      response,
	}
	controller.
		createRequestStruct().
		validateRegistrationRequest()
	if controller.validationResult {
		_, _ = fmt.Fprintf(response, "ok :: %v", controller.requestStruct)
	} else {
		controllerResponse.JsonFormError(controller.fieldErrors, response)
	}
}

func (controller *RegistrationController) createRequestStruct() *RegistrationController {
	fieldAliasList := map[string]string{
		"FieldNamePassword2": FieldNamePassword2,
		"FieldNameLogin": FieldNameLogin,
		"FieldNameInterests": FieldNameInterests,
		"FieldNameCity": FieldNameCity,
	}
	controllerResponse.FillStructureFromRequest(&controller.request, &controller.requestStruct, fieldAliasList)
	return controller
}

func (controller *RegistrationController) validateRegistrationRequest() *RegistrationController {
	controller.validationResult, controller.fieldErrors = validation.ValidateStructure(controller.requestStruct)
	controller.validationResult =
		controller.validationResult &&
		controller.validatePassword() &&
		controller.validateLogin() &&
		controller.validateInterests() &&
		controller.validateCity()
	return controller
}

func (controller *RegistrationController) validatePassword() bool {
	if controller.requestStruct.Password1 != controller.requestStruct.Password2 {
		controller.fieldErrors[FieldNamePassword2] = fmt.Errorf("passwords do not match")
		return false
	}
	return true
}

func (controller *RegistrationController) validateLogin() bool {
	loginExists, err := repository.LoginExists(controller.requestStruct.Login)
	if err != nil {
		controller.fieldErrors[FieldNameLogin] = fmt.Errorf(controllerResponse.ServerErrorMessage)
		return false
	}
	if loginExists {
		controller.fieldErrors[FieldNameLogin] = fmt.Errorf("login already exists")
	}
	return loginExists == false
}

func (controller *RegistrationController) validateInterests() bool {
	if len(controller.requestStruct.Interests) == 0 {
		return true
	}
	invalidInterestIdsList, err := repository.GetInvalidInterestIds(controller.requestStruct.Interests)
	if err != nil {
		controller.fieldErrors[FieldNameInterests] = fmt.Errorf(controllerResponse.ServerErrorMessage)
		return false
	}
	if len(invalidInterestIdsList) > 0 {
		controller.fieldErrors[FieldNameInterests] = fmt.Errorf(
			"incorrect interest ids: %v", invalidInterestIdsList)
		return false
	}
	return true
}

func (controller *RegistrationController) validateCity() bool {
	if len(controller.requestStruct.City) == 0 {
		return true
	}
	cityIdExists, err := repository.CheckCityIdExists(controller.requestStruct.City)
	if err != nil {
		controller.fieldErrors[FieldNameCity] = fmt.Errorf(controllerResponse.ServerErrorMessage)
		return false
	}
	if cityIdExists == false {
		controller.fieldErrors[FieldNameCity] = fmt.Errorf(
			"incorrect city id: %s", controller.requestStruct.City)
		return false
	}
	return true
}