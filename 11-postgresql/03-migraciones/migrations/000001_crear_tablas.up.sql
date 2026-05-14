-- Migración 1 UP: crea las tablas base del sistema de pólizas.

CREATE TABLE IF NOT EXISTS clientes (
    id         SERIAL PRIMARY KEY,
    nombre     VARCHAR(100)  NOT NULL,
    email      VARCHAR(150)  NOT NULL UNIQUE,
    dni        VARCHAR(8)    NOT NULL UNIQUE,
    created_at TIMESTAMP     NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS polizas (
    id          SERIAL PRIMARY KEY,
    cliente_id  INTEGER       NOT NULL REFERENCES clientes(id) ON DELETE CASCADE,
    tipo        VARCHAR(50)   NOT NULL CHECK (tipo IN ('SOAT', 'Vida', 'Vehicular', 'Hogar')),
    prima       NUMERIC(10,2) NOT NULL CHECK (prima > 0),
    activa      BOOLEAN       NOT NULL DEFAULT true,
    created_at  TIMESTAMP     NOT NULL DEFAULT NOW()
);
