# Estructura Go + Fiber + PostgreSQL + gRPC Simplificada

```txt
server/
├── logs/              # Aquí se guardarán los logs diarios
├── cmd/
│   └── grpc/
│       └── main.go                 # gRPC server
├── internal/
│   ├── models/                     # Modelos de datos (GORM)
│   │   ├── base.go
│   │   ├── user.go
│   │   ├── product.go
│   │   └── order.go
│   ├── repositories/               # Acceso a datos
│   │   ├── interfaces.go           # Interfaces de repositorios
│   │   ├── user_repository.go
│   │   ├── product_repository.go
│   │   └── order_repository.go
│   ├── services/                   # Lógica de negocio
│   │   ├── user_service.go
│   │   ├── product_service.go
│   │   ├── order_service.go
│   │   └── auth_service.go
│   ├── handlers/                   # HTTP handlers (Fiber)
│   │   ├── user_handler.go
│   │   ├── product_handler.go
│   │   ├── order_handler.go
│   │   ├── auth_handler.go
│   │   └── health_handler.go
│   ├── grpc/                       # gRPC handlers
│   │   ├── user_handler.go
│   │   ├── product_handler.go
│   │   └── auth_handler.go
│   ├── middleware/                 # HTTP middleware
│   │   ├── auth.go
│   │   ├── cors.go
│   │   ├── logger.go
│   │   ├── rate_limit.go
│   │   └── validator.go
│   ├── routes/                     # Rutas HTTP
│   │   ├── api.go
│   │   ├── user_routes.go
│   │   ├── product_routes.go
│   │   └── auth_routes.go
│   ├── dto/                        # Data Transfer Objects
│   │   ├── user_dto.go
│   │   ├── product_dto.go
│   │   ├── order_dto.go
│   │   └── auth_dto.go
│   └── external/                   # Servicios externos
│       ├── email/
│       │   └── email_service.go
│       ├── payment/
│       │   └── payment_service.go
│       └── storage/
│           └── file_service.go
├── pkg/                            # Utilidades compartidas
│   ├── config/
│   │   ├── config.go
│   │   └── database.go
│   ├── database/
│   │   ├── connection.go
│   │   ├── migration.go
│   │   └── seeder.go
│   ├── logger/
│   │   └── logger.go
│   ├── validator/
│   │   └── validator.go
│   ├── jwt/
│   │   └── jwt.go
│   ├── utils/
│   │   ├── response.go
│   │   ├── pagination.go
│   │   └── helpers.go
│   └── errors/
│       └── errors.go
├── proto/                          # Protocol Buffers
│   ├── user/
│   │   └── user.proto
│   ├── product/
│   │   └── product.proto
│   └── common/
│       └── common.proto
├── pb/                             # Generated protobuf files
│   ├── user/
│   │   ├── user.pb.go
│   │   └── user_grpc.pb.go
│   └── product/
│       ├── product.pb.go
│       └── product_grpc.pb.go
├── migrations/                     # Migraciones de base de datos
│   ├── 001_create_users_table.sql
│   ├── 002_create_products_table.sql
│   └── 003_create_orders_table.sql
├── tests/                          # Tests
│   ├── unit/
│   │   ├── services/
│   │   └── repositories/
│   ├── integration/
│   │   └── api/
│   └── mocks/
├── scripts/                        # Scripts utilitarios
│   ├── build.sh
│   ├── migrate.sh
│   ├── proto-gen.sh
│   └── test.sh
├── docker/
│   ├── Dockerfile
│   ├── docker-compose.yml
│   └── docker-compose.dev.yml
├── configs/
│   ├── config.yaml
│   ├── config.dev.yaml
│   └── config.prod.yaml
├── .env.example
├── .gitignore
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

## EXPLICACIÓN DE LA ESTRUCTURA

### **1. `/cmd`** - Puntos de entrada

- **api/main.go**: Servidor HTTP con Fiber
- **grpc/main.go**: Servidor gRPC

### **2. `/internal`** - Código específico de la aplicación

#### **`/models`** - Modelos GORM

```go
// internal/models/user.go
type User struct {
    Base
    Name     string `json:"name" gorm:"not null"`
    Email    string `json:"email" gorm:"unique;not null"`
    Password string `json:"-" gorm:"not null"`
}
```

#### **`/repositories`** - Acceso a datos

```go
// internal/repositories/interfaces.go
type UserRepository interface {
    Create(user *models.User) error
    GetByID(id uint) (*models.User, error)
    GetByEmail(email string) (*models.User, error)
    Update(user *models.User) error
    Delete(id uint) error
}
```

#### **`/services`** - Lógica de negocio

```go
// internal/services/user_service.go
type UserService struct {
    userRepo repositories.UserRepository
}

