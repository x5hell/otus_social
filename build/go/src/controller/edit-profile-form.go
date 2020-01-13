package controller

import (
	"component/controllerResponse"
	"component/template"
	"fmt"
	"model"
	"net/http"
)

func EditProfileForm(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenUserTemplate("edit-profile-form.html")
	if err != nil {
		controllerResponse.TemplateGeneratingError(response, err)
	} else {
		userId, userAuthorized := model.GetUserId().(int)
		if userAuthorized == false {
			controllerResponse.SessionExpiredError(response, fmt.Errorf(controllerResponse.SessionExpiredErrorMessage))
		} else {
			data, err := model.GetEditProfileFormData(userId)
			if err != nil {
				controllerResponse.GetTemplateDataError(response, err)
			} else {
				err = htmlTemplate.
					ExecuteTemplate(response, template.LayoutName, data)
				controllerResponse.TemplateFillError(response, err)
			}
		}
	}
}