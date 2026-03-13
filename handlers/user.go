package handlers

import (
	"github.com/juackomdz/control-asistencia/models"
	"github.com/juackomdz/control-asistencia/models/dtos"
	"github.com/juackomdz/control-asistencia/repository"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) Save(c echo.Context) error {

	var body dtos.RegistroDTO
	if err := c.Bind(&body); err != nil {
		return c.JSON(500, echo.Map{"mensaje": "Datos invalidos"})
	}

	data := models.Usuarios{
		Rut:            body.Rut,
		NombreCompleto: body.Nombre,
		Email:          body.Email,
		Password:       body.Pass,
		Role:           body.Role,
	}
	_, err := h.repo.FindByEmail(body.Email)

	//not found lo toma como error
	if err != nil {
		if err := h.repo.Create(&data); err != nil {
			return c.JSON(500, echo.Map{"mensaje": "Error al crear usuario"})
		}
		return c.JSON(400, echo.Map{"mensaje": "Creado con exito"})
	} else {
		return c.JSON(400, echo.Map{"mensaje": "Usuario ya existe"})
	}
}

func (h *UserHandler) Test(c echo.Context) error {

	var body dtos.LoginDTO

	if err := c.Bind(&body); err != nil {
		return c.JSON(500, echo.Map{"mensaje": err.Error()})
	}
	user, err := h.repo.FindByEmail(body.Email)
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, user)
}
