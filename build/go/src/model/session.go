package model

import (
	"component/handler"
	"component/session"
	"component/session/memory"
	"net/http"
)

const UserIdName = "userId"

var sessionData session.Session

func InitSession(response http.ResponseWriter, request *http.Request) {
	memory.Init()
	sessionManager, err := session.NewManager("memory", "gosessionid", 3600)
	handler.ErrorLog(err)
	sessionData = sessionManager.SessionStart(response, request)
}

func GetSessionData() session.Session {
	return sessionData
}

func GetUserId() interface{} {
	return GetSessionData().Get(UserIdName)
}

func Authorized() bool {
	return GetUserId() != nil
}

