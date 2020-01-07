package main

import (
	"controller"
	"log"
	"net/http"
)

const AppPort = "8001"

func main()  {
	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/index", controller.UserProfileList)
	mux.HandleFunc("/registration-form", controller.RegistrationForm)
	mux.HandleFunc("/registration", controller.Registration)
	mux.HandleFunc("/login-form", controller.LoginForm)
	mux.HandleFunc("/edit-profile-form", controller.EditProfileForm)
	log.Fatal(http.ListenAndServe("0.0.0.0:" + AppPort, mux))
}