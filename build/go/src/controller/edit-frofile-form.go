package controller

import (
	template2 "component/template"
	"fmt"
	"net/http"
)

func EditProfileForm(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template2.OpenUserTemplate("edit-profile-form.html")
	if err != nil {
		fmt.Fprintf(response, "error: %v", err)
	} else {
		htmlTemplate.ExecuteTemplate(response, template2.LayoutName, nil)
	}
}