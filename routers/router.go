package routers

import (
	h "github.com/juackomdz/control-asistencia/handlers"
	m "github.com/juackomdz/control-asistencia/middleware"
	"github.com/labstack/echo/v4"
)

func RouterPath(e *echo.Echo) {

	api := e.Group("/api/v1")

	//Rutas acceso
	api.POST("/register", h.Registro)
	api.POST("/login", h.Ingreso)
	api.POST("/refresh", h.Refresh)

	a := api.Group("/auth", m.MiddleJWT())
	//Rutas asistencia
	a.POST("/check-in", h.CheckIn)
	a.POST("/check-out", h.CheckOut)

	//Rutas estadisticas
	a.GET("/data/:user_id", h.ResumenHoras)
}
