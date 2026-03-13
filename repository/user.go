package repository

import (
	"github.com/juackomdz/control-asistencia/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.Usuarios) error
	FindByEmail(email string) (*models.Usuarios, error)
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

func (r *userRepository) FindByEmail(email string) (*models.Usuarios, error) {
	var user models.Usuarios
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
