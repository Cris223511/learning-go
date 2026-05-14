-- Tabla de pólizas. Cada microservicio tiene su propia tabla (o BD en producción real).
-- polizas-svc es el único dueño de esta tabla: ningún otro servicio la toca directamente.

CREATE TABLE IF NOT EXISTS polizas (
    id         SERIAL PRIMARY KEY,
    cliente_id VARCHAR(50)   NOT NULL,
    tipo       VARCHAR(50)   NOT NULL CHECK (tipo IN ('SOAT', 'Vida', 'Vehicular', 'Hogar')),
    prima      NUMERIC(10,2) NOT NULL CHECK (prima > 0),
    activa     BOOLEAN       NOT NULL DEFAULT true,
    created_at TIMESTAMP     NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_polizas_cliente ON polizas(cliente_id);
