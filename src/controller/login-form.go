package controller

import (
	"../component/template"
	"fmt"
	"net/http"
)

func LoginForm(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenGuestTemplate("login-form.html")
	if err != nil {
		fmt.Fprintf(response, "error: %v", err)
	} else {
		htmlTemplate.ExecuteTemplate(response, template.LayoutName, nil)
	}
}