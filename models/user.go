package models

type User struct {
	Id    string `json:_id`
	Email string `json:email`
	Pwd   string `json:password`
}
