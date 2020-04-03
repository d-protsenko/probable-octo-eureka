package models

type User struct {
	id    string `json:id`
	email string `json:email`
	pwd   string `json:password`
}
