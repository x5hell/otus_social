package controller

import (
	"component/template"
	"fmt"
	"net/http"
)

func UserProfileList(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenGuestTemplate("user-profile-list.html")
	if err != nil {
		fmt.Fprintf(response, "error: %v", err)
	} else {
		htmlTemplate.ExecuteTemplate(response, template.LayoutName, nil)
	}
}
