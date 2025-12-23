package routers

import (
	"log"

	h "github.com/juackomdz/control-asistencia/handlers"
	"github.com/juackomdz/control-asistencia/middleware"
	"github.com/labstack/echo/v4"
)

func RouterPath(e *echo.Echo) {

	api := e.Group("/api/v1")

	//Rutas acceso
	api.POST("/register", h.Registro)
	api.POST("/login", h.Ingreso)

	a := api.Group("/auth", middleware.MiddleJWT())
	//Rutas asistencia
	a.POST("/check-in", h.CheckIn)
	a.POST("/check-out", h.CheckOut)

	//Rutas estadisticas
	a.GET("/data/:user_id", h.ResumenHoras)

	api.GET("/test", func(c echo.Context) error {

		tok := middleware.GeneraJWT("correo@correo.cl", 1)

		log.Print(tok)
		return c.JSON(200, "ok")
	}, middleware.MiddleJWT())

}
