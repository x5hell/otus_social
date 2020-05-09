package controller

import (
	"model"
	"net/http"
)

func Logout(response http.ResponseWriter, request *http.Request)  {
	model.GetSessionManager().SessionDestroy(response, request)
	response.Header().Add("Cache-Control", "no-cache")
	http.Redirect(response, request, "/login-form", 301)
}
