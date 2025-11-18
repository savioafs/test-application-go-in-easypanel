# syntax=docker/dockerfile:1

# Etapa 1 – build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Dependências
COPY go.mod go.sum ./
RUN go mod download

# Código
COPY . .

# Build binário estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Etapa 2 – imagem final mínima
FROM alpine:3.20

WORKDIR /app

# Copia binário da etapa de build
COPY --from=builder /app/app .

# Define env padrão
ENV GIN_MODE=release
ENV PORT=8080

# Documenta a porta (só informativo, mas bom ter)
EXPOSE 8080

# Comando de start
CMD ["./app"]
