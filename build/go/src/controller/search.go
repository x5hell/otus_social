package controller

import (
	"component/controllerResponse"
	"encoding/json"
	"model"
	"net/http"
)

func Search(response http.ResponseWriter, request *http.Request)  {
	controllerResponse.ParseRequest(response, request, "GET", searchAction)
}

func searchAction(response http.ResponseWriter, request *http.Request){
	requestStruct := createSearchRequest(request)
	searchResult, fieldErrors := model.Search(requestStruct)
	if len(fieldErrors) == 0 {
		jsonSearchResult, err := json.Marshal(searchResult)
		if err == nil {
			controllerResponse.JsonOkMessage(string(jsonSearchResult), response)
		} else {
			controllerResponse.JsonErrorMessage(err.Error(), response)
		}
	} else {
		controllerResponse.JsonFormError(fieldErrors, response)
	}
}

func createSearchRequest(request *http.Request) model.SearchRequest {
	var modelSearchRequest model.SearchRequest
	fieldAliasList := model.GetRegistrationFieldAliasList()
	controllerResponse.FillStructureFromRequest(request, &modelSearchRequest, fieldAliasList)
	return modelSearchRequest
}
