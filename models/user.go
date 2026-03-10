package models

import "gorm.io/gorm"

type Usuarios struct {
	gorm.Model
	Rut            string `gorm:"type:varchar(12)"`
	NombreCompleto string `gorm:"type:varchar(100)"`
	Email          string `gorm:"unique,type:varchar(50)"`
	Password       string `gorm:"type:varchar(100)"`
	Role           string `gorm:"type:varchar(15)"`
	Registros      []RegistroAsistencias
}
