package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConexionDB() *gorm.DB {
	//lee variables de entorno del docker-compose.yml
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(conectionString), &gorm.Config{})
		if err != nil {
			break
		}
		log.Println("Esperando a la DB, reintentando en 3s...")
		time.Sleep(3 * time.Second)
		log.Println("Conexión exitosa a la base de datos")
	}
	if err != nil {
		panic("No se pudo conectar a la base de datos: " + err.Error())
	}

	log.Println("✅ Conexión exitosa a la base de datos")
	return DB
}
