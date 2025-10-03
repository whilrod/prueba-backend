package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func conexionDB() *gorm.DB {
	//lee variables de entorno del docker-compose.yml
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(conectionString), &gorm.Config{})
	if err != nil {
		panic("No se pudo conectar a la base de datos")
	}
	log.Println("Conexión exitosa a la base de datos")
	return db
}
