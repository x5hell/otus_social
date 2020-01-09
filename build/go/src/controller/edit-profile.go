package controller

import (
	"component/controllerResponse"
	"fmt"
	"model"
	"net/http"
)

func EditProfile(response http.ResponseWriter, request *http.Request)  {
	controllerResponse.ParseRequest(response, request, "POST", editProfileAction)
}

func editProfileAction(response http.ResponseWriter, request *http.Request){
	requestStruct := createEditProfileRequest(request)
	fieldErrors := model.EditProfile(requestStruct)
	if len(fieldErrors) == 0 {
		controllerResponse.JsonOkMessage(fmt.Sprintf("%s", "ok"), response)
	} else {
		controllerResponse.JsonFormError(fieldErrors, response)
	}
}

func createEditProfileRequest(request *http.Request) model.EditProfileRequest {
	var modelRequest model.EditProfileRequest
	fieldAliasList := model.GetEditProfileFieldAliasList()
	controllerResponse.FillStructureFromRequest(request, &modelRequest, fieldAliasList)
	return modelRequest
}
