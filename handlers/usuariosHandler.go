package handlers

import (
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	db "github.com/juackomdz/control-asistencia/database"
	"github.com/juackomdz/control-asistencia/middleware"
	"github.com/juackomdz/control-asistencia/models"
	"github.com/labstack/echo/v4"
)

func Registro(c echo.Context) error {

	var register models.RegistroDTO

	if err := c.Bind(&register); err != nil {
		return c.JSON(500, echo.Map{"mensaje": "Error al recuperar campos"})
	}

	exist := db.Dbase.Where("email=?", register.Email).Find(&models.Usuarios{})
	if exist.RowsAffected == 1 {
		return c.JSON(400, echo.Map{"mensaje": "Usuario ya existe"})
	} else {

		hash, err := bcrypt.GenerateFromPassword([]byte(register.Pass), bcrypt.DefaultCost)
		if err != nil {
			log.Print(err)
		}

		db.Conectar().Create(&models.Usuarios{Rut: register.Rut, NombreCompleto: register.Nombre, Email: register.Email, Password: string(hash), Role: register.Role})

		return c.JSON(201, echo.Map{"mensaje": "Registro exitoso"})
	}

}

func Ingreso(c echo.Context) error {

	var login models.LoginDTO

	if err := c.Bind(&login); err != nil {
		return c.JSON(500, echo.Map{"mensaje": "Error al recuperar campos"})
	}

	var user models.Usuarios
	row := db.Dbase.Where("email = ?", login.Email).Find(&user)

	if row.RowsAffected == 1 {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Pass)); err != nil {
			return c.JSON(401, echo.Map{"mensaje": "Credenciales incorrectas"})
		} else {
			tokenA, tokenR := middleware.GenerateJWT(user.Email, user.ID, user.Role)
			return c.JSON(200, echo.Map{
				"auth_token":    tokenA,
				"refresh_token": tokenR,
			})
		}
	} else {
		return c.JSON(404, echo.Map{"mensaje": "Usuario no encontrado"})
	}

}

func Listar(c echo.Context) error {

	var dto = []models.ListUsersDTO{}

	db.Dbase.Table("usuarios").Select([]string{"id", "rut", "nombre_completo", "email"}).Find(&dto)

	return c.JSON(200, dto)
}

func Modificar(c echo.Context) error {

	id := c.Param("user_id")

	if id == "" {
		return c.JSON(400, echo.Map{"error": "user_id es requerido"})
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, echo.Map{"error": "user_id debe ser un número válido"})
	}

	exist := db.Dbase.Where("id = ?", userID).Find(&models.Usuarios{})

	if exist.RowsAffected == 1 {

		var body models.UpdateUserDTO

		if err := c.Bind(&body); err != nil {
			return c.JSON(500, echo.Map{"mensaje": "Error al recuperar campos"})
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(body.Pass), bcrypt.DefaultCost)
		if err != nil {
			log.Print(err)
		}

		db.Dbase.Where("id = ?", userID).Updates(&models.Usuarios{NombreCompleto: body.Nombre, Email: body.Email, Password: string(hash)})

		return c.JSON(200, echo.Map{"mensaje": "Modificado con exito"})
	}

	return c.JSON(500, echo.Map{"mensaje": "Error al actualizar usuario"})
}

func Refresh(c echo.Context) error {

	h := c.Request().Header.Get("Refresh")

	if len(h) == 0 {
		return c.JSON(400, echo.Map{"mensaje": "Falta token"})
	}

	newToken, err := middleware.RefreshToken(h)
	if err != nil {
		log.Print(err)
	}

	return c.JSON(200, echo.Map{"token": newToken})
}
