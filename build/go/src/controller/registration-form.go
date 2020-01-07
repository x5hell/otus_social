package controller

import (
	"component/template"
	"fmt"
	"model"
	"net/http"
)

func RegistrationForm(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenUserTemplate("registration-form.html")
	if err != nil {
		fmt.Fprintf(response, "error: %v", err)
	} else {
		data, err := model.GetRegistrationFormData()
		if err != nil {
			fmt.Fprintf(response, "error: %v", err)
		} else {
			htmlTemplate.
				ExecuteTemplate(response, template.LayoutName, data)
		}
	}
}