package database

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/juackomdz/control-asistencia/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Dbase *gorm.DB

func Conectar() *gorm.DB {

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true,
		},
	)

	sqliteDb, err := gorm.Open(sqlite.Open(os.Getenv("BBDD")), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		log.Println("No se pudo conectar a la bd: ", err)
	}

	Dbase = sqliteDb

	return Dbase
}

func Migrar() {
	Conectar().AutoMigrate(&models.Usuarios{}, &models.RegistroAsistencias{})
}
