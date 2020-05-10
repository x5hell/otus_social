package controller

import (
	"component/controllerResponse"
	"model"
	"net/http"
	"strconv"
)

func Login(response http.ResponseWriter, request *http.Request)  {
	controllerResponse.ParseRequest(response, request, "POST", loginAction)
}

func loginAction(response http.ResponseWriter, request *http.Request){
	requestStruct := createLoginRequest(request)
	userId, fieldErrors := model.Login(requestStruct)
	if userId > 0 {
		controllerResponse.JsonOkMessage(strconv.Itoa(userId), response)
	} else {
		controllerResponse.JsonFormError(fieldErrors, response)
	}
}

func createLoginRequest(request *http.Request) model.LoginRequest {
	var modelLoginRequest model.LoginRequest
	controllerResponse.FillStructureFromRequest(request, &modelLoginRequest, map[string]string{})
	return modelLoginRequest
}
