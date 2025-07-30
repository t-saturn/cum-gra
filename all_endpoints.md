# Documentación API - Sistema SSO

## Endpoints de Autenticación

### 1. Verificar Credenciales

**Ruta:** `/auth/verify`  
**Método:** `POST`  
**Descripción:** Verifica las credenciales del usuario sin crear sesión ni tokens

**Datos de entrada:**

```json
{
  "email": "usuario@ejemplo.com",
  "password": "contraseña123",
  "application_id": "app_001",
  "device_info": {
    "user_agent": "Mozilla/5.0...",
    "ip": "192.168.1.100",
    "device_id": "device_123",
    "browser_name": "Chrome",
    "browser_version": "91.0",
    "os": "Windows",
    "os_version": "10",
    "device_type": "desktop",
    "timezone": "America/Lima",
    "language": "es-PE"
  }
}
```

**Datos de respuesta:**

```json
{
  "success": true,
  "message": "Credenciales válidas",
  "data": {
    "attempt_id": "507f1f77bcf86cd799439011",
    "user_id": "user_123",
    "status": "success",
    "validated_at": "2025-07-30T10:30:00Z",
    "validation_response": {
      "user_id": "user_123",
      "service_response": "valid_user",
      "validated_by": "auth_service",
      "validation_time": 150
    }
  }
}
```

### 2. Login Completo

**Ruta:** `/auth/login`  
**Método:** `POST`  
**Descripción:** Realiza login completo creando sesión y tokens

**Datos de entrada:**

```json
{
  "email": "usuario@ejemplo.com",
  "password": "contraseña123",
  "application_id": "app_001",
  "remember_me": true,
  "device_info": {
    "user_agent": "Mozilla/5.0...",
    "ip": "192.168.1.100",
    "device_id": "device_123",
    "browser_name": "Chrome",
    "browser_version": "91.0",
    "os": "Windows",
    "os_version": "10",
    "device_type": "desktop",
    "timezone": "America/Lima",
    "language": "es-PE",
    "location": {
      "country": "Peru",
      "country_code": "PE",
      "region": "Lima",
      "city": "Lima",
      "coordinates": [-77.0428, -12.0464],
      "isp": "Telefonica del Peru",
      "organization": "Movistar"
    }
  }
}
```

**Datos de respuesta:**

```json
{
  "success": true,
  "message": "Login exitoso",
  "data": {
    "user_id": "user_123",
    "session": {
      "session_id": "550e8400-e29b-41d4-a716-446655440000",
      "status": "active",
      "expires_at": "2025-07-31T10:30:00Z",
      "created_at": "2025-07-30T10:30:00Z"
    },
    "tokens": {
      "access_token": {
        "token_id": "550e8400-e29b-41d4-a716-446655440001",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "token_type": "access",
        "expires_at": "2025-07-30T11:30:00Z"
      },
      "refresh_token": {
        "token_id": "550e8400-e29b-41d4-a716-446655440002",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "token_type": "refresh",
        "expires_at": "2025-08-06T10:30:00Z"
      }
    },
    "attempt_id": "507f1f77bcf86cd799439012"
  }
}
```

### 3. Cerrar Sesión

**Ruta:** `/auth/logout`  
**Método:** `POST`  
**Descripción:** Cierra la sesión actual y revoca los tokens

