package cookie

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"net/http"
)

var blockKey []byte
var hashKey []byte
var secureCookie *securecookie.SecureCookie
var path string
var domain string

func InitializeSecureCookie(hKey string, bKey string, cookiePath string, cookieDomain string) {
	blockKey = []byte(bKey)
	hashKey = []byte(hKey)
	path = cookiePath
	domain = cookieDomain
	secureCookie = securecookie.New(hashKey, blockKey)
}

func SetCookie(cookieName string, key string, value string, maxAge int, w http.ResponseWriter) {
	cookieValue := map[string]string{
		key: value,
	}
	if encoded, err := secureCookie.Encode(cookieName, cookieValue); err == nil {
		cookie := &http.Cookie{
			Name:   cookieName,
			Value:  encoded,
			Path:   path,
			MaxAge: maxAge,
			//HttpOnly: true,
			//Secure:   true,
		}
		http.SetCookie(w, cookie)
	} else {
		fmt.Println(err)
	}
}

func ReadCookie(cookieName string, key string, r *http.Request) string {
	if cookie, err := r.Cookie(cookieName); err == nil {
		value := make(map[string]string)
		if err = secureCookie.Decode(cookieName, cookie.Value, &value); err == nil {
			return value[key]
		}
	}
	return ""
}
