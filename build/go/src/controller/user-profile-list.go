package controller

import (
	"component/template"
	"entity"
	"net/http"
	"repository"
)

const UserProfileListLimit = 10

type LastUserProfileLIst struct {
	Users map[int]entity.User
}

func UserProfileList(response http.ResponseWriter, request *http.Request)  {
	htmlTemplate, err := template.OpenGuestTemplate("user-profile-list.html")
	if err != nil {
		panic(err)
	}
	lastUsers, err := repository.GetLastUsers(UserProfileListLimit)
	if err != nil {
		panic(err)
	}
	data := LastUserProfileLIst{Users: lastUsers}
	err = htmlTemplate.ExecuteTemplate(response, template.LayoutName, data)
	if err != nil {
		panic(err)
	}
}
