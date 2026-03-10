package dtos

import "time"

type AsistenciaDTO struct {
	UsuariosID  uint       `json:"usuario_id"`
	HoraEntrada *time.Time `json:"hora_entrada"`
	HoraSalida  *time.Time `json:"hora_salida"`
}
