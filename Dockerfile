FROM docker.io/golang:1.24.2-alpine AS builder

WORKDIR /app

# Copiar apenas os arquivos necessários para baixar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código fonte e arquivos de configuração
COPY . .

# Compilar a aplicação com flags para reduzir o tamanho do binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./cmd/main.go

# Imagem final mais leve
FROM alpine:3.19

# Instalar apenas o necessário e limpar o cache
RUN apk add --no-cache ca-certificates tzdata && \
    rm -rf /var/cache/apk/* && \
    rm -rf /tmp/* && \
    rm -rf /var/tmp/*

WORKDIR /app

# Copiar apenas o binário compilado e arquivos de configuração
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Criar um usuário não root para executar a aplicação
RUN adduser -D apiuser
USER apiuser

# Expor a porta 8080
EXPOSE 4000 80

# Comando para executar a aplicação
CMD ["./main"]
