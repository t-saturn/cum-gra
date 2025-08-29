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

# Env públicas usadas en build (lado cliente)
ARG NEXT_PUBLIC_API_BASE
ARG NEXT_PUBLIC_FRONT_BASE
ENV NEXT_PUBLIC_API_BASE=${NEXT_PUBLIC_API_BASE}
ENV NEXT_PUBLIC_FRONT_BASE=${NEXT_PUBLIC_FRONT_BASE}

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
ENV PORT=6160
ENV HOSTNAME=0.0.0.0

# Reinyecta públicas si las necesitas para SSR/handlers
ARG NEXT_PUBLIC_API_BASE
ARG NEXT_PUBLIC_FRONT_BASE
ENV NEXT_PUBLIC_API_BASE=${NEXT_PUBLIC_API_BASE}
ENV NEXT_PUBLIC_FRONT_BASE=${NEXT_PUBLIC_FRONT_BASE}

# Paquetes para runtime (usa deps si tenías lockfile)
COPY package*.json ./
RUN if [ -f package-lock.json ]; then \
  npm ci --omit=dev --legacy-peer-deps; \
  else \
  npm install --omit=dev --legacy-peer-deps; \
  fi

# Copia artefactos de build y estáticos
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public

EXPOSE 6160
CMD ["npm", "start", "--", "-p", "6160"]
