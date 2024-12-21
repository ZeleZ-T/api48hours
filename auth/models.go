package auth

type User struct {
	email    string `json:"email"`
	password string `json:"password"`
}
