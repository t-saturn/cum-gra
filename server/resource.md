# Arquitectura Hexagonal Go + Fiber + PostgreSQL + gRPC

```txt
server/
├── cmd/
│   ├── api/
│   │   └── main.go                 # HTTP API server entry point
│   ├── grpc/
│   │   └── main.go                 # gRPC server entry point
│   └── gateway/
│       └── main.go                 # gRPC-Gateway entry point (futuro)
├── internal/
│   ├── core/                       # DOMINIO (Centro del hexágono)
│   │   ├── domain/
│   │   │   ├── entities/
│   │   │   │   ├── user.go
│   │   │   │   ├── product.go
│   │   │   │   ├── order.go
│   │   │   │   └── base_entity.go  # Entity base con campos comunes
│   │   │   ├── valueobjects/
│   │   │   │   ├── email.go
│   │   │   │   ├── money.go
│   │   │   │   ├── phone.go
│   │   │   │   └── address.go
│   │   │   ├── aggregates/
│   │   │   │   ├── user_aggregate.go
│   │   │   │   └── order_aggregate.go
│   │   │   ├── events/
│   │   │   │   ├── domain_event.go
│   │   │   │   ├── user_events.go
│   │   │   │   └── order_events.go
│   │   │   └── errors/
│   │   │       ├── domain_errors.go
│   │   │       └── error_codes.go
│   │   ├── ports/                  # PUERTOS (Interfaces - Dependency Inversion)
│   │   │   ├── repositories/
│   │   │   │   ├── user_repository.go
│   │   │   │   ├── product_repository.go
│   │   │   │   ├── order_repository.go
│   │   │   │   └── unit_of_work.go
│   │   │   ├── services/
│   │   │   │   ├── email_service.go
│   │   │   │   ├── payment_service.go
│   │   │   │   ├── notification_service.go
│   │   │   │   └── cache_service.go
│   │   │   ├── handlers/
│   │   │   │   ├── user_handler.go
│   │   │   │   ├── product_handler.go
│   │   │   │   └── order_handler.go
│   │   │   └── publishers/
│   │   │       └── event_publisher.go
│   │   └── usecases/               # CASOS DE USO (Application Services)
│   │       ├── user/
│   │       │   ├── create_user.go
│   │       │   ├── update_user.go
│   │       │   ├── get_user.go
│   │       │   └── delete_user.go
│   │       ├── product/
│   │       │   ├── create_product.go
│   │       │   ├── update_product.go
│   │       │   ├── list_products.go
│   │       │   └── delete_product.go
│   │       ├── order/
│   │       │   ├── create_order.go
│   │       │   ├── process_order.go
│   │       │   └── cancel_order.go
│   │       └── auth/
│   │           ├── login.go
│   │           ├── register.go
│   │           └── refresh_token.go
│   └── adapters/                   # ADAPTADORES (Implementaciones)
│       ├── primary/                # Adaptadores Primarios (Driving)
│       │   ├── http/
│       │   │   └── fiber/
│       │   │       ├── server.go
│       │   │       ├── config/
│       │   │       │   └── fiber_config.go
│       │   │       ├── middleware/
│       │   │       │   ├── auth.go
│       │   │       │   ├── cors.go
│       │   │       │   ├── logger.go
│       │   │       │   ├── rate_limit.go
│       │   │       │   ├── recovery.go
│       │   │       │   └── validator.go
│       │   │       ├── handlers/
│       │   │       │   ├── user_handler.go
│       │   │       │   ├── product_handler.go
│       │   │       │   ├── order_handler.go
│       │   │       │   ├── auth_handler.go
│       │   │       │   └── health_handler.go
│       │   │       ├── routes/
│       │   │       │   ├── user_routes.go
│       │   │       │   ├── product_routes.go
│       │   │       │   ├── order_routes.go
│       │   │       │   ├── auth_routes.go
│       │   │       │   └── routes.go
│       │   │       ├── dto/
│       │   │       │   ├── request/
│       │   │       │   │   ├── user_request.go
│       │   │       │   │   ├── product_request.go
│       │   │       │   │   └── auth_request.go
│       │   │       │   ├── response/
│       │   │       │   │   ├── user_response.go
│       │   │       │   │   ├── product_response.go
│       │   │       │   │   ├── auth_response.go
│       │   │       │   │   └── base_response.go
│       │   │       │   └── mappers/
│       │   │       │       ├── user_mapper.go
│       │   │       │       └── product_mapper.go
│       │   │       └── websocket/
│       │   │           ├── hub.go
│       │   │           ├── client.go
│       │   │           └── handler.go
│       │   └── grpc/
│       │       ├── server.go
│       │       ├── config/
│       │       │   └── grpc_config.go
│       │       ├── interceptors/
│       │       │   ├── auth.go
│       │       │   ├── logging.go
│       │       │   ├── recovery.go
│       │       │   └── validation.go
│       │       ├── handlers/
│       │       │   ├── user_handler.go
│       │       │   ├── product_handler.go
│       │       │   └── auth_handler.go
│       │       └── pb/              # Generated protobuf files
│       │           ├── user/
│       │           │   ├── user.pb.go
│       │           │   └── user_grpc.pb.go
│       │           ├── product/
│       │           │   ├── product.pb.go
│       │           │   └── product_grpc.pb.go
│       │           └── common/
│       │               ├── common.pb.go
│       │               └── health.pb.go
│       └── secondary/              # Adaptadores Secundarios (Driven)
│           ├── persistence/
│           │   ├── gorm/
│           │   │   ├── connection.go
│           │   │   ├── transaction.go
│           │   │   ├── hooks/
│           │   │   │   ├── base_hooks.go
│           │   │   │   ├── audit_hooks.go
│           │   │   │   └── soft_delete_hooks.go
│           │   │   ├── migrations/
│           │   │   │   ├── auto_migrate.go
│           │   │   │   ├── manual_migrations.go
│           │   │   │   └── seeders/
│           │   │   │       ├── user_seeder.go
│           │   │   │       └── product_seeder.go
│           │   │   ├── repositories/
│           │   │   │   ├── base_repository.go
│           │   │   │   ├── user_repository.go
│           │   │   │   ├── product_repository.go
│           │   │   │   ├── order_repository.go
│           │   │   │   └── unit_of_work.go
│           │   │   ├── models/
│           │   │   │   ├── base_model.go
│           │   │   │   ├── user_model.go
│           │   │   │   ├── product_model.go
│           │   │   │   ├── order_model.go
│           │   │   │   └── associations.go
│           │   │   ├── scopes/
│           │   │   │   ├── common_scopes.go
│           │   │   │   ├── user_scopes.go
│           │   │   │   └── product_scopes.go
│           │   │   └── plugins/
│           │   │       ├── audit_plugin.go
│           │   │       └── metrics_plugin.go
│           │   ├── redis/
│           │   │   ├── connection.go
│           │   │   ├── cache_service.go
│           │   │   └── session_store.go
│           │   └── memory/          # Para testing
│           │       └── repositories/
│           │           ├── user_repository.go
│           │           └── product_repository.go
│           ├── external/
│           │   ├── email/
│           │   │   ├── smtp_service.go
│           │   │   ├── sendgrid_service.go
│           │   │   └── ses_service.go
│           │   ├── payment/
│           │   │   ├── stripe_service.go
│           │   │   ├── paypal_service.go
│           │   │   └── mercadopago_service.go
│           │   ├── storage/
│           │   │   ├── s3_service.go
│           │   │   ├── gcs_service.go
│           │   │   └── local_service.go
│           │   └── notifications/
│           │       ├── fcm_service.go
│           │       └── push_service.go
│           └── messaging/
│               ├── rabbitmq/
│               │   ├── connection.go
│               │   ├── publisher.go
│               │   └── consumer.go
│               ├── kafka/
│               │   ├── producer.go
│               │   └── consumer.go
│               └── nats/
│                   ├── publisher.go
│                   └── subscriber.go
├── pkg/                            # UTILIDADES COMPARTIDAS
│   ├── config/
│   │   ├── config.go
│   │   └── env.go
│   ├── logger/
│   │   ├── logger.go
│   │   ├── logrus.go
│   │   └── fields.go
│   ├── validator/
│   │   ├── validator.go
│   │   └── custom_rules.go
│   ├── jwt/
│   │   ├── jwt.go
│   │   ├── claims.go
│   │   └── middleware.go
│   ├── security/
│   │   ├── crypto.go
│   │   ├── hash.go
│   │   └── rsa.go
│   ├── utils/
│   │   ├── time.go
│   │   ├── strings.go
│   │   ├── pagination.go
│   │   └── response.go
│   ├── errors/
│   │   ├── app_error.go
│   │   ├── error_handler.go
│   │   └── error_codes.go
│   └── metrics/
│       ├── prometheus.go
│       └── middleware.go
├── proto/                          # PROTOCOL BUFFERS DEFINITIONS
│   ├── user/
│   │   └── user.proto
│   ├── product/
│   │   └── product.proto
│   ├── order/
│   │   └── order.proto
│   └── common/
│       ├── common.proto
│       └── health.proto
├── tests/                          # TESTS
│   ├── unit/
│   │   ├── core/
│   │   │   ├── domain/
│   │   │   │   ├── entities/
│   │   │   │   └── valueobjects/
│   │   │   └── usecases/
│   │   │       ├── user/
│   │   │       └── product/
│   │   ├── adapters/
│   │   │   ├── http/
│   │   │   ├── grpc/
│   │   │   └── persistence/
│   │   └── pkg/
│   ├── integration/
│   │   ├── database/
│   │   ├── api/
│   │   └── grpc/
│   ├── e2e/
│   │   ├── scenarios/
│   │   └── fixtures/
│   ├── mocks/
│   │   ├── repositories/
│   │   └── services/
│   └── testdata/
│       ├── fixtures/
│       └── seeds/
├── scripts/                        # SCRIPTS
│   ├── build/
│   │   ├── build.sh
│   │   └── build-docker.sh
│   ├── database/
│   │   ├── migrate.sh
│   │   ├── seed.sh
│   │   ├── reset.sh
│   │   ├── backup.sh
│   │   └── gorm-gen.sh             # GORM code generation
│   ├── proto/
│   │   ├── generate.sh
│   │   └── validate.sh
│   ├── development/
│   │   ├── setup.sh
│   │   └── run-local.sh
│   └── deployment/
│       ├── deploy.sh
│       └── rollback.sh
├── deployments/                    # DEPLOYMENT
│   ├── docker/
│   │   ├── Dockerfile
│   │   ├── Dockerfile.dev
│   │   ├── docker-compose.yml
│   │   ├── docker-compose.dev.yml
│   │   └── docker-compose.test.yml
│   ├── k8s/
│   │   ├── namespace.yaml
│   │   ├── configmap.yaml
│   │   ├── secret.yaml
│   │   ├── deployment.yaml
│   │   ├── service.yaml
│   │   ├── ingress.yaml
│   │   └── hpa.yaml
│   ├── helm/
│   │   ├── Chart.yaml
│   │   ├── values.yaml
│   │   └── templates/
│   └── terraform/
│       ├── main.tf
│       ├── variables.tf
│       └── outputs.tf
├── configs/                        # CONFIGURATION FILES
│   ├── config.yaml
│   ├── config.dev.yaml
│   ├── config.prod.yaml
│   └── config.test.yaml
├── .env.example
├── .env.dev
├── .env.test
├── .gitignore
├── .golangci.yml
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

## ARQUITECTURA HEXAGONAL MEJORADA

### CORE (Centro del Hexágono)

```
├── domain/        # Entidades, Value Objects, Agregados y Eventos
├── ports/         # Interfaces (Dependency Inversion Principle)
└── usecases/      # Casos de uso específicos (Single Responsibility)
```

### ADAPTADORES

```
├── primary/       # Adaptadores de entrada (HTTP, gRPC)
└── secondary/     # Adaptadores de salida (DB, APIs externas)
```

## PRINCIPIOS SOLID APLICADOS

### 1. Single Responsibility Principle (SRP)

- **Casos de uso separados**: Cada usecase tiene una sola responsabilidad
- **Handlers específicos**: Un handler por entidad/funcionalidad
- **Separación de concerns**: HTTP/gRPC/DB en adaptadores diferentes

### 2. Open/Closed Principle (OCP)

- **Interfaces en ports/**: Fácil extensión sin modificación
- **Múltiples implementaciones**: SMTP/SendGrid para email
- **Adaptadores intercambiables**: Postgres/MySQL/MongoDB

### 3. Liskov Substitution Principle (LSP)

- **Implementaciones de interfaces**: Todas cumplen el contrato
- **Polimorfismo**: Servicios externos intercambiables

### 4. Interface Segregation Principle (ISP)

- **Interfaces específicas**: UserRepository, ProductRepository separados
- **Puertos granulares**: EmailService, PaymentService independientes

### 5. Dependency Inversion Principle (DIP)

- **Ports como abstracciones**: Core depende de interfaces
- **Inyección de dependencias**: Adaptadores implementan interfaces

## FLUJO DE DATOS

### HTTP Request Flow

```
HTTP Request → Fiber Middleware → Fiber Handler → UseCase → Domain → GORM Repository → PostgreSQL
```

### gRPC Request Flow

```
gRPC Request → gRPC Interceptor → gRPC Handler → UseCase → Domain → GORM Repository → PostgreSQL
```

### GORM Transaction Flow

```
UseCase → Unit of Work → GORM Transaction → Multiple Repositories → Commit/Rollback
```

### Event-Driven Flow

```
Domain Event → Event Publisher → Message Queue → Event Handler → UseCase
```

## TECNOLOGÍAS Y LIBRERÍAS

### Core Stack

- **Framework HTTP**: Fiber v2
- **gRPC**: Google gRPC
- **ORM**: GORM v2 + PostgreSQL Driver
- **Database**: PostgreSQL
- **Cache**: Redis
- **Logging**: Logrus
- **Validation**: go-playground/validator
- **Migration**: GORM AutoMigrate + Manual Migrations

### Messaging & Events

- **Message Brokers**: RabbitMQ, Kafka, NATS
- **Event Sourcing**: Preparado para implementar

### Monitoring & Observability

- **Metrics**: Prometheus
- **Tracing**: Preparado para Jaeger/Zipkin
- **Health Checks**: Implementado

### Security

- **JWT**: Token-based authentication
- **Crypto**: Bcrypt, RSA
- **Rate Limiting**: Configurado

## PREPARACIÓN PARA FUTURO

### gRPC-Gateway

- **Gateway server**: Punto de entrada preparado
- **Proto definitions**: Organizadas por dominio
- **REST to gRPC**: Mapping automático

### WebSockets

- **Hub pattern**: Manejo de conexiones
- **Client management**: Gestión de clientes
- **Real-time events**: Integración con domain events

### Event Sourcing

- **Domain events**: Base implementada
- **Event store**: Preparado para implementar
- **CQRS**: Separación comando/consulta lista

## BENEFICIOS DE ESTA ESTRUCTURA

### **Mantenibilidad**

- Separación clara de responsabilidades
- Código organizado por dominio
- Fácil localización de funcionalidades

### **Testabilidad**

- Mocks organizados por tipo
- Tests unitarios, integración y E2E
- Fixtures y test data separados

### **Escalabilidad**

- Microservicios ready
- Horizontal scaling preparado
- Load balancing support

### **Flexibilidad**

- Múltiples adaptadores
- Tecnologías intercambiables
- Deployment options variadas

### **Observabilidad**

- Logging estructurado con Logrus
- Métricas con Prometheus
- Health checks implementados

### **Seguridad**

- Autenticación JWT
- Autorización por roles
- Rate limiting
- Input validation

### GORM Integration Benefits

### **ORM Advantages**

- **Code Generation**: GORM Gen para queries type-safe
- **Auto Migrations**: Esquemas sincronizados automáticamente
- **Rich Associations**: Relaciones complejas simplificadas
- **Hooks System**: Callbacks para auditoría y validación
- **Soft Delete**: Eliminación lógica built-in
- **Scopes**: Queries reutilizables y composables

### **Performance Features**

- **Eager/Lazy Loading**: Control de carga de relaciones
- **Batch Operations**: Inserts/Updates masivos
- **Raw SQL Support**: Queries complejas cuando sea necesario
- **Connection Pooling**: Gestión automática de conexiones
- **Prepared Statements**: Queries optimizadas

### **Developer Experience**

- **Auto Migration**: Sincronización automática de esquemas
- **Model Tagging**: Configuración declarativa
- **Plugin System**: Extensibilidad avanzada
- **Debug Mode**: SQL logging para desarrollo
- **Error Handling**: Manejo robusto de errores DB

Esta estructura está optimizada para proyectos Go modernos, siguiendo las mejores prácticas de la comunidad y preparada para escalar desde un monolito hasta microservicios cuando sea necesario.
