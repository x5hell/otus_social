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
		fmt.Fprintf(response, "error: %v", err)
	} else {
		userId, userAuthorized := model.GetUserId().(int)
		if userAuthorized == false {
			fmt.Fprintf(response, controllerResponse.SessionExpiredMessage)
		} else {
			data, err := model.GetEditProfileFormData(userId)
			if err != nil {
				fmt.Fprintf(response, "error: %v", err)
			} else {
				htmlTemplate.
					ExecuteTemplate(response, template.LayoutName, data)
			}
		}

	}
}