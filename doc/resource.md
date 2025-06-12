# Estructura Go + Fiber + PostgreSQL + gRPC - Sistema de Gestión Centralizada de Usuarios

## Estructura de Directorios Refinada

```
server/
├── logs/
├── cmd/
│   └── grpc/
│       └── main.go               # Arranca el servidor gRPC
├── internal/
│   ├── models/
│   │   └── user.go               # Entidad User
│   ├── services/
│   │   └── user_service.go       # Lógica de negocio
│   ├── repositories/
│   │   └── user_repo.go          # Acceso a datos (GORM)
│   └── grpc/
│       └── server.go             # Configuración de gRPC (interceptors, handlers)
├── pkg/
│   ├── config/
│   │   └── config.go             # Carga de variables de entorno
│   └── database/
│       └── connection.go         # Conexión a PostgreSQL
├── proto/
│   ├── user.proto                # Definición del servicio y mensajes gRPC
│   └── auth.proto                # Auth (si aplica)
├── pb/                           # Código generado desde *.proto
├── migrations/
├── tests/
│   ├── unit/
│   └── integration/
├── .github/
│   └── workflows/
│       ├── ci.yml
│       ├── cd.yml
│       └── proto-lint.yml                  # Linting de protobuf
├── .env.example
├── .env.dev.example
├── .env.prod.example
├── .golangci.yml
├── buf.yaml                                # Configuración de buf para protobuf
├── buf.gen.yaml                            # Generación de código protobuf
├── Makefile
├── go.mod
└── go.sum
```

## Cambios Principales y Justificación

### 1. **Organización de Handlers gRPC**

- **Añadido**: `internal/grpc/handlers/` - Separar handlers de servicios para mejor separación de responsabilidades
- Los handlers implementan las interfaces generadas por protobuf
- Los servicios contienen la lógica de negocio pura

### 2. **Interceptors Mejorados**

- **Movido**: De raíz a `internal/grpc/server/interceptors/`
- **Añadido**: Interceptors específicos para validación y métricas
- Mejor organización de middleware gRPC

### 3. **DTOs (Data Transfer Objects)**

- **Añadido**: `internal/dto/` para objetos de transferencia entre capas
- Facilita conversión entre protobuf, domain models y database models

### 4. **Mejoras en Repositorios**

- **Añadido**: Organización por dominio dentro de postgresql
- **Añadido**: Repositorio base común para operaciones CRUD estándar
- **Añadido**: Interfaces específicas por dominio

### 5. **Utilidades gRPC**

- **Añadido**: Conversores entre protobuf y modelos de dominio
- **Añadido**: Helpers para respuestas gRPC
- **Añadido**: Códigos de estado gRPC específicos

### 6. **Configuración Protobuf**

- **Añadido**: `buf.yaml` y `buf.gen.yaml` para mejor gestión de protobuf
- **Añadido**: Scripts de generación automática

### 7. **Health Check Service**

- **Añadido**: Servicio estándar de health check para gRPC
- Importante para orquestadores como Kubernetes

### 8. **Observabilidad**

- **Añadido**: Configuración para Jaeger (tracing distribuido)
- **Mejorado**: Logging específico para gRPC
- **Añadido**: Métricas específicas

## Flujo de Petición gRPC Actualizado

```
Cliente gRPC Request
    ↓
Interceptors (auth, logging, recovery, validation, metrics)
    ↓
gRPC Handler (organizational/user_handler.go)
    ↓
Service (user_management/user_service.go)
    ↓
Repository (postgresql/user_management/user.go)
    ↓
Database
    ↓
Response + Audit Log
```

Esta estructura proporciona:

- **Mejor separación de responsabilidades**
- **Escalabilidad** por dominios de negocio
- **Testabilidad** mejorada
- **Mantenibilidad** a largo plazo
- **Observabilidad** completa
- **Cumplimiento de estándares gRPC**
