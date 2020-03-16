package controller

import (
	"component/controllerResponse"
	"model"
	"net/http"
	"repository"
)


func Search(response http.ResponseWriter, request *http.Request)  {
	controllerResponse.ParseRequest(response, request, "GET", searchAction)
}

func searchAction(response http.ResponseWriter, request *http.Request){
	requestStruct := createSearchRequest(request)
	searchResult, fieldErrors := model.Search(requestStruct)
	if len(fieldErrors) == 0 {
		controllerResponse.JsonOkData(searchResult, response)
	} else {
		controllerResponse.JsonFormError(fieldErrors, response)
	}
}

func createSearchRequest(request *http.Request) repository.UserSearchRequest {
	var modelSearchRequest repository.UserSearchRequest
	fieldAliasList := model.GetSearchFieldAliasList()
	controllerResponse.FillStructureFromRequest(request, &modelSearchRequest, fieldAliasList)
	return modelSearchRequest
}
