package handlers

import (
	"strconv"

	db "github.com/juackomdz/control-asistencia/database"
	"github.com/labstack/echo/v4"
)

type Result struct {
	HoraEntrada string `json:"hora_entrada"`
	HoraSalida  string `json:"hora_salida"`
	Horas       int    `json:"horas_trabajadas"`
	Minutos     int    `json:"minutos_trabajados"`
}

func ResumenHoras(c echo.Context) error {

	var res []Result

	id := c.Param("user_id")

	if id == "" {
		return c.JSON(400, echo.Map{"error": "user_id es requerido"})
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, echo.Map{"error": "user_id debe ser un número válido"})
	}

	query := db.Conectar().Table("registro_asistencias").
		Select([]string{
			"time(hora_entrada,'localtime') as hora_entrada",
			"time(hora_salida,'localtime') as hora_salida",
			"cast((julianday(hora_salida) - julianday(hora_entrada)) * 24 % 24 as integer) as horas",
			"cast((julianday(hora_salida) - julianday(hora_entrada)) * 24 * 60 % 60 as integer) as minutos",
		}).
		Where("usuarios_id = ?", userID)

	if err := query.Scan(&res).Error; err != nil {
		return c.JSON(500, echo.Map{"error": "Error al obtener el resumen de horas"})
	}

	return c.JSON(200, res)
}
