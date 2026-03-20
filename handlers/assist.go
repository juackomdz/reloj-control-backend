package handlers

import (
	"github.com/juackomdz/control-asistencia/models"
	"github.com/juackomdz/control-asistencia/models/dtos"
	"github.com/juackomdz/control-asistencia/repository"
	"github.com/labstack/echo/v4"
)

type AssistHandler struct {
	repo repository.AssistRepository
}

func NewAssistHandler(repo repository.AssistRepository) *AssistHandler {
	return &AssistHandler{
		repo: repo,
	}
}

func (h *AssistHandler) CheckIn(c echo.Context) error {

	var body dtos.AsistenciaDTO

	if err := c.Bind(&body); err != nil {
		return c.JSON(500, echo.Map{"mensaje": "Datos invalidos"})
	}

	data := models.RegistroAsistencias{
		UsuariosID:  body.UsuariosID,
		HoraEntrada: body.HoraEntrada,
	}

	if rows := h.repo.FindRegistro(body.UsuariosID); rows == 0 {
		if err := h.repo.CheckIn(&data); err != nil {
			return c.JSON(500, echo.Map{"mensaje": err.Error()})
		}
		return c.JSON(201, echo.Map{"mensaje": "Registro exitoso"})
	} else {
		return c.JSON(400, echo.Map{"mensaje": "Ya registro entrada"})
	}

}

func (h *AssistHandler) CheckOut(c echo.Context) error {

	var body dtos.AsistenciaDTO

	if err := c.Bind(&body); err != nil {
		return c.JSON(500, echo.Map{"mensaje": "Datos invalidos"})
	}
	if rows := h.repo.FindRegistro(body.UsuariosID); rows == 0 {
		return c.JSON(400, echo.Map{"mensaje": "Ya registro salida"})
	} else {
		if err := h.repo.CheckOut(body.UsuariosID, body.HoraSalida); err != nil {
			return c.JSON(500, echo.Map{"mensaje": "Datos invalidos"})
		}
		return c.JSON(200, echo.Map{"mensaje": "Registro exitoso"})
	}
}

func (h *AssistHandler) Test(c echo.Context) error {

	row := h.repo.FindRegistro(3)
	return c.JSON(200, echo.Map{"mensaje": row})
}
