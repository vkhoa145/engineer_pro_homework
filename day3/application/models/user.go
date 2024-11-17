package models

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	UserProfile string `json:"user_profile"`
}
