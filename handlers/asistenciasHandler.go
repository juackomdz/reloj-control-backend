package handlers

import (
	"time"

	db "github.com/juackomdz/control-asistencia/database"
	"github.com/juackomdz/control-asistencia/models"
	"github.com/labstack/echo/v4"
)

func CheckIn(c echo.Context) error {

	ahora := time.Now()

	id := c.Get("user").(uint)

	existe := db.Conectar().Where(&models.RegistroAsistencias{UsuariosID: id}).Last(&models.RegistroAsistencias{}, "hora_salida is null")

	if existe.RowsAffected == 0 {
		datos := models.RegistroAsistencias{UsuariosID: id, HoraEntrada: &ahora}
		db.Conectar().Create(&datos)
		return c.JSON(200, echo.Map{"mensaje": "Registro Exitoso"})
	} else {
		return c.JSON(400, echo.Map{"mensaje": "Ya registro entrada"})
	}

}

func CheckOut(c echo.Context) error {

	ahora := time.Now()
	id := c.Get("user").(uint)

	exito := db.Conectar().Model(&models.RegistroAsistencias{}).Where(&models.RegistroAsistencias{UsuariosID: id}).Where("hora_salida is null").Update("hora_salida", &ahora)

	if exito.RowsAffected == 0 {
		return c.JSON(400, echo.Map{"mensaje": "Ya registro salida"})
	}

	return c.JSON(200, echo.Map{"mensaje": "Registro Exitoso"})
}
