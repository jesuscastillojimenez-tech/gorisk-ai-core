# ==========================================
# STAGE 1: Builder (Node.js)
# ==========================================
FROM node:20-alpine AS builder

WORKDIR /app

# Copiar package.json e instalar dependencias
COPY package.json package-lock.json ./
RUN npm install

# Copiar el resto del código y construir para producción
COPY . .
RUN npm run build

# ==========================================
# STAGE 2: Runner (Nginx)
# ==========================================
FROM nginx:alpine

# Copiar los archivos estáticos generados por Vite (carpeta dist) al directorio de Nginx
COPY --from=builder /app/dist /usr/share/nginx/html

# Exponer el puerto 80 para servir el frontend
EXPOSE 80

# Iniciar Nginx
CMD ["nginx", "-g", "daemon off;"]