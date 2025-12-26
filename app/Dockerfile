# ============ Deps ============
FROM node:20-bookworm-slim AS deps
WORKDIR /app
ENV NEXT_TELEMETRY_DISABLED=1
COPY package*.json ./
# Usa ci si hay lockfile; si no, instala normal
RUN if [ -f package-lock.json ]; then \
  npm ci --legacy-peer-deps; \
  else \
  npm install --legacy-peer-deps; \
  fi

# ============ Build ============
FROM node:20-bookworm-slim AS builder
WORKDIR /app
ENV NEXT_TELEMETRY_DISABLED=1

# Reutiliza deps ya resueltas
COPY --from=deps /app/node_modules ./node_modules
# Copia el código
COPY . .
# Compila (genera .next)
RUN npm run build

# ============ Runtime ============
FROM node:20-bookworm-slim AS runner
WORKDIR /app
ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1
ENV PORT=6161
ENV HOSTNAME=0.0.0.0

# Instala dependencias en modo producción
COPY package*.json ./
RUN if [ -f package-lock.json ]; then \
  npm ci --omit=dev --legacy-peer-deps; \
  else \
  npm install --omit=dev --legacy-peer-deps; \
  fi

# Copia artefactos de build y estáticos
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public

EXPOSE 6161
CMD ["npm", "start", "--", "-p", "6161"]
