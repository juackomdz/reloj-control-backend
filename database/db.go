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

func init() {

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true,
		},
	)

	sqliteDb, err := gorm.Open(sqlite.Open("test1.db"), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		log.Println("No se pudo conectar a la bd: ", err)
	}

	Dbase = sqliteDb
}

func Conectar() *gorm.DB {
	return Dbase
}

func Migrar() {
	Conectar().AutoMigrate(&models.Usuarios{}, &models.RegistroAsistencias{})
}
