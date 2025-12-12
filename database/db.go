package database

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Conectar() *gorm.DB {

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			IgnoreRecordNotFoundError: true,
		},
	)

	db, err := gorm.Open(sqlite.Open("test1.db"), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		log.Println("No se pudo conectar a la bd: ", err)
	}
	return db
}
