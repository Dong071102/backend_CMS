package models

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
