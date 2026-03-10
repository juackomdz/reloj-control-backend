package dtos

type RegistroDTO struct {
	Rut    string `json:"rut"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Pass   string `json:"password"`
	Role   string `json:"role"`
}

type UpdateUserDTO struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Pass   string `json:"password"`
}

type ListUsersDTO struct {
	Id             uint   `json:"id"`
	Rut            string `json:"rut"`
	NombreCompleto string `json:"nombre_completo"`
	Email          string `json:"email"`
}
