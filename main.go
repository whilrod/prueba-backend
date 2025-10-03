package main

import (
	"log"
	"net/http"
	"prueba-backend/config"
	"prueba-backend/controllers"
	"prueba-backend/repository"
	"prueba-backend/routes"
)

func main() {
	config.ConexionDB()

	userRepo := repository.NewUserRepository(config.DB)
	userController := controllers.NewUserController(userRepo)

	mux := http.NewServeMux()
	routes.RegisterUserRoutes(mux, userController)

	// Servidor
	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
