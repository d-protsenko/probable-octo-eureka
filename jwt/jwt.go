package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(
	issuer string, secret []byte, issuedTime time.Time,
	tokenDuration int64, subject string, audience string,
) string {
	startTime := issuedTime.Unix()
	claims := &jwt.StandardClaims{
		ExpiresAt: startTime + tokenDuration,
		IssuedAt:  startTime,
		Issuer:    issuer,
		Subject:   subject,
		Audience:  audience,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		claims,
	)
	tokenString, err := token.SignedString(secret)
	if err == nil {
		return tokenString
	}
	return ""
}

func ParseToken(
	tokenString string, secret []byte,
) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if token == nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	} else {
		return nil, err
	}
}
