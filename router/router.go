package router

import (
	"awesomeProject/controllers"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func InitializeRouter() {
	Router = mux.NewRouter()
	Router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	Router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	Router.HandleFunc("/sign-up", controllers.SignUp).Methods("POST")
	Router.HandleFunc("/sign-in", controllers.SignIn).Methods("POST")
	Router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	Router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
