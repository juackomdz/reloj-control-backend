package handlers

import (
	"github.com/juackomdz/control-asistencia/database"
	"github.com/labstack/echo/v4"
)

type Data struct {
	Data1 string
	Data2 string
}

type Result struct {
	HoraEntrada string `json:"hora_entrada"`
	HoraSalida  string `json:"hora_salida"`
	Horas       int    `json:"horas_trabajadas"`
	Minutos     int    `json:"minutos_trabajados"`
}

func ResumenHoras(c echo.Context) error {

	var res []Result

	id := c.Param("user_id")

	database.Conectar().Raw(`
		select
		time(hora_entrada,'localtime') as hora_entrada,
		time(hora_salida,'localtime') as hora_salida,
		cast((julianday(hora_salida) - julianday(hora_entrada)) * 24 % 24 as integer) as horas,
		cast((julianday(hora_salida) - julianday(hora_entrada)) * 24 * 60 % 60 as integer) as minutos
		from registro_asistencias
		where usuarios_id = ?
	`, id).Scan(&res)

	return c.JSON(200, res)
}