**Datos de entrada:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "session_id": "550e8400-e29b-41d4-a716-446655440000",
  "logout_type": "user_initiated"
}
```

**Datos de respuesta:**

```json
{
  "success": true,
  "message": "Logout exitoso",
  "data": {
    "session_id": "550e8400-e29b-41d4-a716-446655440000",
    "revoked_at": "2025-07-30T12:30:00Z",
    "tokens_revoked": [
      "550e8400-e29b-41d4-a716-446655440001",
      "550e8400-e29b-41d4-a716-446655440002"
    ]
  }
}
```

## Endpoints de Tokens

### 4. Refrescar Token

**Ruta:** `/auth/token/refresh`  
**Método:** `POST`  
**Descripción:** Genera nuevos tokens usando el refresh token

**Datos de entrada:**

```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "device_info": {
    "user_agent": "Mozilla/5.0...",
    "ip": "192.168.1.100",
    "device_id": "device_123"
  }
}
```

**Datos de respuesta:**

```json
{
  "success": true,
  "message": "Token refrescado exitosamente",
  "data": {
    "access_token": {
      "token_id": "550e8400-e29b-41d4-a716-446655440003",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "token_type": "access",
      "expires_at": "2025-07-30T13:30:00Z"
    },
    "refresh_token": {
      "token_id": "550e8400-e29b-41d4-a716-446655440004",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
      "token_type": "refresh",
      "expires_at": "2025-08-06T12:30:00Z"
    },
    "session_id": "550e8400-e29b-41d4-a716-446655440000",
    "refresh_count": 1
  }
}
```

### 5. Validar Token

**Ruta:** `/auth/token/validate`  
**Método:** `POST`  
**Descripción:** Valida si un access token es válido y activo

**Datos de entrada:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "application_id": "app_001"
}
```

**Datos de respuesta:**

```json
{
  "success": true,
  "message": "Token válido",
  "data": {
    "token_id": "550e8400-e29b-41d4-a716-446655440001",
    "user_id": "user_123",
    "session_id": "550e8400-e29b-41d4-a716-446655440000",
    "status": "active",
    "token_type": "access",
    "expires_at": "2025-07-30T11:30:00Z",
    "last_used": "2025-07-30T10:45:00Z",
    "is_valid": true
  }
}
```

## Endpoints de Sesiones

### 6. Ver Sesión Actual

**Ruta:** `/auth/session/me`  
**Método:** `GET`  
**Descripción:** Obtiene información de la sesión actual del usuario

**Headers:**

```js
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Datos de respuesta:**

```json
{
  "success": true,
  "message": "Sesión actual obtenida",
  "data": {
    "session_id": "550e8400-e29b-41d4-a716-446655440000",
    "user_id": "user_123",
    "status": "active",
    "is_active": true,
    "created_at": "2025-07-30T10:30:00Z",
    "last_activity": "2025-07-30T12:15:00Z",
    "expires_at": "2025-07-31T10:30:00Z",
    "device_info": {
      "browser_name": "Chrome",
      "browser_version": "91.0",
      "os": "Windows",
      "os_version": "10",
      "device_type": "desktop",
      "ip": "192.168.1.100",
      "location": {
        "country": "Peru",
        "city": "Lima"
      }
    },
    "active_tokens": [
      "550e8400-e29b-41d4-a716-446655440001",
      "550e8400-e29b-41d4-a716-446655440002"
    ]
  }
}
```

### 7. Listar Sesiones

**Ruta:** `/auth/sessions`  
**Método:** `GET`  
**Descripción:** Lista todas las sesiones activas del usuario

**Headers:**

```js
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Query Parameters:**

- `status` (opcional): Filtrar por estado (`active`, `inactive`, `revoked`, `expired`)
- `limit` (opcional): Número máximo de resultados (default: 10)
- `offset` (opcional): Desplazamiento para paginación (default: 0)

**Datos de respuesta:**

```json
{
  "success": true,
  "message": "Sesiones obtenidas",
  "data": {
    "sessions": [
      {
        "session_id": "550e8400-e29b-41d4-a716-446655440000",
        "status": "active",
        "is_active": true,
        "created_at": "2025-07-30T10:30:00Z",
        "last_activity": "2025-07-30T12:15:00Z",
        "expires_at": "2025-07-31T10:30:00Z",
        "device_info": {
          "browser_name": "Chrome",
          "os": "Windows",
          "device_type": "desktop",
          "location": {
            "country": "Peru",
            "city": "Lima"
          }
        },
        "is_current": true
      },
      {
        "session_id": "550e8400-e29b-41d4-a716-446655440005",
        "status": "active",
        "is_active": true,
        "created_at": "2025-07-29T08:00:00Z",
        "last_activity": "2025-07-29T18:30:00Z",
        "expires_at": "2025-07-30T08:00:00Z",
        "device_info": {
          "browser_name": "Safari",
          "os": "iOS",
          "device_type": "mobile",
          "location": {
            "country": "Peru",
            "city": "Lima"
          }
        },
        "is_current": false
      }
    ],
    "total": 2,
    "limit": 10,
    "offset": 0
  }
}
```

