-- Tabla de clientes. Solo clientes-svc la gestiona.
-- En microservicios reales cada servicio tendría su propia base de datos,
-- acá comparten la misma instancia de PostgreSQL pero con tablas separadas.

CREATE TABLE IF NOT EXISTS clientes (
    id         VARCHAR(20)  PRIMARY KEY,
    nombre     VARCHAR(100) NOT NULL,
    email      VARCHAR(150) NOT NULL UNIQUE,
    dni        VARCHAR(8)   NOT NULL UNIQUE,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW()
);
