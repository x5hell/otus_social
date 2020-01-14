package controller

import (
	"component/controllerResponse"
	"component/template"
	"model"
	"net/http"
)

func RegistrationForm(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenUserTemplate("registration-form.html")
	if err != nil {
		controllerResponse.TemplateGeneratingError(response, err)
	} else {
		data, err := model.GetRegistrationFormData()
		if err != nil {
			controllerResponse.GetTemplateDataError(response, err)
		} else {
			err = htmlTemplate.ExecuteTemplate(response, template.LayoutName, data)
			if err != nil {
				controllerResponse.TemplateFillError(response, err)
			}
		}
	}
}