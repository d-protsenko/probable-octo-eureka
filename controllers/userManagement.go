package controllers

import (
	"awesomeProject/cfg"
	"awesomeProject/cookie"
	"awesomeProject/jwt"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

var accessJWTCookieName = "jwtA"
var refreshJWTCookieName = "jwtR"
var accessTokenCookieKey = "access"
var refreshTokenCookieKey = "refresh"

func getSecret() []byte {
	return []byte(cfg.MainJwtConfig.Secret)
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get all Users")
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get single User")
}

func SignUp(writer http.ResponseWriter, request *http.Request) {
	log.Println("Create User")
}

func SignIn(writer http.ResponseWriter, request *http.Request) {
	userId := uuid.New().String()
	currentTime := time.Now()
	accessToken := jwt.GenerateToken(
		cfg.MainJwtConfig.Issuer,
		getSecret(),
		currentTime,
		cfg.MainJwtConfig.AccessTokenDuration,
		userId,
		"")
	refreshToken := jwt.GenerateToken(
		cfg.MainJwtConfig.Issuer,
		getSecret(),
		currentTime,
		cfg.MainJwtConfig.RefreshTokenDuration,
		userId,
		"")
	cookie.SetCookie(accessJWTCookieName, accessTokenCookieKey, accessToken, int(cfg.MainJwtConfig.AccessTokenDuration), writer)
	cookie.SetCookie(refreshJWTCookieName, refreshTokenCookieKey, refreshToken, int(cfg.MainJwtConfig.RefreshTokenDuration), writer)
	log.Println("User successfully signed in and A/R tokens set")
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	accessToken := cookie.ReadCookie(accessJWTCookieName, accessTokenCookieKey, request)
	refreshToken := cookie.ReadCookie(refreshJWTCookieName, refreshTokenCookieKey, request)
	if accessParsedClaims, err := jwt.ParseToken(accessToken, getSecret()); err == nil {
		aud := accessParsedClaims["aud"]
		sub := accessParsedClaims["sub"]
		fmt.Println("Access token aud:", aud)
		fmt.Println("Access token sub:", sub)
		log.Println("ACCESS GRANTED")
	}
	if refreshParsedClaims, err := jwt.ParseToken(refreshToken, getSecret()); err == nil {
		aud := refreshParsedClaims["aud"]
		sub := refreshParsedClaims["sub"]
		fmt.Println("Refresh token aud:", aud)
		fmt.Println("Refresh token sub:", sub)
	}
	log.Println("Update User endpoint finished")
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	log.Println("Delete User")
}
