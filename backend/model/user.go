package model

type User struct {
	Username string `json:"username" gorm:"primary_key"`
	Password string `json:"password"`
}

type UserRequest struct {
	Username string `json:"username"`
}
