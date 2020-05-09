package model

import (
	"component/handler"
	"component/session"
	"component/session/memory"
	"net/http"
)

const UserIdName = "userId"

var sessionData session.Session
var sessionManager *session.Manager

func InitSession(response http.ResponseWriter, request *http.Request) {
	memory.Init()
	var err error
	sessionManager, err = session.NewManager("memory", "gosessionid", 3600)
	handler.ErrorLog(err)
	sessionData = sessionManager.SessionStart(response, request)
}

func GetSessionManager() *session.Manager {
	return sessionManager
}

func GetSessionData() session.Session {
	return sessionData
}

func GetUserIdFromSession() interface{} {
	return GetSessionData().Get(UserIdName)
}

func Authorized() bool {
	return GetUserIdFromSession() != nil
}

