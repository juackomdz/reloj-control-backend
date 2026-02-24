package main

import (
	"log"

	"github.com/juackomdz/control-asistencia/database"
	r "github.com/juackomdz/control-asistencia/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database.Conectar()
	database.Migrar()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:4173"},
	}))
	r.RouterPath(e)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, echo.Map{"mensaje": "Todo OK"})
	})

	if err := e.Start(":3001"); err != nil {
		log.Println("Error al iniciar servidor")
	}
}
