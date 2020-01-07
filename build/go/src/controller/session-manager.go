package controller

import (
	"component/handler"
	"component/session"
	"net/http"
)

var sessionManager *session.Manager

func InitSession(response http.ResponseWriter, request *http.Request) {
	sessionManager, err := session.NewManager("memory","gosessionid",3600)
	handler.ErrorLog(err)
	sessionManager.SessionStart(response, request)

}

func GetSessionManager() *session.Manager {
	return sessionManager
}