package dtos

type LoginDTO struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}
