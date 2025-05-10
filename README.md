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



### Erros Comuns

#### Ubuntu

Se encontrar o erro "address already in use" (endereço já em uso):

1. Verificar processos usando a porta
```sh
sudo lsof -i -P -n
```

2. Encerrar o processo
```sh
sudo kill -9 PID
```

Onde PID é o número do processo identificado no comando anterior.

### Notas por Sistema Operacional

#### Windows

1. **WSL (Windows Subsystem for Linux)**
   - Recomendamos fortemente usar o WSL2 para desenvolvimento
   - Instale o Docker Desktop com suporte WSL2
   - Execute todos os comandos dentro do ambiente WSL

2. **Permissões**
   - Execute o Docker Desktop como administrador
   - Certifique-se que o WSL2 está configurado corretamente:
   ```powershell
   wsl --set-default-version 2
   ```

#### macOS

1. **Docker Desktop**
   - Certifique-se de alocar memória suficiente para o Docker Desktop
   - Recomendado: 8GB RAM, 2 CPUs
   - Ajuste em: Docker Desktop > Settings > Resources

2. **Portas**
   - Se encontrar conflitos de porta, verifique serviços nativos:
   ```sh
   sudo lsof -i :5432  # POSTGRES
   sudo lsof -i :6379  # Redis
   sudo lsof -i :4566  # LocalStack
   ```

#### Linux

1. **Permissões Docker**
   - Adicione seu usuário ao grupo docker:
   ```sh
   sudo usermod -aG docker $USER
   newgrp docker
   ```

2. **Serviços Locais**
   - Certifique-se que não há instâncias locais do POstgres ou Redis:
   ```sh
   sudo systemctl stop postgres
   sudo systemctl stop redis
   ```

3. **Firewall**
   - Se estiver usando UFW, libere as portas necessárias:
   ```sh
   sudo ufw allow 3306
   sudo ufw allow 6379
   sudo ufw allow 4566
   ```