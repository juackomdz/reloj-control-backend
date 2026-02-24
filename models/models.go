package models

import (
	"time"

	"gorm.io/gorm"
)

type RegistroAsistencias struct {
	gorm.Model
	UsuariosID  uint
	HoraEntrada *time.Time
	HoraSalida  *time.Time
}

type Usuarios struct {
	gorm.Model
	Rut            string `gorm:"type:varchar(12)"`
	NombreCompleto string `gorm:"type:varchar(100)"`
	Email          string `gorm:"unique,type:varchar(50)"`
	Password       string `gorm:"type:varchar(100)"`
	Role           string `gorm:"type:varchar(15)"`
	Registros      []RegistroAsistencias
}
