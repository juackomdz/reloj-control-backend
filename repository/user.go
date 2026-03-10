package repository

import (
	"github.com/juackomdz/control-asistencia/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.Usuarios) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.Usuarios) error {

	return r.db.Create(user).Error
}
