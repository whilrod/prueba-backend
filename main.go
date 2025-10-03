package main

import (
	"log"
	"net/http"
	"prueba-backend/config"
	"prueba-backend/controllers"
	"prueba-backend/models"
	"prueba-backend/repository"
	"prueba-backend/routes"
)

func main() {
	config.ConexionDB()

	// Migración automática de gorm.model
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("❌ Error en migración:", err)
	}

	userRepo := repository.NewUserRepository(config.DB)
	userController := controllers.NewUserController(userRepo)

	mux := http.NewServeMux()
	routes.RegisterUserRoutes(mux, userController)

	// Servidor
	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