### 8. Revocar Sesión Específica

**Ruta:** `/auth/sessions/{session_id}`  
**Método:** `DELETE`  
**Descripción:** Revoca una sesión específica y todos sus tokens

**Headers:**

```js
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Path Parameters:**

- `session_id`: ID de la sesión a revocar

**Datos de entrada:**

```json
{
  "reason": "user_requested",
  "revoke_all_tokens": true
}
```

**Datos de respuesta:**

```json
{
  "success": true,
  "message": "Sesión revocada exitosamente",
  "data": {
    "session_id": "550e8400-e29b-41d4-a716-446655440005",
    "revoked_at": "2025-07-30T12:30:00Z",
    "revocation_reason": "user_requested",
    "tokens_revoked": [
      "550e8400-e29b-41d4-a716-446655440006",
      "550e8400-e29b-41d4-a716-446655440007"
    ]
  }
}
```

## Endpoint de Salud

### 9. Estado del Servicio

**Ruta:** `/auth/health`  
**Método:** `GET`  
**Descripción:** Verifica el estado de salud del servicio de autenticación

**Datos de respuesta:**

```json
{
  "success": true,
  "message": "Servicio operativo",
  "data": {
    "status": "healthy",
    "timestamp": "2025-07-30T12:30:00Z",
    "version": "1.0.0",
    "uptime": "24h 15m 30s",
    "databases": {
      "mongodb": {
        "status": "connected",
        "response_time": "15ms"
      },
      "postgresql": {
        "status": "connected",
        "response_time": "8ms"
      }
    },
    "dependencies": {
      "captcha_service": "healthy",
      "notification_service": "healthy"
    }
  }
}
```

## Códigos de Error Comunes

### Respuesta de Error

```json
{
  "success": false,
  "error": {
    "code": "INVALID_CREDENTIALS",
    "message": "Las credenciales proporcionadas no son válidas",
    "details": {
      "field": "password",
      "reason": "incorrect_password"
    }
  },
  "timestamp": "2025-07-30T12:30:00Z"
}
```

### Principales Códigos de Error

| Código                  | Descripción                  | HTTP Status |
| ----------------------- | ---------------------------- | ----------- |
| `INVALID_CREDENTIALS`   | Credenciales incorrectas     | 401         |
| `TOKEN_EXPIRED`         | Token expirado               | 401         |
| `TOKEN_INVALID`         | Token inválido o malformado  | 401         |
| `SESSION_EXPIRED`       | Sesión expirada              | 401         |
| `SESSION_NOT_FOUND`     | Sesión no encontrada         | 404         |
| `USER_NOT_FOUND`        | Usuario no encontrado        | 404         |
| `VALIDATION_ERROR`      | Error de validación de datos | 400         |
| `RATE_LIMIT_EXCEEDED`   | Límite de intentos excedido  | 429         |
| `INTERNAL_SERVER_ERROR` | Error interno del servidor   | 500         |
| `SERVICE_UNAVAILABLE`   | Servicio no disponible       | 503         |

## Notas Importantes

1. **Autenticación**: Los endpoints marcados requieren el header `Authorization: Bearer {access_token}`
2. **Rate Limiting**: Todos los endpoints tienen límites de velocidad configurados
3. **CORS**: El servicio está configurado para aceptar requests de dominios autorizados
4. **Logs**: Todos los intentos de autenticación se registran para auditoría
5. **Timestamps**: Todos los timestamps están en formato ISO 8601 UTC
6. **UUIDs**: Los IDs de sesión y token son UUIDs v4
7. **Hashing**: Los tokens se almacenan hasheados por seguridad
