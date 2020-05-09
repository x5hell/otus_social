package controller

import (
	"component/controllerResponse"
	"component/database"
	"model"
	"net/http"
)

func ChangeDbWorkMode(response http.ResponseWriter, request *http.Request)  {
	controllerResponse.ParseRequest(response, request, "GET", changeDbWorkModeAction)
}

func changeDbWorkModeAction(response http.ResponseWriter, request *http.Request){
	requestStruct := createChangeDbWorkModeRequest(request)
	err := model.ChangeDbWorkMode(requestStruct)

	if err == nil {
		workMode, _ := database.GetWorkMode()
		controllerResponse.JsonOkMessage(workMode, response)
	} else {
		controllerResponse.JsonFormError(map[string]error{model.WorkModeFieldName: err}, response)
	}
}

func createChangeDbWorkModeRequest(request *http.Request) model.ChangeDbWorkModeRequest {
	var modelRequest model.ChangeDbWorkModeRequest
	controllerResponse.FillStructureFromRequest(request, &modelRequest, map[string]string{})
	return modelRequest
}