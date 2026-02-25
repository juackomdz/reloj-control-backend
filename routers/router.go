package routers

import (
	h "github.com/juackomdz/control-asistencia/handlers"
	m "github.com/juackomdz/control-asistencia/middleware"
	"github.com/labstack/echo/v4"
)

func RouterPath(e *echo.Echo) {

	api := e.Group("/api/v1")

	//Rutas acceso
	api.POST("/login", h.Ingreso)
	api.POST("/refresh", h.Refresh)

	//Rutas admin
	api.PATCH("/users/:user_id", h.Modificar, m.MiddleJWT(), m.RequireRole(m.RoleAdmin))
	api.GET("/users", h.Listar, m.MiddleJWT(), m.RequireRole(m.RoleAdmin))
	api.POST("/register", h.Registro, m.MiddleJWT(), m.RequireRole(m.RoleAdmin))

	//Rutas user
	api.POST("/check-in", h.CheckIn, m.MiddleJWT(), m.RequireRole(m.RoleUser))
	api.POST("/check-out", h.CheckOut, m.MiddleJWT(), m.RequireRole(m.RoleUser))

	//Rutas estadisticas
	api.GET("/data/:user_id", h.ResumenHoras)
}
