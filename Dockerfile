# Etapa 1 – build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copia go.mod/go.sum e baixa dependências (cache melhor)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Build binário estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Etapa 2 – imagem final mínima
FROM alpine:3.20

WORKDIR /app

# Copia binário
COPY --from=builder /app/app .

# Porta que a app expõe
EXPOSE 8888

# Variáveis (se quiser default)
ENV GIN_MODE=release

CMD ["./app"]
