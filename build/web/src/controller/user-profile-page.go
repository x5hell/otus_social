package controller

import (
	"component/controllerResponse"
	"component/template"
	"model"
	"net/http"
)

func UserProfilePage(response http.ResponseWriter, request *http.Request) {
	controllerResponse.ParseRequest(response, request, "GET", UserProfilePageAction)
}

func UserProfilePageAction(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenUserTemplate("user-profile-page.html")
	if err != nil {
		controllerResponse.TemplateGeneratingError(response, err)
	} else {
		requestStruct := CreateUserProfilePageRequest(request)

		data, err := model.GetUserProfilePageData(requestStruct)
		if err != nil {
			controllerResponse.GetTemplateDataError(response, err)
		}
		err = htmlTemplate.ExecuteTemplate(response, template.LayoutName, data)
		if err != nil {
			controllerResponse.TemplateFillError(response, err)
		}
	}
}

func CreateUserProfilePageRequest(request *http.Request) model.UserProfilePageRequest {
	var modelUserProfilePageRequest model.UserProfilePageRequest
	controllerResponse.FillStructureFromRequest(
		request, &modelUserProfilePageRequest, model.GetUserProfilePageRequestAliasList())
	return modelUserProfilePageRequest
}