package main

import (
	"controller"
	"log"
	"net/http"
)

const AppPort = "8001"

func main()  {
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", controller.UserProfileList)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/registration-form", controller.RegistrationForm)
	http.HandleFunc("/registration", controller.Registration)
	http.HandleFunc("/login-form", controller.LoginForm)
	http.HandleFunc("/edit-profile-form", controller.EditProfileForm)
	log.Fatal(http.ListenAndServe("0.0.0.0:" + AppPort, nil))
}