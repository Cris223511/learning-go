-- Migración 2 UP: agrega campos adicionales a las tablas existentes.

ALTER TABLE clientes
    ADD COLUMN IF NOT EXISTS telefono  VARCHAR(15),
    ADD COLUMN IF NOT EXISTS fecha_nac DATE;

ALTER TABLE polizas
    ADD COLUMN IF NOT EXISTS fecha_vencimiento DATE,
    ADD COLUMN IF NOT EXISTS notas             TEXT;

-- Índice para acelerar búsquedas de pólizas por cliente.
CREATE INDEX IF NOT EXISTS idx_polizas_cliente_id ON polizas(cliente_id);
