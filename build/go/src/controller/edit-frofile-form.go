package controller

import (
	"component/template"
	"fmt"
	"net/http"
)

func EditProfileForm(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenUserTemplate("edit-profile-form.html")
	if err != nil {
		fmt.Fprintf(response, "error: %v", err)
	} else {
		htmlTemplate.ExecuteTemplate(response, template.LayoutName, nil)
	}
}