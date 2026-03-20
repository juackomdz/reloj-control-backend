package repository

import (
	"time"

	"github.com/juackomdz/control-asistencia/models"
	"gorm.io/gorm"
)

type AssistRepository interface {
	CheckIn(*models.RegistroAsistencias) error
	CheckOut(id uint, horaSalida *time.Time) error
	FindRegistro(id uint) int64
}

type assistRepository struct {
	db *gorm.DB
}

func NewAssistRepository(db *gorm.DB) AssistRepository {
	return &assistRepository{
		db: db,
	}
}

func (r *assistRepository) CheckIn(body *models.RegistroAsistencias) error {
	return r.db.Create(&body).Error
}

func (r *assistRepository) CheckOut(id uint, horaSalida *time.Time) error {
	return r.db.Model(&models.RegistroAsistencias{}).
		Where("usuarios_id = ? AND hora_salida IS NULL", id).
		Update("hora_salida", &horaSalida).Error
}

func (r *assistRepository) FindRegistro(id uint) int64 {
	return r.db.Where(&models.RegistroAsistencias{UsuariosID: id}).Last(&models.RegistroAsistencias{}, "hora_salida is null").RowsAffected
}
