package main

import (
	"awesomeProject/cfg"
	"awesomeProject/jwt"
	"awesomeProject/models"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var users []models.User

var jwtConfig cfg.JwtConfig
var dbConfig cfg.DBConfig

func getSecret() []byte {
	return []byte(jwtConfig.Secret)
}

func main() {
	router := mux.NewRouter()

	cfg.ReadJwtProperties("./properties/jwt-configuration.json", &jwtConfig)
	cfg.ReadDBProperties("./properties/db-configuration.json", &dbConfig)

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	token := jwt.GenerateToken(
		jwtConfig.Issuer,
		getSecret(),
		time.Now(),
		jwtConfig.BaseTokenDuration,
		uuid.New().String(),
		"aud")
	fmt.Println(token)

	parsedClaims, err := jwt.ParseToken(token, getSecret())
	if err == nil {
		aud := parsedClaims["aud"]
		sub := parsedClaims["sub"]
		fmt.Println("aud:", aud)
		fmt.Println("sub:", sub)
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getUsers(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get all Users")
}

func getUser(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get single User")
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	log.Println("Create User")
}

func updateUser(writer http.ResponseWriter, request *http.Request) {
	log.Println("Update User")
}

func deleteUser(writer http.ResponseWriter, request *http.Request) {
	log.Println("Delete User")
}
