package models

import "gorm.io/gorm"

//estructura del modelo de usuario
type User struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	Nombre string `json:"nombre"`
	Email  string `json:"email" gorm:"unique"`
}

/*
TableName sobrescribe el nombre de la tabla por defecto
para uso de idioma ingles
*/
func (User) TableName() string {
	return "usuarios"
}
