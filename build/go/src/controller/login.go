package controller

import (
	"component/controllerResponse"
	"fmt"
	"model"
	"net/http"
)

func Login(response http.ResponseWriter, request *http.Request)  {
	controllerResponse.ParseRequest(response, request, "POST", loginAction)
}

func loginAction(response http.ResponseWriter, request *http.Request){
	requestStruct := createLoginRequest(request)
	userId, fieldErrors := model.Login(requestStruct)
	if userId > 0 {
		controllerResponse.JsonOkMessage(fmt.Sprintf("%d", userId), response)
	} else {
		controllerResponse.JsonFormError(fieldErrors, response)
	}
}

func createLoginRequest(request *http.Request) model.LoginRequest {
	var modelLoginRequest model.LoginRequest
	controllerResponse.FillStructureFromRequest(request, &modelLoginRequest, map[string]string{})
	return modelLoginRequest
}
