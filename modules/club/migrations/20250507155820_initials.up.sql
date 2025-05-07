-- create o schema club
CREATE SCHEMA IF NOT EXISTS club;

-- set o schema club como padrão
SET search_path TO club;

-- create table organizations
CREATE TABLE IF NOT EXISTS organizations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    slug VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

-- create index for deleted_at
CREATE INDEX IF NOT EXISTS idx_organizations_deleted_at ON organizations(deleted_at);

-- create table users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    organization_id INT NOT NULL,
    phone VARCHAR(255) NULL,
    created_at timestamptz NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

-- indexes to users
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);
CREATE INDEX IF NOT EXISTS idx_users_organization_id ON users(organization_id);

