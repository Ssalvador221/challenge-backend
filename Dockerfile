# Usa uma imagem base do Go
FROM golang:1.23-alpine AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos do módulo Go e baixa as dependências
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copia o código fonte para o contêiner
COPY ./src/ ./src/

# Compila o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -o contact-form ./src

# Imagem final
FROM alpine:latest

# Instala pacotes necessários para SSL
RUN apk --no-cache add ca-certificates

# Copia o binário compilado da imagem builder
COPY --from=builder /app/contact-form /app/contact-form

# Define a porta que a aplicação escutará
EXPOSE 80

# Comando para executar o aplicativo
CMD ["/app/contact-form"]
