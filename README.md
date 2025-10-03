# Prueba Desarrollador Backend Go + PostgreSQL + Frontend + Docker

## 🔹 Descripción
Proyecto CRUD completo de **usuarios** utilizando **Go**, **GORM**, **PostgreSQL** y un **frontend básico con HTML y CSS**.  
Permite **crear, leer, actualizar y eliminar usuarios** mediante un **API REST** y un frontend visual.

---

## 🔹 Tecnologías usadas
- Go 1.21
- PostgreSQL 15
- GORM (ORM de Go)
- HTML + CSS
- Docker y Docker Compose

---

## 🔹 Requisitos
- Docker y Docker Compose instalados
- Go 1.21 (opcional si quieres ejecutar fuera de Docker)

---

## 🔹 Estructura del proyecto
```bash
    prueba-backend/
│
├── main.go                 # Punto de entrada del proyecto
├── go.mod                  # Módulo de Go
├── go.sum
│
├── config/                 # Configuración de la DB y variables de entorno
│   └── config.go
│
├── controllers/            # Lógica de controladores HTTP
│   └── userController.go
│
├── repository/             # Implementación de repositorios (interfaz + implementación)
│   ├── userRepository.go
│   └── implementationInterface.go
│
├── models/                 # Definición de structs GORM
│   └── user.go
│
├── routes/                 # Registro de rutas HTTP
│   └── routes.go
│
├── frontend/               # Archivos estáticos para el frontend
│   ├── index.html
│   └── styles.css
│
├── docker-compose.yml      # Contenedores (Postgres + Go API)
├── Dockerfile              # Imagen del backend Go
├── init/                   # Scripts de inicialización de la DB
│   └── init.sql
└── README.md               # Documentación del proyecto
```

## 🔹 Instalación y ejecución
1. Clonar el repositorio:
```bash
git clone https://github.com/whilrod/prueba-backend.git
cd prueba-backend
```
---

2. Levantar los contenedores:
```bash
    docker-compose up --build
```
Esto iniciará:

- Contenedor postgres_db con la base de datos inicializada.
- Contenedor go_api corriendo en http://localhost:8080

3. Acceder al frontend en tu navegador:

- http://localhost:8080

## 🔹 Documentación de uso del API REST
#### ENDPOINTS
| Método | Ruta | Descripción | Body/Params |
| :--- | :---: | ---: | ---: |
| GET | /users | Lista a todos los usuarios | — |
| GET | /users/{id} | Obtiene un usuario por ID | id en URL |
| POST | /users | Crea un usuario | JSON: { "nombre": "Juan", "email": "juan@example.com" } |
| PUT | /users/{id} | Actualiza un usuario | JSON: { "nombre": "Nuevo", "email": "nuevo@example.com" } |
| DELETE | /users/{id} | Elimina un usuario | id en URL |

#### Ejemplos de uso

- Consultar la lista de usuarios (GET)
![alt text](image.png)
- Consultar un usuario por ID (GET)
![alt text](image-1.png)
- Crear un usuario (POST)
![alt text](image-5.png)
- Actualizar un usuario (PUT)
![alt text](image-4.png)
![alt text](image-2.png)
- Eliminar un usuario (DELETE)
![alt text](image-3.png)

## 🔹 Script de creación de base de datos (init.sql)

```sql
-- Crear usuario para la app
CREATE USER app_user WITH PASSWORD 'app_pass';

-- Darle permisos
GRANT ALL PRIVILEGES ON DATABASE prueba_db TO app_user;

-- Tabla usuarios
CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insertar fila inicial
INSERT INTO usuarios (nombre, email) VALUES
('Juan Pérez', 'juan@example.com'),
('Ana Gómez', 'ana@example.com'),
('Carlos Ruiz', 'carlos@example.com');
```
* Nota: Este script se ejecuta automáticamente

## 🔹 Frontend

- Archivos HTML y CSS servidos desde Go.
- Permite interactuar con todos los endpoints del CRUD.
- Formulario de creación/edición, lista de usuarios y botones de guardar, editar y eliminar.
- Estilos simples pero agradables con CSS puro.
![alt text](image-6.png)