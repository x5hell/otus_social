package main

import (
	"component/database"
	"component/environment"
	"controller"
	"log"
	"net/http"
)

const EnvSiteInternalPort = "WEB_SITE_INTERNAL_PORT"

const EnvDbDatabase = "MYSQL_DATABASE"
const EnvDbPort = "MYSQL_PORT"
const EnvMasterHost = "MYSQL_MASTER_HOST"
const EnvMasterUser = "MYSQL_MASTER_USER"
const EnvMasterPassword = "MYSQL_MASTER_PASSWORD"
const EnvSlaveHost = "MYSQL_SLAVE_HOST"
const EnvSlaveUser = "MYSQL_SLAVE_USER"
const EnvSlavePassword = "MYSQL_SLAVE_PASSWORD"

func main()  {
	initRegistryConnections()
	defer database.CloseRegistryConnections()
	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/user-profile-list", controller.Middleware(
		controller.UserProfileList, controller.UseSession))

	mux.HandleFunc("/search-form", controller.Middleware(
		controller.SearchForm, controller.UseSession))

	mux.HandleFunc("/search", controller.Middleware(
		controller.Search, controller.UseSession))

	mux.HandleFunc("/user-profile-page", controller.Middleware(
		controller.UserProfilePage, controller.UseSession))

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
	mux.HandleFunc("/change-db-work-mode", controller.Middleware(
		controller.ChangeDbWorkMode, controller.UseSession))

	mux.HandleFunc("/", controller.Middleware(
		nil, controller.RedirectMainPage))

	log.Fatal(http.ListenAndServe("0.0.0.0:" + getPort(), mux))
}

func initRegistryConnections() {
	masterEnv := database.ConnectionSettingsEnvironment{
		Host:     EnvMasterHost,
		Port:     EnvDbPort,
		Database: EnvDbDatabase,
		User:     EnvMasterUser,
		Password: EnvMasterPassword,
	}
	master := database.Connection{Environment:masterEnv}
	slaveEnv := database.ConnectionSettingsEnvironment{
		Host:     EnvSlaveHost,
		Port:     EnvDbPort,
		Database: EnvDbDatabase,
		User:     EnvSlaveUser,
		Password: EnvSlavePassword,
	}
	slave := database.Connection{Environment:slaveEnv}
	database.InitConnectionRegistry(master, slave)
}

func getPort() string {
	port, err := environment.Get(EnvSiteInternalPort)
	if err != nil {
		log.Fatal(err)
	}
	return port
}