package models

type User struct {
    User_ID        int    `json:"user_id"`
    User_Name      string `json:"user_name"`
	User_Last_name string `json:"user_last_name"`
    User_Email     string `json:"user_email"`
    CreatedAt string `json:"created_at"`
}