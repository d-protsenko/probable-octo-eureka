package main

import (
	"awesomeProject/cfg"
	"awesomeProject/repositories"
	"awesomeProject/router"
	"log"
	"net/http"
)

func Init() {
	log.Println("Initializing configurations...")
	cfg.InitializeConfigurations()
	log.Println("JWT, DB, Cookie configurations initialized.")
	log.Println("Initializing router....")
	router.InitializeRouter()
	log.Println("Router initialized.")
}

func main() {
	Init()
	log.Println("Connecting to DB...")
	repositories.ConnectToDb()
	log.Println("Connected to DB")
	log.Fatal(http.ListenAndServe(":8080", router.Router))
}
