package domain

type Room struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RoomService interface {
	Create(room *Room) error
	Update(room *Room) error
	Delete(id int) error
	FindByID(id int) (*Room, error)
}

type RoomRepository interface {
	Create(room *Room) error
	Update(room *Room) error
	Delete(id int) error
	FindByID(id int) (*Room, error)
}
