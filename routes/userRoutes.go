package routes

import (
	"net/http"
	"prueba-backend/controllers"
)

// RegisterUserRoutes registra las rutas relacionadas con usuarios
func RegisterUserRoutes(mux *http.ServeMux, userController *controllers.UserController) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userController.GetAllUsers(w, r)
		case http.MethodPost:
			userController.CreateUser(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		// Aquí se asume que /users/{id}
		switch r.Method {
		case http.MethodGet:
			userController.GetUserByID(w, r)
		case http.MethodPut:
			userController.UpdateUser(w, r)
		case http.MethodDelete:
			userController.DeleteUser(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
}
