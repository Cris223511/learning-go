-- Migración 1 DOWN: elimina las tablas en orden inverso (primero la que tiene FK).
DROP TABLE IF EXISTS polizas;
DROP TABLE IF EXISTS clientes;
