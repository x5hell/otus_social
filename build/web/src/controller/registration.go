package controller

import (
	"component/controllerResponse"
	"model"
	"net/http"
	"strconv"
)

func Registration(response http.ResponseWriter, request *http.Request)  {
	controllerResponse.ParseRequest(response, request, "POST", registrationAction)
}

func registrationAction(response http.ResponseWriter, request *http.Request){
	requestStruct := createRegistrationRequest(request)
	userId, fieldErrors := model.Registration(requestStruct)
	if userId > 0 {
		controllerResponse.JsonOkMessage(strconv.Itoa(userId), response)
	} else {
		controllerResponse.JsonFormError(fieldErrors, response)
	}
}

func createRegistrationRequest(request *http.Request) model.RegistrationRequest {
	var modelRegistrationRequest model.RegistrationRequest
	fieldAliasList := model.GetRegistrationFieldAliasList()
	controllerResponse.FillStructureFromRequest(request, &modelRegistrationRequest, fieldAliasList)
	return modelRegistrationRequest
}
