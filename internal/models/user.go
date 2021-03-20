package models

type User struct {
	*Model
	Username string `json:"username"`
	Password string `json:"-"`
}
