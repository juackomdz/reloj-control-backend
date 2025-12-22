package routers

import (
	h "github.com/juackomdz/control-asistencia/handlers"
	"github.com/labstack/echo/v4"
)

func RouterPath(e *echo.Echo) {

	api := e.Group("/api/v1")

	//Rutas asistencia
	api.POST("/check-in", h.CheckIn)
	api.POST("/check-out", h.CheckOut)

	//Rutas acceso
	api.POST("/register", h.Registro)
	api.POST("/login", h.Ingreso)

	//Rutas estadisticas
	api.GET("/data/:user_id", h.ResumenHoras)

}
