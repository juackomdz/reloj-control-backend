package models

import "time"

type AsistenciaDTO struct {
	UsuariosID  uint       `json:"usuario_id"`
	HoraEntrada *time.Time `json:"hora_entrada"`
	HoraSalida  *time.Time `json:"hora_salida"`
}

type RegistroDTO struct {
	Rut    string `json:"rut"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Pass   string `json:"password"`
	Role   string `json:"role"`
}

type LoginDTO struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}
