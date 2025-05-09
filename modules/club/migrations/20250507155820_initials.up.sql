-- create o schema club
CREATE SCHEMA IF NOT EXISTS club;

-- set o schema club como padrão
SET search_path TO club;

-- create schema_migrations table in club schema
CREATE TABLE IF NOT EXISTS club.schema_migrations (
    version bigint NOT NULL PRIMARY KEY,
    dirty boolean NOT NULL
);

-- create table contractors
CREATE TABLE IF NOT EXISTS contractors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,            -- Nome do contractor (Universo feminino, corpo e saúde, etc)
    cpf_cnpj VARCHAR(14) UNIQUE NOT NULL,  -- CPF ou CNPJ (formato com máscara)
    email VARCHAR(255) NOT NULL UNIQUE,
    slug VARCHAR(255) NOT NULL UNIQUE,     -- slug do contractor (o-universo-feminino)
    phone VARCHAR(255) NULL,
    password VARCHAR(255) NULL,
    image VARCHAR(255) NULL,              -- Imagem do contractor (opcional)
    address VARCHAR(255) NULL,            -- Endereço do contractor (opcional)
    city VARCHAR(100) NULL,               -- Cidade (opcional)
    state VARCHAR(50) NULL,               -- Estado (opcional)
    zip_code VARCHAR(10) NULL,            -- CEP (opcional)
    description TEXT NULL,                -- Descrição do contractor (opcional)
    meta_title VARCHAR(255) NULL,         -- Meta title do contractor (opcional)
    meta_description TEXT NULL,           -- Meta description do contractor (opcional)
    meta_keywords TEXT NULL,              -- Meta keywords do contractor (opcional)    
    created_at timestamptz NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

-- create indexes for contractors
CREATE INDEX IF NOT EXISTS idx_contractors_deleted_at ON contractors(deleted_at);
CREATE INDEX IF NOT EXISTS idx_contractors_slug ON contractors(slug);
CREATE INDEX IF NOT EXISTS idx_contractors_email ON contractors(email);
CREATE INDEX IF NOT EXISTS idx_contractors_cpf_cnpj ON contractors(cpf_cnpj);

-- create table partners
CREATE TABLE IF NOT EXISTS partners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,            -- Nome do parceiro
    cpf_cnpj VARCHAR(14) UNIQUE NOT NULL,  -- CPF ou CNPJ do parceiro
    email VARCHAR(255) NOT NULL UNIQUE,
    slug VARCHAR(255) NOT NULL UNIQUE,     -- slug do parceiro (clinica-da-mama, sabin, exame, clinica da mulher)
    status smallint NOT NULL DEFAULT 1,    -- Status do parceiro (1 - Ativo, 0 - Inativo)
    phone VARCHAR(255) NULL,
    password VARCHAR(255) NULL,            -- Senha do parceiro (opcional)
    image VARCHAR(255) NULL,              -- Imagem do parceiro (opcional)
    address VARCHAR(255) NULL,            -- Endereço do parceiro (opcional)
    city VARCHAR(100) NULL,               -- Cidade (opcional)
    state VARCHAR(50) NULL,               -- Estado (opcional)
    zip_code VARCHAR(10) NULL,            -- CEP (opcional)
    description TEXT NULL,                -- Descrição do parceiro (opcional)
    meta_title VARCHAR(255) NULL,         -- Meta title do parceiro (opcional)
    meta_description TEXT NULL,           -- Meta description do parceiro (opcional)
    meta_keywords TEXT NULL,              -- Meta keywords do parceiro (opcional)
    created_at timestamptz NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

-- create indexes for partners
CREATE INDEX IF NOT EXISTS idx_partners_deleted_at ON partners(deleted_at);
CREATE INDEX IF NOT EXISTS idx_partners_slug ON partners(slug);
CREATE INDEX IF NOT EXISTS idx_partners_email ON partners(email);
CREATE INDEX IF NOT EXISTS idx_partners_cpf_cnpj ON partners(cpf_cnpj);

-- create table users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(20),
    cpf_cnpj VARCHAR(14),                     -- CPF ou CNPJ
    password VARCHAR(255) NULL,
    image VARCHAR(255) NULL,
    address VARCHAR(255) NULL,
    city VARCHAR(100) NULL,
    state VARCHAR(50) NULL,
    zip_code VARCHAR(10) NULL,
    
    created_at timestamptz NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);
-- create indexes for users
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);
CREATE INDEX IF NOT EXISTS idx_users_cpf_cnpj ON users(cpf_cnpj);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);


