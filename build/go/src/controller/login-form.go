package controller

import (
	"component/controllerResponse"
	"component/template"
	"net/http"
)

func LoginForm(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenUserTemplate("login-form.html")
	if err != nil {
		controllerResponse.TemplateGeneratingError(response, err)
	} else {
		err = htmlTemplate.ExecuteTemplate(response, template.LayoutName, nil)
		if err != nil {
			controllerResponse.TemplateFillError(response, err)
		}
	}
}