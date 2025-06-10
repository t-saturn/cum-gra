```txt
server/
├── cmd/
│   └── api/
│       └── main.go                 # Punto de entrada de la aplicación
├── internal/
│   ├── core/                       # DOMINIO (Centro del hexágono)
│   │   ├── domain/
│   │   │   ├── entities/
│   │   │   │   ├── user.go
│   │   │   │   └── product.go
│   │   │   ├── valueobjects/
│   │   │   │   ├── email.go
│   │   │   │   └── money.go
│   │   │   └── errors/
│   │   │       └── domain_errors.go
│   │   ├── ports/                  # PUERTOS (Interfaces)
│   │   │   ├── repositories/
│   │   │   │   ├── user_repository.go
│   │   │   │   └── product_repository.go
│   │   │   ├── services/
│   │   │   │   ├── email_service.go
│   │   │   │   └── payment_service.go
│   │   │   └── handlers/
│   │   │       ├── user_handler.go
│   │   │       └── product_handler.go
│   │   └── services/               # CASOS DE USO (Application Services)
│   │       ├── user_service.go
│   │       ├── product_service.go
│   │       └── auth_service.go
│   └── adapters/                   # ADAPTADORES (Implementaciones)
│       ├── primary/                # Adaptadores Primarios (Driving)
│       │   ├── http/
│       │   │   ├── fiber/
│       │   │   │   ├── server.go
│       │   │   │   ├── middleware/
│       │   │   │   │   ├── auth.go
│       │   │   │   │   ├── cors.go
│       │   │   │   │   └── logger.go
│       │   │   │   ├── handlers/
│       │   │   │   │   ├── user_handler.go
│       │   │   │   │   ├── product_handler.go
│       │   │   │   │   └── health_handler.go
│       │   │   │   ├── routes/
│       │   │   │   │   ├── user_routes.go
│       │   │   │   │   ├── product_routes.go
│       │   │   │   │   └── routes.go
│       │   │   │   └── dto/
│       │   │   │       ├── user_dto.go
│       │   │   │       └── product_dto.go
│       │   │   └── grpc/           # Si usas gRPC
│       │   │       └── server.go
│       │   └── cli/                # Si tienes comandos CLI
│       │       └── commands.go
│       └── secondary/              # Adaptadores Secundarios (Driven)
│           ├── persistence/
│           │   ├── postgres/
│           │   │   ├── connection.go
│           │   │   ├── migrations/
│           │   │   ├── repositories/
│           │   │   │   ├── user_repository.go
│           │   │   │   └── product_repository.go
│           │   │   └── models/
│           │   │       ├── user_model.go
│           │   │       └── product_model.go
│           │   ├── redis/
│           │   │   ├── connection.go
│           │   │   └── cache_service.go
│           │   └── memory/          # Para testing
│           │       └── repositories/
│           ├── external/
│           │   ├── email/
│           │   │   ├── smtp_service.go
│           │   │   └── sendgrid_service.go
│           │   ├── payment/
│           │   │   ├── stripe_service.go
│           │   │   └── paypal_service.go
│           │   └── storage/
│           │       ├── s3_service.go
│           │       └── local_service.go
│           └── messaging/
│               ├── rabbitmq/
│               │   └── publisher.go
│               └── kafka/
│                   └── producer.go
├── pkg/                            # UTILIDADES COMPARTIDAS
│   ├── config/
│   │   └── config.go
│   ├── logger/
│   │   └── logger.go
│   ├── validator/
│   │   └── validator.go
│   ├── jwt/
│   │   └── jwt.go
│   └── utils/
│       ├── crypto.go
│       └── time.go
├── tests/                          # TESTS
│   ├── unit/
│   ├── integration/
│   └── e2e/
├── docs/                           # DOCUMENTACIÓN
│   ├── api/
│   │   └── swagger.yaml
│   └── architecture/
├── scripts/                        # SCRIPTS
│   ├── migrate.sh
│   └── seed.sh
├── deployments/                    # DEPLOYMENT
│   ├── docker/
│   │   └── Dockerfile
│   └── k8s/
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

# EXPLICACIÓN DE LA ARQUITECTURA

CORE (Centro del Hexágono):
├── domain/ # Entidades y reglas de negocio puras
├── ports/ # Interfaces que definen contratos
└── services/ # Casos de uso y lógica de aplicación

ADAPTERS:
├── primary/ # Reciben peticiones (HTTP, CLI, gRPC)
└── secondary/ # Implementan servicios externos (DB, APIs)

FLUJO DE DATOS:
HTTP Request → Fiber Handler → Application Service → Domain → Repository → Database
BENEFICIOS:
✅ Independiente de frameworks
✅ Testeable (mock de interfaces)
✅ Flexible (cambiar adaptadores sin afectar el core)
✅ Mantenible (separación clara de responsabilidades)
✅ Escalable (agregar nuevos adaptadores fácilmente)
