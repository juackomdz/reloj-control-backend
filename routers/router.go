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
}
