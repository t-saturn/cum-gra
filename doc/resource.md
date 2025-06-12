# Estructura Go + Fiber + PostgreSQL + gRPC - Sistema de Gestión Centralizada de Usuarios

```txt
central-user-manager/
├── logs/                          # Logs diarios del sistema
├── cmd/
│   ├── api/
│   │   └── main.go               # Servidor HTTP con Fiber
│   ├── grpc/
│   │   └── main.go               # Servidor gRPC
│   └── migrate/
│       └── main.go               # Herramienta de migraciones
├── internal/
│   ├── models/                    # Modelos GORM para todas las entidades
│   │   ├── base.go               # Modelo base con campos comunes
│   │   ├── organic_unit.go       # Unidades orgánicas
│   │   ├── user.go               # Usuarios
│   │   ├── system.go             # Sistemas
│   │   ├── module.go             # Módulos de sistemas
│   │   ├── role.go               # Roles
│   │   ├── permission.go         # Permisos
│   │   ├── structural_position.go # Posiciones estructurales
│   │   ├── personnel_movement.go  # Movimientos de personal
│   │   ├── session.go            # Sesiones activas
│   │   ├── audit_log.go          # Logs de auditoría
│   │   ├── verification_token.go # Tokens de verificación
│   │   ├── mfa_device.go         # Dispositivos MFA
│   │   ├── group.go              # Grupos de usuarios
│   │   ├── api_token.go          # Tokens API
│   │   └── history/              # Modelos de historial
│   │       ├── organic_unit_history.go
│   │       ├── user_history.go
│   │       ├── system_history.go
│   │       ├── module_history.go
│   │       ├── role_history.go
│   │       ├── structural_position_history.go
│   │       ├── role_module_access_history.go
│   │       └── user_system_role_history.go
│   ├── repositories/              # Acceso a datos
│   │   ├── interfaces/           # Interfaces de repositorios
│   │   │   ├── organic_unit_repository.go
│   │   │   ├── user_repository.go
│   │   │   ├── system_repository.go
│   │   │   ├── module_repository.go
│   │   │   ├── role_repository.go
│   │   │   ├── permission_repository.go
│   │   │   ├── structural_position_repository.go
│   │   │   ├── personnel_movement_repository.go
│   │   │   ├── session_repository.go
│   │   │   ├── audit_log_repository.go
│   │   │   ├── verification_token_repository.go
│   │   │   ├── mfa_device_repository.go
│   │   │   ├── group_repository.go
│   │   │   └── api_token_repository.go
│   │   ├── postgresql/           # Implementaciones PostgreSQL
│   │   │   ├── organic_unit_repository.go
│   │   │   ├── user_repository.go
│   │   │   ├── system_repository.go
│   │   │   ├── module_repository.go
│   │   │   ├── role_repository.go
│   │   │   ├── permission_repository.go
│   │   │   ├── structural_position_repository.go
│   │   │   ├── personnel_movement_repository.go
│   │   │   ├── session_repository.go
│   │   │   ├── audit_log_repository.go
│   │   │   ├── verification_token_repository.go
│   │   │   ├── mfa_device_repository.go
│   │   │   ├── group_repository.go
│   │   │   └── api_token_repository.go
│   │   └── factory.go            # Factory de repositorios
│   ├── services/                  # Lógica de negocio
│   │   ├── organizational/       # Servicios organizacionales
│   │   │   ├── organic_unit_service.go
│   │   │   ├── structural_position_service.go
│   │   │   └── personnel_movement_service.go
│   │   ├── user_management/      # Servicios de gestión de usuarios
│   │   │   ├── user_service.go
│   │   │   ├── user_position_service.go
│   │   │   └── user_profile_service.go
│   │   ├── access_control/       # Servicios de control de acceso
│   │   │   ├── system_service.go
│   │   │   ├── module_service.go
│   │   │   ├── role_service.go
│   │   │   ├── permission_service.go
│   │   │   ├── role_assignment_service.go
│   │   │   └── access_validation_service.go
│   │   ├── authentication/       # Servicios de autenticación
│   │   │   ├── auth_service.go
│   │   │   ├── session_service.go
│   │   │   ├── token_service.go
│   │   │   ├── mfa_service.go
│   │   │   └── password_service.go
│   │   ├── security/            # Servicios de seguridad
│   │   │   ├── audit_service.go
│   │   │   ├── security_policy_service.go
│   │   │   └── api_token_service.go
│   │   ├── notification/        # Servicios de notificación
│   │   │   ├── email_service.go
│   │   │   ├── sms_service.go
│   │   │   └── notification_service.go
│   │   └── group/              # Servicios de grupos
│   │       └── group_service.go
│   ├── handlers/                # HTTP handlers (Fiber)
│   │   ├── api/
│   │   │   ├── v1/
│   │   │   │   ├── organizational/
│   │   │   │   │   ├── organic_unit_handler.go
│   │   │   │   │   ├── structural_position_handler.go
│   │   │   │   │   └── personnel_movement_handler.go
│   │   │   │   ├── user_management/
│   │   │   │   │   ├── user_handler.go
│   │   │   │   │   ├── user_position_handler.go
│   │   │   │   │   └── user_profile_handler.go
│   │   │   │   ├── access_control/
│   │   │   │   │   ├── system_handler.go
│   │   │   │   │   ├── module_handler.go
│   │   │   │   │   ├── role_handler.go
│   │   │   │   │   ├── permission_handler.go
│   │   │   │   │   └── role_assignment_handler.go
│   │   │   │   ├── authentication/
│   │   │   │   │   ├── auth_handler.go
│   │   │   │   │   ├── session_handler.go
│   │   │   │   │   ├── mfa_handler.go
│   │   │   │   │   └── password_handler.go
│   │   │   │   ├── security/
│   │   │   │   │   ├── audit_handler.go
│   │   │   │   │   └── api_token_handler.go
│   │   │   │   └── group/
│   │   │   │       └── group_handler.go
│   │   │   └── health_handler.go
│   │   └── admin/               # Handlers administrativos
│   │       └── system_admin_handler.go
│   ├── grpc/                    # gRPC handlers
│   │   ├── organizational/
│   │   │   ├── organic_unit_handler.go
│   │   │   └── structural_position_handler.go
│   │   ├── user_management/
│   │   │   └── user_handler.go
│   │   ├── access_control/
│   │   │   ├── system_handler.go
│   │   │   ├── module_handler.go
│   │   │   ├── role_handler.go
│   │   │   └── permission_handler.go
│   │   ├── authentication/
│   │   │   ├── auth_handler.go
│   │   │   └── session_handler.go
│   │   └── security/
│   │       └── audit_handler.go
│   ├── middleware/              # HTTP middleware
│   │   ├── auth/
│   │   │   ├── jwt_auth.go
│   │   │   ├── session_auth.go
│   │   │   ├── api_token_auth.go
│   │   │   └── mfa_auth.go
│   │   ├── authorization/
│   │   │   ├── rbac.go
│   │   │   ├── system_access.go
│   │   │   └── module_access.go
│   │   ├── security/
│   │   │   ├── rate_limit.go
│   │   │   ├── audit_log.go
│   │   │   ├── cors.go
│   │   │   └── security_headers.go
│   │   ├── validation/
│   │   │   ├── request_validator.go
│   │   │   └── input_sanitizer.go
│   │   └── logging/
│   │       └── logger.go
│   ├── routes/                  # Rutas HTTP
│   │   ├── api/
│   │   │   ├── v1/
│   │   │   │   ├── api_routes.go
│   │   │   │   ├── organizational_routes.go
│   │   │   │   ├── user_management_routes.go
│   │   │   │   ├── access_control_routes.go
│   │   │   │   ├── authentication_routes.go
│   │   │   │   ├── security_routes.go
│   │   │   │   └── group_routes.go
│   │   │   └── router.go
│   │   └── admin/
│   │       └── admin_routes.go
│   ├── dto/                     # Data Transfer Objects
│   │   ├── organizational/
│   │   │   ├── organic_unit_dto.go
│   │   │   ├── structural_position_dto.go
│   │   │   └── personnel_movement_dto.go
│   │   ├── user_management/
│   │   │   ├── user_dto.go
│   │   │   ├── user_position_dto.go
│   │   │   └── user_profile_dto.go
│   │   ├── access_control/
│   │   │   ├── system_dto.go
│   │   │   ├── module_dto.go
│   │   │   ├── role_dto.go
│   │   │   ├── permission_dto.go
│   │   │   └── role_assignment_dto.go
│   │   ├── authentication/
│   │   │   ├── auth_dto.go
│   │   │   ├── session_dto.go
│   │   │   ├── token_dto.go
│   │   │   └── mfa_dto.go
│   │   ├── security/
│   │   │   ├── audit_dto.go
│   │   │   └── api_token_dto.go
│   │   ├── group/
│   │   │   └── group_dto.go
│   │   └── common/
│   │       ├── pagination_dto.go
│   │       ├── filter_dto.go
│   │       └── response_dto.go
│   ├── external/                # Servicios externos
│   │   ├── notification/
│   │   │   ├── email/
│   │   │   │   ├── smtp_client.go
│   │   │   │   └── template_engine.go
│   │   │   └── sms/
│   │   │       └── sms_client.go
│   │   ├── storage/
│   │   │   ├── file_storage.go
│   │   │   └── image_processor.go
│   │   └── integration/
│   │       ├── ldap_client.go
│   │       └── external_system_client.go
│   └── domain/                  # Lógica de dominio
│       ├── events/
│       │   ├── user_events.go
│       │   ├── role_events.go
│       │   └── audit_events.go
│       └── policies/
│           ├── password_policy.go
│           ├── session_policy.go
│           └── access_policy.go
├── pkg/                         # Utilidades compartidas
│   ├── config/
│   │   ├── config.go
│   │   ├── database.go
│   │   ├── redis.go
│   │   ├── notification.go
│   │   └── security.go
│   ├── database/
│   │   ├── connection.go
│   │   ├── migration.go
│   │   ├── seeder.go
│   │   └── transaction.go
│   ├── cache/
│   │   ├── redis_client.go
│   │   ├── cache_manager.go
│   │   └── session_store.go
│   ├── logger/
│   │   ├── logger.go
│   │   ├── audit_logger.go
│   │   └── structured_logger.go
│   ├── validator/
│   │   ├── validator.go
│   │   ├── custom_validators.go
│   │   └── sanitizer.go
│   ├── security/
│   │   ├── jwt/
│   │   │   ├── jwt_manager.go
│   │   │   └── token_validator.go
│   │   ├── encryption/
│   │   │   ├── password_hasher.go
│   │   │   ├── crypto_utils.go
│   │   │   └── key_manager.go
│   │   ├── mfa/
│   │   │   ├── totp_generator.go
│   │   │   ├── sms_sender.go
│   │   │   └── backup_codes.go
│   │   └── rate_limiter/
│   │       └── rate_limiter.go
│   ├── utils/
│   │   ├── response/
│   │   │   ├── api_response.go
│   │   │   └── error_response.go
│   │   ├── pagination/
│   │   │   └── paginator.go
│   │   ├── converter/
│   │   │   └── type_converter.go
│   │   ├── time/
│   │   │   └── time_utils.go
│   │   └── string/
│   │       └── string_utils.go
│   ├── constants/
│   │   ├── status_constants.go
│   │   ├── role_constants.go
│   │   ├── permission_constants.go
│   │   └── error_constants.go
│   └── errors/
│       ├── custom_errors.go
│       ├── error_handler.go
│       └── error_codes.go
├── proto/                       # Protocol Buffers
│   ├── organizational/
│   │   ├── organic_unit.proto
│   │   └── structural_position.proto
│   ├── user_management/
│   │   └── user.proto
│   ├── access_control/
│   │   ├── system.proto
│   │   ├── module.proto
│   │   ├── role.proto
│   │   └── permission.proto
│   ├── authentication/
│   │   ├── auth.proto
│   │   └── session.proto
│   ├── security/
│   │   └── audit.proto
│   └── common/
│       ├── common.proto
│       ├── pagination.proto
│       └── timestamp.proto
├── pb/                          # Generated protobuf files
│   ├── organizational/
│   │   ├── organic_unit.pb.go
│   │   ├── organic_unit_grpc.pb.go
│   │   ├── structural_position.pb.go
│   │   └── structural_position_grpc.pb.go
│   ├── user_management/
│   │   ├── user.pb.go
│   │   └── user_grpc.pb.go
│   ├── access_control/
│   │   ├── system.pb.go
│   │   ├── system_grpc.pb.go
│   │   ├── module.pb.go
│   │   ├── module_grpc.pb.go
│   │   ├── role.pb.go
│   │   ├── role_grpc.pb.go
│   │   ├── permission.pb.go
│   │   └── permission_grpc.pb.go
│   ├── authentication/
│   │   ├── auth.pb.go
│   │   ├── auth_grpc.pb.go
│   │   ├── session.pb.go
│   │   └── session_grpc.pb.go
│   ├── security/
│   │   ├── audit.pb.go
│   │   └── audit_grpc.pb.go
│   └── common/
│       ├── common.pb.go
│       ├── pagination.pb.go
│       └── timestamp.pb.go
├── migrations/                  # Migraciones de base de datos
│   ├── 001_create_organic_units_table.sql
│   ├── 002_create_organic_units_history_table.sql
│   ├── 003_create_users_table.sql
│   ├── 004_create_users_history_table.sql
│   ├── 005_create_systems_table.sql
│   ├── 006_create_systems_history_table.sql
│   ├── 007_create_modules_table.sql
│   ├── 008_create_modules_history_table.sql
│   ├── 009_create_roles_table.sql
│   ├── 010_create_roles_history_table.sql
│   ├── 011_create_role_module_access_table.sql
│   ├── 012_create_role_module_access_history_table.sql
│   ├── 013_create_user_system_roles_table.sql
│   ├── 014_create_user_system_roles_history_table.sql
│   ├── 015_create_permissions_table.sql
│   ├── 016_create_role_permissions_table.sql
│   ├── 017_create_module_permissions_table.sql
│   ├── 018_create_structural_positions_table.sql
│   ├── 019_create_structural_positions_history_table.sql
│   ├── 020_create_user_structural_positions_table.sql
│   ├── 021_create_personnel_movements_table.sql
│   ├── 022_create_password_history_table.sql
│   ├── 023_create_active_sessions_table.sql
│   ├── 024_create_session_history_table.sql
│   ├── 025_create_audit_logs_table.sql
│   ├── 026_create_verification_tokens_table.sql
│   ├── 027_create_mfa_devices_table.sql
│   ├── 028_create_groups_table.sql
│   ├── 029_create_user_groups_table.sql
│   ├── 030_create_session_expiration_policies_table.sql
│   ├── 031_create_api_tokens_table.sql
│   └── 032_create_indexes_and_constraints.sql
├── seeds/                       # Datos semilla
│   ├── organic_units_seed.sql
│   ├── users_seed.sql
│   ├── systems_seed.sql
│   ├── modules_seed.sql
│   ├── roles_seed.sql
│   ├── permissions_seed.sql
│   └── admin_user_seed.sql
├── tests/                       # Tests
│   ├── unit/
│   │   ├── services/
│   │   │   ├── organizational/
│   │   │   ├── user_management/
│   │   │   ├── access_control/
│   │   │   ├── authentication/
│   │   │   └── security/
│   │   ├── repositories/
│   │   │   └── postgresql/
│   │   └── utils/
│   ├── integration/
│   │   ├── api/
│   │   │   ├── v1/
│   │   │   │   ├── organizational_test.go
│   │   │   │   ├── user_management_test.go
│   │   │   │   ├── access_control_test.go
│   │   │   │   ├── authentication_test.go
│   │   │   │   └── security_test.go
│   │   │   └── admin_test.go
│   │   ├── grpc/
│   │   └── database/
│   ├── e2e/
│   │   ├── user_management_flow_test.go
│   │   ├── role_assignment_flow_test.go
│   │   └── authentication_flow_test.go
│   ├── fixtures/
│   │   ├── users.json
│   │   ├── roles.json
│   │   └── permissions.json
│   └── mocks/
│       ├── repositories/
│       ├── services/
│       └── external/
├── scripts/                     # Scripts utilitarios
│   ├── build/
│   │   ├── build.sh
│   │   ├── build-docker.sh
│   │   └── cross-compile.sh
│   ├── database/
│   │   ├── migrate.sh
│   │   ├── rollback.sh
│   │   ├── seed.sh
│   │   └── backup.sh
│   ├── development/
│   │   ├── dev-setup.sh
│   │   ├── generate-proto.sh
│   │   ├── generate-mocks.sh
│   │   └── lint.sh
│   └── deployment/
│       ├── deploy.sh
│       ├── health-check.sh
│       └── log-rotate.sh
├── docker/
│   ├── Dockerfile
│   ├── Dockerfile.dev
│   ├── docker-compose.yml
│   ├── docker-compose.dev.yml
│   ├── docker-compose.test.yml
│   └── nginx/
│       ├── nginx.conf
│       └── ssl/
├── configs/
│   ├── config.yaml
│   ├── config.dev.yaml
│   ├── config.test.yaml
│   ├── config.prod.yaml
│   └── security/
│       ├── jwt-keys/
│       │   ├── private.key
│       │   └── public.key
│       └── tls/
│           ├── server.crt
│           └── server.key
├── docs/                        # Documentación
│   ├── api/
│   │   ├── openapi.yaml
│   │   └── postman_collection.json
│   ├── grpc/
│   │   └── grpc_documentation.md
│   ├── architecture/
│   │   ├── system_architecture.md
│   │   ├── database_schema.md
│   │   └── security_model.md
│   ├── deployment/
│   │   ├── installation_guide.md
│   │   ├── configuration_guide.md
│   │   └── troubleshooting.md
│   └── user_guides/
│       ├── admin_guide.md
│       ├── user_guide.md
│       └── api_guide.md
├── deployments/                 # Configuraciones de despliegue
│   ├── kubernetes/
│   │   ├── namespace.yaml
│   │   ├── configmap.yaml
│   │   ├── secret.yaml
│   │   ├── deployment.yaml
│   │   ├── service.yaml
│   │   └── ingress.yaml
│   ├── terraform/
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   └── outputs.tf
│   └── ansible/
│       ├── playbook.yml
│       └── inventory/
├── monitoring/                  # Configuraciones de monitoreo
│   ├── prometheus/
│   │   └── config.yml
│   ├── grafana/
│   │   └── dashboards/
│   └── alertmanager/
│       └── config.yml
├── .github/                     # GitHub Actions
│   └── workflows/
│       ├── ci.yml
│       ├── cd.yml
│       └── security-scan.yml
├── .env.example
├── .env.dev.example
├── .env.prod.example
├── .gitignore
├── .golangci.yml
├── Makefile
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## EXPLICACIÓN DETALLADA DE LA ESTRUCTURA

### **1. Organización por Dominios**

La estructura está organizada por dominios funcionales principales:

#### **`organizational/`** - Gestión Organizacional

- **Unidades Orgánicas**: Jerarquía organizacional
- **Posiciones Estructurales**: Cargos dentro de las unidades
- **Movimientos de Personal**: Cambios y transferencias

#### **`user_management/`** - Gestión de Usuarios

- **Usuarios**: CRUD y gestión de perfiles
- **Posiciones de Usuario**: Asignación de cargos
- **Perfiles**: Gestión de información personal

#### **`access_control/`** - Control de Acceso

- **Sistemas**: Gestión de sistemas integrados
- **Módulos**: Funcionalidades de cada sistema
- **Roles**: Definición de roles por sistema
- **Permisos**: Permisos granulares
- **Asignaciones**: Gestión de roles y accesos

#### **`authentication/`** - Autenticación

- **Autenticación**: Login/logout y validación
- **Sesiones**: Gestión de sesiones activas
- **Tokens**: JWT y tokens de API
- **MFA**: Autenticación multifactor
- **Contraseñas**: Políticas y gestión

#### **`security/`** - Seguridad

- **Auditoría**: Logs y trazabilidad completa
- **Tokens API**: Gestión de acceso programático
- **Políticas**: Configuración de seguridad

#### **`group/`** - Grupos

- **Grupos**: Agrupación lógica de usuarios

### **2. Características Principales**

#### **Auditoría Completa**

- Todas las tablas principales tienen su tabla `_history`
- Logs de auditoría para todas las operaciones
- Trazabilidad completa de cambios

#### **Soft Delete**

- Implementación de borrado lógico en todas las entidades
- Campos `is_deleted`, `deleted_at`, `deleted_by`
- Preservación de integridad referencial

#### **Versionado**

- Control de versiones para detectar conflictos
- Campo `version` en entidades principales

#### **Seguridad Robusta**

- Autenticación multifactor
- Gestión de sesiones con políticas configurables
- Tokens API con scopes
- Rate limiting y protección CORS

#### **Escalabilidad**

- Separación clara de responsabilidades
- Interfaces para fácil testing y mocking
- Preparado para microservicios
- Cache con Redis para mejor rendimiento

### **3. Modelos GORM Principales**

#### **Modelo Base**

```go
// internal/models/base.go
type BaseModel struct {
    ID        uint       `json:"id" gorm:"primaryKey"`
    IsDeleted bool       `json:"-" gorm:"not null;default:false"`
    DeletedAt *time.Time `json:"-"`
    DeletedBy *uint      `json:"-"`
    Version   int        `json:"version" gorm:"not null;default:1"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    CreatedBy *uint      `json:"created_by"`
    UpdatedBy *uint      `json:"updated_by"`
}
```

#### **Usuario Completo**

```go
// internal/models/user.go
type User struct {
    BaseModel
    Username        string         `json:"username" gorm:"uniqueIndex:idx_username_deleted;size:50;not null"`
    Email           string         `json:"email" gorm:"uniqueIndex:idx_email_deleted;size:255;not null"`
    PasswordHash    string         `json:"-" gorm:"size:255;not null"`
    FirstName       string         `json:"first_name" gorm:"size:50;not null"`
    LastName        string         `json:"last_name" gorm:"size:50;not null"`
    FullName        string         `json:"full_name" gorm:"size:255;not null"`
    Phone           string         `json:"phone" gorm:"size:20"`
    Address         string         `json:"address" gorm:"type:text"`
    ProfilePicture  string         `json:"profile_picture" gorm:"size:500"`
    OrganicUnitID   *uint          `json:"organic_unit_id"`
    Status          int16          `json:"status" gorm:"not null;default:1"`

    // Relaciones
    OrganicUnit     *OrganicUnit   `json:"organic_unit,omitempty"`
    SystemRoles     []UserSystemRole `json:"system_roles,omitempty"`
    Positions       []UserStructuralPosition `json:"positions,omitempty"`
    Groups          []UserGroup    `json:"groups,omitempty"`
```
