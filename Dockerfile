# ============ Builder ============
FROM golang:1.24-bookworm AS builder
# Alternativa ultra exacta:
# FROM golang:1.24.4-bookworm AS builder

ENV CGO_ENABLED=0 \
  GO111MODULE=on \
  GOTOOLCHAIN=auto

WORKDIR /src

# Cache de dependencias
COPY go.mod ./
# Si tienes go.sum, mejor incluirlo para cache más estable
# COPY go.sum ./
RUN go mod download

# Copiamos el resto del código
COPY . .

# Compilar binario estatico
RUN go build -o /out/server ./cmd/server/main.go

# ============ Runtime ============
# Si tu app hace llamadas HTTPS, usa static-debian12; si no, puedes dejar "static"
FROM gcr.io/distroless/static-debian12:nonroot
# Alternativa minimalista sin CA certs:
# FROM gcr.io/distroless/static:nonroot

WORKDIR /app

# Copiamos el binario
COPY --from=builder /out/server /app/server

# El volumen montará /app/keys
EXPOSE 9190
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
