# ============ Build ============
FROM node:20-alpine AS builder
WORKDIR /app

ENV NEXT_TELEMETRY_DISABLED=1

# Instalar dependencias del sistema necesarias
RUN apk add --no-cache openssl python3 make g++ build-base

# Copiar archivos de dependencias
COPY package*.json ./

# Instalar dependencias
RUN npm install --legacy-peer-deps

# Copiar todo el proyecto
COPY . .

# Build de la app
RUN npm run build

# ============ Runtime ============
FROM node:20-alpine AS runner
WORKDIR /app

ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=6161
ENV HOSTNAME=0.0.0.0

# Instalar utilitarios del sistema
RUN apk add --no-cache openssl curl tzdata

# Configurar zona horaria
RUN cp /usr/share/zoneinfo/America/Lima /etc/localtime && \
    echo "America/Lima" > /etc/timezone

# Copiar desde builder
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next ./.next

EXPOSE 6161

CMD ["npm", "start", "--", "-p", "6161"]