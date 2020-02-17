package controller

import (
	"component/controllerResponse"
	"component/template"
	"model"
	"net/http"
)

func SearchForm(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenUserTemplate("search-form.html")
	if err != nil {
		controllerResponse.TemplateGeneratingError(response, err)
	} else {
		data, err := model.GetSearchFormData()
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
