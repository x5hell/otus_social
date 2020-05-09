package handler

import (
	"log"
	"runtime/debug"
)

func ErrorLog(err error) {
	if err != nil {
		log.Println(err)
		log.Println(string(debug.Stack()))
	}
}