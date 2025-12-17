package handlers

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/juackomdz/control-asistencia/database"
	"github.com/juackomdz/control-asistencia/models"
	"github.com/labstack/echo/v4"
)

func Registro(c echo.Context) error {

	var registro models.RegistroDTO

	if err := c.Bind(&registro); err != nil {
		return c.JSON(500, echo.Map{"mensaje": "Error al recuperar campos"})
	}

	existe := database.Conectar().Where("email=?", registro.Email).Find(&models.Usuarios{})
	if existe.RowsAffected == 1 {
		return c.JSON(400, echo.Map{"mensaje": "Usuario ya existe"})
	} else {

		hash, err := bcrypt.GenerateFromPassword([]byte(registro.Pass), bcrypt.DefaultCost)
		if err != nil {
			log.Print(err)
		}

		database.Conectar().Create(&models.Usuarios{Rut: registro.Rut, NombreCompleto: registro.Nombre, Email: registro.Email, Password: string(hash)})

		return c.JSON(201, echo.Map{"mensaje": "Registro exitoso"})
	}

}

func Ingreso(c echo.Context) error {

	var login models.LoginDTO

	if err := c.Bind(&login); err != nil {
		return c.JSON(500, echo.Map{"mensaje": "Error al recuperar campos"})
	}

	var user models.Usuarios
	row := database.Conectar().Where("email = ?", login.Email).Find(&user)

	if row.RowsAffected == 1 {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Pass)); err != nil {
			return c.JSON(401, echo.Map{"mensaje": "Credenciales incorrectas"})
		} else {
			return c.JSON(200, echo.Map{"user_id": user.ID})
		}
	} else {
		return c.JSON(404, echo.Map{"mensaje": "Usuario no encontrado"})
	}

}
