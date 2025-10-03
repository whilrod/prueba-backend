-- Conectar a la base creada automáticamente
\connect prueba_db;

-- Crear usuario para la app
CREATE USER app_user WITH PASSWORD 'app_pass';

-- Darle permisos
GRANT ALL PRIVILEGES ON DATABASE prueba_db TO app_user;

-- Tabla de prueba
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
