package handlers

import (
	"log"

	"github.com/juackomdz/control-asistencia/models"
	"github.com/juackomdz/control-asistencia/models/dtos"
	"github.com/juackomdz/control-asistencia/repository"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

	hash, errHash := bcrypt.GenerateFromPassword([]byte(body.Pass), bcrypt.DefaultCost)
	if errHash != nil {
		log.Print(errHash)
	}
	data := models.Usuarios{
		Rut:            body.Rut,
		NombreCompleto: body.Nombre,
		Email:          body.Email,
		Password:       string(hash),
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

// TODO eliminar metodo test
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
