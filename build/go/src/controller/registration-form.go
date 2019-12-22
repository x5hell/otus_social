package controller

import (
	template2 "component/template"
	"fmt"
	"net/http"
)

func RegistrationForm(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template2.OpenGuestTemplate("registration-form.html")
	if err != nil {
		fmt.Fprintf(response, "error: %v", err)
	} else {
		htmlTemplate.ExecuteTemplate(response, template2.LayoutName, nil)
	}
}