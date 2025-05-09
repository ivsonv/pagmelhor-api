-- create o schema club
CREATE SCHEMA IF NOT EXISTS club;

-- set o schema club como padrão
SET search_path TO club;

-- create schema_migrations table in club schema
CREATE TABLE IF NOT EXISTS club.schema_migrations (
    version bigint NOT NULL PRIMARY KEY,
    dirty boolean NOT NULL
);