func (s *UserService) CreateUser(dto *dto.CreateUserDTO) (*models.User, error) {
    // Validaciones de negocio
    // Lógica de creación
}
```

#### **`/handlers`** - Controladores HTTP

```go
// internal/handlers/user_handler.go
type UserHandler struct {
    userService *services.UserService
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    // Manejo de request/response HTTP
}
```

#### **`/grpc`** - Handlers gRPC

```go
// internal/grpc/user_handler.go
type UserGRPCHandler struct {
    userService *services.UserService
}

func (h *UserGRPCHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
    // Manejo de requests gRPC
}
```

### **3. `/pkg`** - Código reutilizable

#### **`/config`** - Configuración

```go
// pkg/config/config.go
type Config struct {
    Server   ServerConfig   `yaml:"server"`
    Database DatabaseConfig `yaml:"database"`
    JWT      JWTConfig      `yaml:"jwt"`
}
```

#### **`/database`** - Conexión y migraciones

```go
// pkg/database/connection.go
func NewConnection(cfg DatabaseConfig) (*gorm.DB, error) {
    // Configuración de GORM
}
```

## **FLUJO DE DATOS SIMPLIFICADO**

### HTTP Request:

```
Request → Middleware → Handler → Service → Repository → Database
```

### gRPC Request:

```
gRPC Request → gRPC Handler → Service → Repository → Database
```

## **VENTAJAS DE ESTA ESTRUCTURA**

### **Simplicidad**

- Menos capas y abstracciones
- Fácil de entender y mantener
- Rápido desarrollo inicial

### **Separación de responsabilidades**

- **Models**: Estructura de datos
- **Repositories**: Acceso a datos
- **Services**: Lógica de negocio
- **Handlers**: Interfaz HTTP/gRPC

### **Testeable**

- Interfaces para mocking
- Tests unitarios por capa
- Tests de integración separados

### **Escalable**

- Fácil agregar nuevas funcionalidades
- Servicios desacoplados
- Preparado para microservicios

## **EJEMPLO DE IMPLEMENTACIÓN**

### **Model (GORM)**

```go
// internal/models/user.go
type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name" gorm:"not null"`
    Email     string    `json:"email" gorm:"unique;not null"`
    Password  string    `json:"-" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### **Repository**

```go
// internal/repositories/user_repository.go
type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
    return r.db.Create(user).Error
}
```

### **Service**

```go
// internal/services/user_service.go
type UserService struct {
    userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(dto *dto.CreateUserDTO) (*models.User, error) {
    // Hash password
    // Validate business rules
    // Create user
}
```

### **Handler**

```go
// internal/handlers/user_handler.go
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    var dto dto.CreateUserDTO
    if err := c.BodyParser(&dto); err != nil {
        return c.Status(400).JSON(utils.ErrorResponse("Invalid input"))
    }

    user, err := h.userService.CreateUser(&dto)
    if err != nil {
        return c.Status(500).JSON(utils.ErrorResponse(err.Error()))
    }

    return c.Status(201).JSON(utils.SuccessResponse(user))
}
```

## **TECNOLOGÍAS INCLUIDAS**

- **HTTP Framework**: Fiber v2
- **gRPC**: Google gRPC
- **ORM**: GORM v2
- **Database**: PostgreSQL
- **Cache**: Redis (opcional)
- **Auth**: JWT
- **Validation**: go-playground/validator
- **Logging**: Structured logging
- **Testing**: Testify + mocks
