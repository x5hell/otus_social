package main

import (
	controller2 "controller"
	"log"
	"net/http"
)

const AppPort  = "8001"

func main()  {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/registration-form", controller2.RegistrationForm)
	http.HandleFunc("/login-form", controller2.LoginForm)
	http.HandleFunc("/edit-profile-form", controller2.EditProfileForm)
	log.Fatal(http.ListenAndServe("0.0.0.0:" + AppPort, nil))
}