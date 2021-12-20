package models

type User struct {
	ID         int    `json:"id"`
	email      string `json:"email"`
	password   string `json:"password"`
	first_name string `json:"first_name"`
	last_name  string `json:"last_name"`
	created_at string `json:"created_at"`
	updated_at string `json:"updated_at"`
}
