package domain

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	RefreshToken string `json:"refresh_token"`
}

type UserService interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
	FindByID(id int) (*User, error)
}

type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
	FindByID(id int) (*User, error)
}
