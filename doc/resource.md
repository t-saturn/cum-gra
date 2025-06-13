# Sistema de Gestión Centralizada de Usuarios (Go + gRPC + PostgreSQL + GORM)

## Estructura del Proyecto

```
server/
├── logs/                                # Archivos de log
├── cmd/
│   └── grpc/
│       └── main.go                      # Punto de entrada del servidor gRPC
├── internal/
│   ├── handlers/
│   ├── models/
│   │   └── user.go                      # Definición de entidades (GORM)
│   ├── services/
│   │   └── user_service.go              # Lógica de negocio (aplicación)
│   ├── repositories/
│   │   └── user_repo.go                 # Acceso a datos vía GORM
│   └── grpc/
│       ├── server.go                    # Configuración del servidor gRPC
│       └── interceptors.go             # Interceptores (auth, logs, recovery, etc.)
├── pkg/
│   ├── logger/
│   ├── config/
│   │   └── config.go                    # Carga de configuración desde .env
│   └── database/
│       └── connection.go               # Conexión a PostgreSQL (GORM)
├── proto/
│   ├── user/                            # Definición del servicio de usuarios
│   └── application/                     # (Opcional) Definición de servicio de autenticación
├── pb/                                  # Código generado desde archivos *.proto
├── tests/
│   ├── unit/                            # Tests unitarios
│   └── integration/                     # Tests de integración (gRPC, DB)
├── .github/
│   └── workflows/
│       ├── ci.yml                       # CI: pruebas, linting, etc.
│       ├── cd.yml                       # CD: despliegue automático
│       └── proto-lint.yml               # Linter para archivos protobuf
├── .env.example                         # Variables de entorno de ejemplo
├── .env.dev.example
├── .env.prod.example
├── .golangci.yml                        # Configuración de GolangCI-Lint
├── buf.yaml                             # Configuración de Buf (protobuf)
├── buf.gen.yaml                         # Generación de código con Buf
├── Makefile                             # Comandos útiles (build, proto-gen, etc.)
├── go.mod
└── go.sum
```

---

## Flujo de una Solicitud gRPC

```
Cliente gRPC Request
    ↓
Interceptors (auth, logging, recovery, validation, metrics)
    ↓
gRPC Handler (internal/grpc/server.go)
    ↓
Service (internal/services/user_service.go)
    ↓
Repository (internal/repositories/user_repo.go)
    ↓
Database (PostgreSQL via GORM)
    ↓
Respuesta → Interceptor (log, audit, trace) → Cliente
```

### Componentes del Flujo

| Componente           | Descripción                                                                                    |
| -------------------- | ---------------------------------------------------------------------------------------------- |
| **Cliente gRPC**     | Una app cliente (CLI, frontend, otro servicio) invoca un RPC definido en `proto/`              |
| **Interceptors**     | Middleware para autenticación, recuperación de pánico, métricas y trazabilidad                 |
| **gRPC Handler**     | Implementa la interfaz generada desde el `.proto`, recibe la solicitud y la reenvía al service |
| **Service Layer**    | Contiene la lógica de negocio, validaciones de dominio y coordinación entre componentes        |
| **Repository Layer** | Abstrae la base de datos usando GORM                                                           |
| **Base de Datos**    | PostgreSQL como almacén relacional persistente                                                 |
| **Respuesta**        | La respuesta se propaga de vuelta por las capas, con logs o auditoría opcional                 |

---

## Flujo de una Solicitud gRPC (con Autenticación y Autorización)

```
Cliente gRPC Request
    ↓
Interceptors (auth/jwt | mTLS, logging, recovery, validation, metrics)
    ↓
Authz Middleware (verifica permisos, roles, scopes)
    ↓
gRPC Handler (internal/grpc/server.go)
    ↓
Service (internal/services/user_service.go)
    ↓
Repository (internal/repositories/user_repo.go)
    ↓
Database (PostgreSQL via GORM)
    ↓
Respuesta → Interceptors (log, audit, trace) → Cliente
```

### 🔐 Seguridad Integrada

| Capa              | Mecanismo         | Descripción                                                                                           |
| ----------------- | ----------------- | ----------------------------------------------------------------------------------------------------- |
| **Autenticación** | `JWT` / `mTLS`    | Se verifica el token (access token OAuth2, JWT), o certificado cliente (mutual TLS) en el interceptor |
| **Autorización**  | `RBAC` / `Scopes` | Verifica si el usuario tiene los permisos/roles necesarios para invocar el método RPC                 |
| **Cifrado**       | `TLS`             | Todo el tráfico gRPC está cifrado usando TLS 1.2+                                                     |
| **Interceptors**  | Middleware        | Logging, recuperación de pánico, validación de entradas, trazabilidad, métricas Prometheus, etc.      |
