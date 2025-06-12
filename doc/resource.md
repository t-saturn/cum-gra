# Estructura Go + Fiber + PostgreSQL + gRPC - Sistema de Gestión Centralizada de Usuarios

## Estructura de Directorios Refinada

```
server/
├── logs/
├── cmd/
│   ├── grpc/
│   │   └── main.go
│   ├── migrate/
│   │   └── main.go
│   └── health/                              # [NUEVO] Health check service
│       └── main.go
├── internal/
│   ├── models/
│   │   ├── base.go
│   │   ├── organic_unit.go
│   │   ├── user.go
│   │   ├── system.go
│   │   ├── module.go
│   │   ├── role.go
│   │   ├── permission.go
│   │   ├── structural_position.go
│   │   ├── personnel_movement.go
│   │   ├── audit_log.go
│   │   ├── group.go
│   │   └── history/
│   │       ├── organic_unit_history.go
│   │       ├── user_history.go
│   │       ├── system_history.go
│   │       ├── module_history.go
│   │       ├── role_history.go
│   │       ├── structural_position_history.go
│   │       ├── role_module_access_history.go
│   │       └── user_system_role_history.go
│   ├── repositories/
│   │   ├── interfaces/
│   │   │   ├── organizational.go           # Interface para repos organizacionales
│   │   │   ├── user_management.go          # Interface para repos de usuarios
│   │   │   ├── access_control.go           # Interface para repos de control de acceso
│   │   │   ├── authentication.go           # Interface para repos de autenticación
│   │   │   └── audit.go                    # [NUEVO] Interface para auditoría
│   │   ├── postgresql/
│   │   │   ├── organizational/
│   │   │   │   ├── organic_unit.go
│   │   │   │   ├── structural_position.go
│   │   │   │   └── personnel_movement.go
│   │   │   ├── user_management/
│   │   │   │   ├── user.go
│   │   │   │   └── group.go
│   │   │   ├── access_control/
│   │   │   │   ├── role.go
│   │   │   │   ├── permission.go
│   │   │   │   └── module.go
│   │   │   ├── authentication/
│   │   │   │   └── session.go              # [NUEVO] Gestión de sesiones
│   │   │   ├── audit/                      # [NUEVO] Repositorios de auditoría
│   │   │   │   ├── audit_log.go
│   │   │   │   └── history.go
│   │   │   └── base/                       # [NUEVO] Repositorio base común
│   │   │       └── repository.go
│   │   └── factory.go
│   ├── services/
│   │   ├── organizational/
│   │   │   ├── organic_unit_service.go
│   │   │   ├── structural_position_service.go
│   │   │   └── personnel_movement_service.go
│   │   ├── user_management/
│   │   │   ├── user_service.go
│   │   │   └── group_service.go
│   │   ├── access_control/
│   │   │   ├── role_service.go
│   │   │   ├── permission_service.go
│   │   │   └── authorization_service.go    # [NUEVO] Lógica de autorización
│   │   ├── authentication/
│   │   │   ├── auth_service.go
│   │   │   └── session_service.go          # [NUEVO] Gestión de sesiones
│   │   ├── security/
│   │   │   ├── audit_service.go
│   │   │   ├── encryption_service.go       # [NUEVO] Servicios de encriptación
│   │   │   └── validation_service.go       # [NUEVO] Validaciones de seguridad
│   │   ├── notification/
│   │   │   └── notification_service.go
│   │   └── group/
│   │       └── group_service.go
│   ├── grpc/
│   │   ├── server/                         # [NUEVO] Configuración del servidor gRPC
│   │   │   ├── server.go
│   │   │   └── interceptors/               # [MODIFICADO] Movido aquí
│   │   │       ├── auth.go
│   │   │       ├── logging.go
│   │   │       ├── recovery.go
│   │   │       ├── validation.go           # [NUEVO] Interceptor de validación
│   │   │       └── metrics.go              # [NUEVO] Interceptor de métricas
│   │   ├── handlers/                       # [NUEVO] Handlers gRPC (implementan proto services)
│   │   │   ├── organizational/
│   │   │   │   ├── organic_unit_handler.go
│   │   │   │   ├── structural_position_handler.go
│   │   │   │   └── personnel_movement_handler.go
│   │   │   ├── user_management/
│   │   │   │   ├── user_handler.go
│   │   │   │   └── group_handler.go
│   │   │   ├── access_control/
│   │   │   │   ├── role_handler.go
│   │   │   │   ├── permission_handler.go
│   │   │   │   └── authorization_handler.go
│   │   │   ├── authentication/
│   │   │   │   └── auth_handler.go
│   │   │   ├── security/
│   │   │   │   └── audit_handler.go
│   │   │   └── health/                     # [NUEVO] Health check handler
│   │   │       └── health_handler.go
│   │   └── middleware/                     # [NUEVO] Middleware específico de gRPC
│   │       ├── context.go                  # Manejo de contexto
│   │       └── metadata.go                 # Manejo de metadata
│   ├── external/
│   │   ├── notification/
│   │   │   ├── email/
│   │   │   │   └── smtp_client.go
│   │   │   └── sms/
│   │   │       └── sms_client.go
│   │   ├── storage/
│   │   │   ├── s3/
│   │   │   │   └── s3_client.go
│   │   │   └── local/
│   │   │       └── file_storage.go
│   │   └── integration/
│   │       ├── ldap/                       # [NUEVO] Integración LDAP
│   │       │   └── ldap_client.go
│   │       └── third_party/
│   │           └── api_client.go
│   ├── domain/
│   │   ├── events/
│   │   │   ├── user_events.go
│   │   │   ├── role_events.go
│   │   │   └── audit_events.go
│   │   ├── policies/
│   │   │   ├── access_policy.go
│   │   │   ├── password_policy.go          # [NUEVO] Políticas de contraseña
│   │   │   └── audit_policy.go
│   │   └── specifications/                 # [NUEVO] Domain specifications
│   │       ├── user_spec.go
│   │       └── role_spec.go
│   └── dto/                                # [NUEVO] Data Transfer Objects
│       ├── organizational/
│       ├── user_management/
│       ├── access_control/
│       ├── authentication/
│       └── common/
├── pkg/
│   ├── config/
│   │   ├── config.go
│   │   ├── grpc.go                         # [NUEVO] Configuración específica gRPC
│   │   └── database.go
│   ├── database/
│   │   ├── postgresql/
│   │   │   ├── connection.go
│   │   │   └── migration.go
│   │   └── transaction/                    # [NUEVO] Manejo de transacciones
│   │       └── manager.go
│   ├── cache/
│   │   ├── redis/
│   │   │   └── client.go
│   │   └── memory/
│   │       └── cache.go
│   ├── logger/
│   │   ├── logger.go
│   │   ├── grpc.go                         # [NUEVO] Logger específico para gRPC
│   │   └── structured.go
│   ├── validator/
│   │   ├── validator.go
│   │   └── grpc.go                         # [NUEVO] Validador para protobuf
│   ├── security/
│   │   ├── jwt/
│   │   │   └── jwt.go
│   │   ├── crypto/
│   │   │   ├── hash.go
│   │   │   └── encryption.go
│   │   └── rbac/                           # [NUEVO] Role-Based Access Control
│   │       └── rbac.go
│   ├── utils/
│   │   ├── converter/                      # [NUEVO] Conversores proto <-> domain
│   │   │   ├── user.go
│   │   │   ├── role.go
│   │   │   └── organizational.go
│   │   ├── pagination/
│   │   │   └── pagination.go
│   │   └── response/
│   │       └── grpc_response.go            # [NUEVO] Helpers para respuestas gRPC
│   ├── constants/
│   │   ├── errors.go
│   │   ├── roles.go
│   │   ├── permissions.go
│   │   └── grpc_codes.go                   # [NUEVO] Códigos de estado gRPC
│   └── errors/
│       ├── domain_errors.go
│       ├── grpc_errors.go                  # [NUEVO] Errores específicos gRPC
│       └── handler.go
├── proto/
│   ├── organizational/
│   │   ├── organic_unit.proto
│   │   ├── structural_position.proto
│   │   └── personnel_movement.proto
│   ├── user_management/
│   │   ├── user.proto
│   │   └── group.proto
│   ├── access_control/
│   │   ├── role.proto
│   │   ├── permission.proto
│   │   └── authorization.proto             # [NUEVO] Servicio de autorización
│   ├── authentication/
│   │   └── auth.proto
│   ├── security/
│   │   └── audit.proto
│   ├── common/
│   │   ├── types.proto                     # Tipos comunes
│   │   ├── pagination.proto                # [NUEVO] Paginación
│   │   ├── timestamp.proto                 # [NUEVO] Timestamps
│   │   └── error.proto                     # [NUEVO] Errores comunes
│   └── health/                             # [NUEVO] Health check service
│       └── health.proto
├── pb/                                     # Código generado desde proto
│   ├── organizational/
│   ├── user_management/
│   ├── access_control/
│   ├── authentication/
│   ├── security/
│   ├── common/
│   └── health/
├── migrations/
│   ├── postgresql/
│   │   ├── 001_initial_schema.up.sql
│   │   ├── 001_initial_schema.down.sql
│   │   ├── 002_add_audit_tables.up.sql
│   │   └── 002_add_audit_tables.down.sql
│   └── seeds/                              # [MODIFICADO] Movido aquí
│       ├── development/
│       │   ├── users.sql
│       │   └── roles.sql
│       └── production/
│           └── initial_admin.sql
├── tests/
│   ├── unit/
│   │   ├── services/
│   │   ├── repositories/
│   │   └── handlers/                       # [NUEVO] Tests para handlers gRPC
│   ├── integration/
│   │   ├── grpc/
│   │   │   ├── organizational_test.go
│   │   │   ├── user_management_test.go
│   │   │   ├── access_control_test.go
│   │   │   └── authentication_test.go
│   │   └── database/
│   ├── e2e/
│   │   ├── scenarios/
│   │   └── grpc_client/                    # [NUEVO] Cliente de prueba gRPC
│   ├── fixtures/
│   │   ├── database/
│   │   └── proto/                          # [NUEVO] Fixtures para protobuf
│   └── mocks/
│       ├── repositories/
│       ├── services/
│       └── grpc/                           # [NUEVO] Mocks para servicios gRPC
├── scripts/
│   ├── build/
│   │   ├── build.sh
│   │   └── docker.sh
│   ├── database/
│   │   ├── migrate.sh
│   │   ├── seed.sh
│   │   └── backup.sh
│   ├── development/
│   │   ├── setup.sh
│   │   ├── proto_gen.sh                    # [NUEVO] Generación de protobuf
│   │   └── mock_gen.sh                     # [NUEVO] Generación de mocks
│   └── deployment/
│       ├── deploy.sh
│       └── rollback.sh
├── docker/
│   ├── Dockerfile
│   ├── Dockerfile.dev
│   ├── docker-compose.yml
│   ├── docker-compose.dev.yml
│   └── grpc/                               # [NUEVO] Configuración específica gRPC
│       └── Dockerfile.grpc
├── configs/
│   ├── config.yaml
│   ├── config.dev.yaml
│   ├── config.prod.yaml
│   └── grpc/                               # [NUEVO] Configuraciones gRPC
│       ├── server.yaml
│       └── interceptors.yaml
├── docs/
│   ├── grpc/
│   │   ├── services.md                     # Documentación de servicios
│   │   ├── interceptors.md                 # Documentación de interceptors
│   │   └── examples/                       # [NUEVO] Ejemplos de uso
│   ├── architecture/
│   │   ├── overview.md
│   │   ├── domain_model.md
│   │   └── security_model.md
│   ├── deployment/
│   │   ├── production.md
│   │   └── development.md
│   └── user_guides/
│       ├── admin.md
│       └── developer.md
├── deployments/
│   ├── kubernetes/
│   │   ├── deployment.yaml
│   │   ├── service.yaml
│   │   └── configmap.yaml
│   ├── helm/
│   └── terraform/
├── monitoring/
│   ├── prometheus/
│   │   └── rules.yaml
│   ├── grafana/
│   │   └── dashboards/
│   └── jaeger/                             # [NUEVO] Tracing distribuido
│       └── config.yaml
├── .github/
│   └── workflows/
│       ├── ci.yml
│       ├── cd.yml
│       └── proto-lint.yml                  # [NUEVO] Linting de protobuf
├── .env.example
├── .env.dev.example
├── .env.prod.example
├── .golangci.yml
├── buf.yaml                                # [NUEVO] Configuración de buf para protobuf
├── buf.gen.yaml                            # [NUEVO] Generación de código protobuf
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
