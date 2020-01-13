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

	mux.HandleFunc("/user-profile-list", controller.Middleware(
		controller.UserProfileList, controller.UseSession))

	mux.HandleFunc("/registration-form", controller.Middleware(
		controller.RegistrationForm, controller.UseSession, controller.AuthorizedMiddleware))
	mux.HandleFunc("/registration", controller.Middleware(
		controller.Registration, controller.UseSession))

	mux.HandleFunc("/login-form", controller.Middleware(
		controller.LoginForm, controller.UseSession, controller.AuthorizedMiddleware))
	mux.HandleFunc("/login", controller.Middleware(
		controller.Login, controller.UseSession))

	mux.HandleFunc("/logout", controller.Middleware(
		controller.Logout, controller.UseSession))

	mux.HandleFunc("/edit-profile-form", controller.Middleware(
		controller.EditProfileForm, controller.UseSession, controller.NotAuthorizedMiddleware))
	mux.HandleFunc("/edit-profile", controller.Middleware(
		controller.EditProfile, controller.UseSession))

	log.Fatal(http.ListenAndServe("0.0.0.0:" + AppPort, mux))
}