-- benefits
CREATE TABLE IF NOT EXISTS benefits (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,                        -- Nome do benefício
    description TEXT NULL,                             -- Descrição do benefício (opcional)
    original_price DECIMAL(10,2) NULL,                 -- Preço original do benefício (opcional)
    discount_type VARCHAR(20) CHECK (discount_type IN ('percent', 'fixed')), -- Tipo de desconto
    discount_value DECIMAL(10,2) NULL,                 -- Valor do desconto (opcional)
    status smallint NOT NULL DEFAULT 1,                -- Status do benefício (1 - Ativo, 0 - Inativo)
    notes TEXT NULL,                                   -- Notas adicionais (opcional)

    contractor_id INTEGER REFERENCES contractors(id),  -- Se for um benefício exclusivo do contractor
    partner_id INTEGER REFERENCES partners(id),        -- Se for um benefício oferecido por um parceiro

    created_at timestamptz NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

-- create indexes for benefits
CREATE INDEX IF NOT EXISTS idx_benefits_deleted_at ON benefits(deleted_at);
CREATE INDEX IF NOT EXISTS idx_benefits_contractor_id ON benefits(contractor_id);
CREATE INDEX IF NOT EXISTS idx_benefits_partner_id ON benefits(partner_id);

-- create table benefit items
CREATE TABLE IF NOT EXISTS benefit_items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,                      -- Nome do item (ex: "Exame laboratorial")
    coverage_percent DECIMAL(5,2),                   -- Percentual de cobertura (ex: 80.00 para 80%)
    limit_per_day DECIMAL(10,2),                     -- Limite de uso por dia (ex: 2 sessões)
    limit_total DECIMAL(10,2),                       -- Limite total de usos por item (opcional)
    max_coverage_value DECIMAL(10,2),                -- Valor máximo de cobertura por uso (R$)
    unlimited BOOLEAN DEFAULT FALSE,                 -- Se o item não tem limite
    notes TEXT,                                      -- Observações adicionais
    status smallint NOT NULL DEFAULT 1,              -- Status do item (1 - Ativo, 0 - Inativo)
    original_price DECIMAL(10,2),                    -- Preço original do item (opcional)
    discount_type VARCHAR(20) CHECK (discount_type IN ('percent', 'fixed')), -- Tipo de desconto
    discount_value DECIMAL(10,2),                    -- Valor do desconto (opcional)

    benefit_id INTEGER NOT NULL REFERENCES benefits(id),
    
    created_at timestamptz NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

-- create indexes for benefit items
CREATE INDEX IF NOT EXISTS idx_benefit_items_deleted_at ON benefit_items(deleted_at);
CREATE INDEX IF NOT EXISTS idx_benefit_items_benefit_id ON benefit_items(benefit_id);

-- create table users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(20),
    cpf_cnpj VARCHAR(14),                     -- CPF ou CNPJ
    password VARCHAR(255) NULL,
    image VARCHAR(255) NULL,
    address VARCHAR(255) NULL,
    city VARCHAR(100) NULL,
    state VARCHAR(50) NULL,
    zip_code VARCHAR(10) NULL,
    
    created_at timestamptz NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

-- create indexes for users
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);
CREATE INDEX IF NOT EXISTS idx_users_cpf_cnpj ON users(cpf_cnpj);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- create table benefit item uses
CREATE TABLE users_benefit_usage (
    id SERIAL PRIMARY KEY,
    voucher_code VARCHAR(255) NOT NULL,                                       -- Código do uso do benefício
    description TEXT,                                                         -- Descrição do uso do benefício
    discount_type VARCHAR(20) CHECK (discount_type IN ('percent', 'fixed')),  -- Tipo de desconto
    discount_value DECIMAL(10,2),                                             -- Valor do desconto (opcional)
    notes TEXT,                                                              -- Notas adicionais (opcional)
    status VARCHAR(10) CHECK (status IN ('used', 'reserved', 'cancelled')),   -- Status do uso do benefício

    user_id INTEGER NOT NULL REFERENCES users(id),
    partner_id INTEGER NOT NULL REFERENCES partners(id),
    benefit_id INTEGER NOT NULL REFERENCES benefits(id),
    benefit_item_id INTEGER REFERENCES benefit_items(id),       -- Opcional: qual item do benefício foi usado

    created_at timestamptz NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
    updated_at timestamptz NULL,
    deleted_at timestamptz NULL
);

-- create indexes for users benefit usage
CREATE INDEX IF NOT EXISTS idx_users_benefit_usage_deleted_at ON users_benefit_usage(deleted_at);
CREATE INDEX IF NOT EXISTS idx_users_benefit_usage_user_id ON users_benefit_usage(user_id);
CREATE INDEX IF NOT EXISTS idx_users_benefit_usage_partner_id ON users_benefit_usage(partner_id);
CREATE INDEX IF NOT EXISTS idx_users_benefit_usage_benefit_id ON users_benefit_usage(benefit_id);
CREATE INDEX IF NOT EXISTS idx_users_benefit_usage_benefit_item_id ON users_benefit_usage(benefit_item_id);

