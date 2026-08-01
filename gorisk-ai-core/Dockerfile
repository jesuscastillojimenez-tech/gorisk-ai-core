# ==========================================
# STAGE 1: Builder
# ==========================================
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Instalar dependencias necesarias para compilar
RUN apk add --no-cache git

# Copiar archivos de módulos y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar el binario de Go (optimizado para tamaño y eliminando debug info)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o gorisk-api cmd/main.go

# ==========================================
# STAGE 2: Runner
# ==========================================
FROM alpine:latest

WORKDIR /root/

# Instalar certificados raíz para que Go pueda hacer llamadas HTTPS a Gemini y Mongo
RUN apk --no-cache add ca-certificates

# Copiar el binario desde la fase de compilación
COPY --from=builder /app/gorisk-api .

# EXCELENTE PRÁCTICA: No copiamos el .env local. 
# Confiamos en tu main.go que sabe usar las variables del entorno del SO/Docker.

# Exponer el puerto donde corre Gin
EXPOSE 8080

# Comando de inicio
CMD ["./gorisk-api"]