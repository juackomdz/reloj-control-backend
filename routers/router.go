package routers

import (
	h "github.com/juackomdz/control-asistencia/handlers"
	"github.com/juackomdz/control-asistencia/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouterPath(e *echo.Echo, db *gorm.DB) {

	userRepo := repository.NewUserRepository(db)
	userHandler := h.NewUserHandler(userRepo)
	assistRepo := repository.NewAssistRepository(db)
	assistHandler := h.NewAssistHandler(assistRepo)

	api := e.Group("/api/v1")

	//test
	api.POST("/user", userHandler.Save)
	api.POST("/test", assistHandler.Test)

	api.POST("/check-in", assistHandler.CheckIn)
	api.POST("/check-out", assistHandler.CheckOut)
	//Rutas acceso
	//api.POST("/login", h.Ingreso)
	//api.POST("/refresh", h.Refresh)

	//Rutas admin
	//api.PATCH("/users/:user_id", h.Modificar, m.MiddleJWT(), m.RequireRole(m.RoleAdmin))
	//api.GET("/users", h.Listar, m.MiddleJWT(), m.RequireRole(m.RoleAdmin))
	//api.POST("/register", h.Registro, m.MiddleJWT(), m.RequireRole(m.RoleAdmin))

	//Rutas user
	//api.POST("/check-in", h.CheckIn, m.MiddleJWT(), m.RequireRole(m.RoleUser))
	//api.POST("/check-out", h.CheckOut, m.MiddleJWT(), m.RequireRole(m.RoleUser))

	//Rutas estadisticas
	//api.GET("/data/:user_id", h.ResumenHoras)
}
