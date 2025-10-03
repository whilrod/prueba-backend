// Se implementa la lógica para interactuar con la base de datos utilizando GORM
package repository

import (
	"prueba-backend/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

/*
NewUserRepository es el constructor que recibe una instancia de *gorm.DB
para inyectar la dependencia de la DB al reposotorio
*/
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

/*
GetAll obtiene todos los registros de la tabla "usuarios"
Retorna un slice con los usuarios encontrados y un error en caso de que ocurra.
*/
func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

/*
GetByID busca un usuario por ID en la tabla "usuarios"
Retorna un puntero al usuario y un error si no existe o ocurre algún problema.
*/
func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

/*
Create inserta un nuevo usuario en la tabla "usuarios"
Retorna un puntero al usuario y un error si no existe o ocurre algún problema
*/
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(&user).Error
}

/*
Update actualiza los datos de un usuario existente
Recibe un puntero a models.User y devuelve un error si la operación falla
*/
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

/*
Delete elimina un usuario de la tabla "users" según su ID
Devuelve un error en caso de que ocurra algún problema
*/
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
