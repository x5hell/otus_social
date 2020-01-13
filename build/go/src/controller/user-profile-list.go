package controller

import (
	"component/controllerResponse"
	"component/template"
	"model"
	"net/http"
)

const UserProfileListLimit = 10

func UserProfileList(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenUserTemplate("user-profile-list.html")
	if err != nil {
		controllerResponse.TemplateGeneratingError(response, err)
	} else {
		data, err := model.GetUserProfileListData(UserProfileListLimit)
		if err != nil {
			controllerResponse.GetTemplateDataError(response, err)
		}
		err = htmlTemplate.ExecuteTemplate(response, template.LayoutName, data)
		if err != nil {
			controllerResponse.TemplateFillError(response, err)
		}
	}

}
