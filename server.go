package main

import (
	"awesomeProject/cfg"
	"awesomeProject/router"
	"log"
	"net/http"
)

func Init() {
	cfg.InitializeConfigurations()
	router.InitializeRouter()
}

func main() {
	Init()
	log.Fatal(http.ListenAndServe(":8080", router.Router))
}
