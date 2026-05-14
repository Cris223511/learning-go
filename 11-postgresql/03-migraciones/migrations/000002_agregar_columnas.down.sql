-- Migración 2 DOWN: revierte exactamente lo que hizo el UP.
DROP INDEX IF EXISTS idx_polizas_cliente_id;

ALTER TABLE polizas
    DROP COLUMN IF EXISTS fecha_vencimiento,
    DROP COLUMN IF EXISTS notas;

ALTER TABLE clientes
    DROP COLUMN IF EXISTS telefono,
    DROP COLUMN IF EXISTS fecha_nac;
