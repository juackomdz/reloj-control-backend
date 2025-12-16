package main

import (
	"log"

	"github.com/juackomdz/control-asistencia/models"
	r "github.com/juackomdz/control-asistencia/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	models.Migrar()

	e := echo.New()
	e.Use(middleware.CORS())
	r.RouterPath(e)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"mensaje": "Todo OK"})
	})

	if err := e.Start(":3001"); err != nil {
		log.Println("Error al iniciar servidor")
	}
}
