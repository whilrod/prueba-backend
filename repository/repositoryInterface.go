// interface para unificar estructura de querys
package repository

import (
	"prueba-backend/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
}
