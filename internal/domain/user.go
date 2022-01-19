package domain

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	RefreshToken string `json:"refresh_token"`
}
