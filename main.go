package main

import (
	"./controller"
	"log"
	"net/http"
)

const AppPort  = "8000"

func main()  {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/registration-form", controller.RegistrationForm)
	http.HandleFunc("/login-form", controller.LoginForm)
	http.HandleFunc("/edit-profile-form", controller.EditProfileForm)
	log.Fatal(http.ListenAndServe("localhost:" + AppPort, nil))
}