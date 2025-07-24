# Etapa 1: Instalar dependencias y construir
FROM node:20 AS build

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

# Etapa 2: Servir los archivos generados
FROM node:20

WORKDIR /app

COPY --from=build /app ./

# Instala Astro como dependencia de producción si usas `npm create astro@latest`
RUN npm install --omit=dev

EXPOSE 4321

CMD ["npx", "astro", "preview", "--host"]

