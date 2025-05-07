# PAGMELHOR

Backend da empresa pagmelhor, somos uma holding que temos um grupo de empresas.

### ENVs Local
```
# SERVER
JWT_SECRET=localhost
JWT_EXPIRESIN=3600
API_PORT=4000

# BANCO DE DADOS
DB_DRIVER=postgres
DB_HOST=localhost
DB_NAME=postgres
DB_PASSWORD=root
DB_USER=root
DB_PORT=5432

# REDIS
REDIS_URL=redis://localhost:6379

# RABBITMQ
RABBITMQ_URL=amqp://root:root@localhost:5672
RABBITMQ_DEFAULT_USER=root
RABBITMQ_DEFAULT_PASS=root
```

### Migrations

#### 1. Instalação (MacOS)
```bash
# Instalar o golang-migrate via Homebrew
brew install golang-migrate
```

#### 2. Estrutura de Migrations
```
modules/
  ├── club/
  │   └── migrations/
  │       └── YYYYMMDDHHMMSS_nome_da_migration.up.sql
  └── payments/
      └── migrations/
          └── YYYYMMDDHHMMSS_nome_da_migration.up.sql
```

#### 3. Criando uma nova Migration
```bash
# Para o módulo club
make migrate-create-club
# Digite o nome da migration quando solicitado
# Exemplo: create_table_users

# Para o módulo payments
make migrate-create-payments
# Digite o nome da migration quando solicitado
# Exemplo: create_table_transactions
```

#### 4. Escrevendo a Migration
Após criar a migration, edite o arquivo gerado com seu SQL. Exemplo:
```sql
-- modules/club/migrations/20240320123456_create_users_table.up.sql
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

#### 5. Executando as Migrations
```bash
# Executar todas as migrations pendentes
make migrate-up-club

# Reverter a última migration
make migrate-down-club
```
#### 6. Boas Práticas
- Mantenha as migrations organizadas por módulo
- Escreva SQL idempotente (usando IF NOT EXISTS, IF EXISTS)
- Teste as migrations em ambiente de desenvolvimento antes de aplicar em produção

#### 7. Variáveis de ambiente necessárias
```env
# Banco de Dados
DB_DRIVER=postgres
DB_HOST=localhost
DB_NAME=postgres
DB_PASSWORD=root
DB_USER=root
DB_PORT=5432
```