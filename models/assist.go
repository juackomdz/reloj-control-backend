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
