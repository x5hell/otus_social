package controller

import (
	"model"
	"net/http"
)

func Middleware(
	handlerFunc func(response http.ResponseWriter, request *http.Request),
	middlewareList ...func(response http.ResponseWriter, request *http.Request) bool,
) func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		var useNextMiddleware bool
		for _, middleware := range middlewareList {
			useNextMiddleware = middleware(response, request)
			if useNextMiddleware == false {
				return
			}
		}
		if useNextMiddleware {
			handlerFunc(response, request)
		}
	}
}

func NotAuthorizedMiddleware(response http.ResponseWriter, request *http.Request) bool {
	if model.GetSessionData().Get(model.UserIdName) == nil {
		http.Redirect(response, request, "/login-form", 301)
		return false
	}
	return true
}

func AuthorizedMiddleware(response http.ResponseWriter, request *http.Request) bool {
	if model.GetSessionData().Get(model.UserIdName) != nil {
		http.Redirect(response, request, "/edit-profile-form", 301)
		return false
	}
	return true
}

func UseSession(response http.ResponseWriter, request *http.Request) bool {
	model.InitSession(response, request)
	return true
}