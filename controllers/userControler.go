/*
Maneja la logica de negocio
*/
package controllers

import (
	"encoding/json"
	"net/http"
	"prueba-backend/models"
	"prueba-backend/repository"
	"strconv"
	"strings"
)

type UserController struct {
	Repo repository.UserRepository
}

/*
NewUserController Constructor: inicializa y devuelve una instancia de la estructura .
Devuelve un puntero a UserController usando Inyección de Dependencias
*/
func NewUserController(repo repository.UserRepository) *UserController {
	return &UserController{Repo: repo}
}

// GET /users
func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Repo.GetAll()
	if err != nil {
		http.Error(w, "Error obteniendo usuarios", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET /users/{id}
func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	user, err := c.Repo.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// POST /users
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := c.Repo.Create(&user); err != nil {
		http.Error(w, "Error creando usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// PUT /users/{id}
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	user.ID = uint(id)

	if err := c.Repo.Update(&user); err != nil {
		http.Error(w, "Error actualizando usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DELETE /users/{id}
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := c.Repo.Delete(uint(id)); err != nil {
		http.Error(w, "Error eliminando usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
