package domain

type Profile struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type ProfileService interface {
	Update(profile *Profile) error
}

type ProfileRepository interface {
	Update(profile *Profile) error
}
