# auth-service-server – API Documentation

Este servicio maneja la autenticación centralizada del sistema SSO usando MongoDB. Registra intentos de acceso, emite tokens, gestiona sesiones, valida autenticación y expone logs de seguridad.

## Estructura de Modelos

### Modelos Principales

- **`AuthAttempt`** - Registra todos los intentos de autenticación
- **`Session`** - Gestiona sesiones de usuarios activas/inactivas
- **`Token`** - Maneja tokens de acceso y refresh
- **`AuthLog`** - Log de auditoría de todas las actividades de autenticación
- **`CaptchaLog`** - Registra validaciones de CAPTCHA
- **`TokenActivityLog`** - Log detallado de actividad de tokens
- **`SessionActivityLog`** - Log detallado de actividad de sesiones

---

## Autenticación por credenciales

### 🔸 Verificar credenciales

- **Método:** `POST`
- **Nombre:** Verificar usuario y contraseña
- **Ruta:** `/auth/verify`

**Body:**

```json
{
  "email": "user@example.com", // opcional
  "dni": "12345678", // opcional (email o dni requerido)
  "password": "password123",
  "application_id": "app-uuid",
  "device_info": {
    "user_agent": "Mozilla/5.0...",
    "ip": "192.168.1.1",
    "device_id": "device-123",
    "browser_name": "Chrome",
    "os": "Windows"
  },
  "captcha_token": "captcha-token" // opcional
}
```

**Respuesta:**

```json
{
  "user_id": "ObjectID",
  "status": "success",
  "validated_by": "credentials",
  "auth_attempt_id": "ObjectID"
}
```

**Modelos afectados:**

- **`AuthAttempt`** - Crea registro con `status: "success"/"failed"`
- **`AuthLog`** - Registra evento de verificación
- **`CaptchaLog`** - Si se proporciona captcha_token
- **`ValidationResponse`** - Respuesta del servicio de validación

**Descripción:** Verifica si las credenciales del usuario son válidas. Registra un `AuthAttempt` (success o failed). No genera tokens ni sesión.

---

### 🔸 Iniciar sesión

- **Método:** `POST`
- **Nombre:** Login completo
- **Ruta:** `/auth/login`

**Body:**

```json
{
  "email": "user@example.com", // opcional
  "dni": "12345678", // opcional (email o dni requerido)
  "password": "password123",
  "application_id": "app-uuid",
  "application_url": "https://app.example.com",
  "device_info": {
    "user_agent": "Mozilla/5.0...",
    "ip": "192.168.1.1",
    "device_id": "device-123",
    "browser_name": "Chrome",
    "browser_version": "91.0",
    "os": "Windows",
    "os_version": "10",
    "device_type": "desktop",
    "timezone": "America/Lima",
    "language": "es-PE",
    "location": {
      "country": "Peru",
      "city": "Lima",
      "coordinates": [-77.0428, -12.0464]
    }
  },
  "captcha_token": "captcha-token"
}
```

**Respuesta:**

```json
{
  "access_token": "JWT...",
  "refresh_token": "JWT...",
  "session_id": "abc123",
  "expires_at": "2025-08-01T12:00:00Z",
  "user_id": "ObjectID"
}
```

**Modelos afectados:**

- **`AuthAttempt`** - Registra intento de login exitoso
- **`Session`** - Crea nueva sesión activa con métricas
- **`Token`** - Crea 2 tokens (access + refresh) pareados
- **`AuthLog`** - Registra evento de login completo
- **`TokenActivityLog`** - Registra creación de tokens
- **`SessionActivityLog`** - Registra creación de sesión
- **`CaptchaLog`** - Si se valida CAPTCHA

**Descripción:** Valida credenciales, crea una sesión y emite tokens. Registra logs y device info. Requiere `application_id`.

---

### 🔸 Cerrar sesión

- **Método:** `POST`
- **Nombre:** Logout
- **Ruta:** `/auth/logout`
- **Headers:** `Authorization: Bearer <access_token>`

**Body:**

```json
{
  "logout_type": "user_initiated", // opcional: user_initiated, admin_revoked, security_breach
  "revoke_all_sessions": false // opcional: si revocar todas las sesiones del usuario
}
```

**Respuesta:**

```json
{
  "message": "Session and tokens revoked successfully",
  "revoked_tokens": 2,
  "session_id": "abc123"
}
```

**Modelos afectados:**

- **`Session`** - Actualiza status a "revoked", establece `revoked_at`
- **`Token`** - Revoca access_token y refresh_token asociados
- **`AuthLog`** - Registra evento de logout
- **`TokenActivityLog`** - Registra revocación de tokens
- **`SessionActivityLog`** - Registra fin de sesión

**Descripción:** Revoca el `access_token` actual y cierra la sesión. También registra un `AuthLog`.

---

## Manejo de Tokens

### 🔸 Refrescar token

- **Método:** `POST`
- **Nombre:** Refresh token
- **Ruta:** `/auth/token/refresh`

**Body:**

```json
{
  "refresh_token": "JWT...",
  "application_id": "app-uuid",
  "device_info": {
    "ip": "192.168.1.1",
    "user_agent": "Mozilla/5.0..."
  }
}
```

**Respuesta:**

```json
{
  "access_token": "nuevo_JWT...",
  "expires_at": "2025-08-01T12:00:00Z",
  "token_id": "new-token-id"
}
```

**Modelos afectados:**

- **`Token`** - Crea nuevo access_token, incrementa `refresh_count` del refresh_token
- **`AuthLog`** - Registra evento de token refresh
- **`TokenActivityLog`** - Registra uso del refresh_token y creación del nuevo access_token
- **`Session`** - Actualiza `last_activity` y métricas de `total_token_refreshes`

**Descripción:** Genera un nuevo `access_token` si el `refresh_token` es válido. Actualiza estadísticas y logs.

---

### 🔸 Validar token

- **Método:** `POST`
- **Nombre:** Validación de token
- **Ruta:** `/auth/token/validate`

**Body:**

```json
{
  "access_token": "JWT...",
  "application_id": "app-uuid", // opcional
  "check_session": true // opcional: verificar si la sesión está activa
}
```

**Respuesta:**

```json
{
  "valid": true,
  "user_id": "ObjectID",
  "session_id": "abc123",
  "application_id": "app-uuid",
  "expires_at": "2025-08-01T12:00:00Z",
  "token_type": "access",
  "scopes": ["read", "write"]
}
```

**Modelos afectados:**

- **`Token`** - Actualiza `last_used`
- **`AuthLog`** - Registra evento de validación de token
- **`TokenActivityLog`** - Registra uso/validación del token
- **`Session`** - Actualiza `last_activity` si `check_session=true`

**Descripción:** Valida si un `access_token` es válido y activo. Usado por gateway o WebSocket.

---

## 🎮 Sesiones

### 🔸 Obtener sesión actual

- **Método:** `GET`
- **Nombre:** Obtener sesión activa
- **Ruta:** `/auth/session/me`
- **Headers:** `Authorization: Bearer <access_token>`

**Respuesta:**

```json
{
  "user_id": "ObjectID",
  "session_id": "abc123",
  "application_id": "app-uuid",
  "status": "active",
  "created_at": "2025-07-21T12:00:00Z",
  "last_activity": "2025-07-21T13:30:00Z",
  "expires_at": "2025-08-01T12:00:00Z",
  "device_info": {
    "browser_name": "Chrome",
    "os": "Windows",
    "device_type": "desktop",
    "ip": "192.168.1.1",
    "location": {
      "country": "Peru",
      "city": "Lima"
    }
  },
  "metrics": {
    "total_requests": 45,
    "total_token_refreshes": 3,
    "applications_accessed": 2,
    "session_duration": 5400
  }
}
```

**Modelos afectados:**

- **`Session`** - Lee información de la sesión actual
- **`Token`** - Verifica validez del access_token
- **`AuthLog`** - Registra consulta de sesión

**Descripción:** Devuelve la información de la sesión actual basada en el token recibido.

---

### 🔸 Listar todas las sesiones activas

- **Método:** `GET`
- **Nombre:** Listar sesiones activas
- **Ruta:** `/auth/sessions`
- **Headers:** `Authorization: Bearer <access_token>`

**Query params:**

- `status` - Filtrar por estado: active, inactive, revoked, expired
- `limit` - Límite de resultados (default: 10)
- `offset` - Offset para paginación

**Respuesta:**

```json
{
  "sessions": [
    {
      "session_id": "abc123",
      "status": "active",
      "created_at": "2025-07-21T12:00:00Z",
      "last_activity": "2025-07-21T13:30:00Z",
      "device_info": {
        "browser_name": "Chrome",
        "os": "Windows",
        "device_type": "desktop",
        "ip": "192.168.1.1",
        "location": {
          "country": "Peru",
          "city": "Lima"
        }
      },
      "is_current": true
    }
  ],
  "total": 3,
  "active_count": 2
}
```

**Modelos afectados:**

- **`Session`** - Lee todas las sesiones del usuario
- **`Token`** - Verifica access_token del usuario
- **`AuthLog`** - Registra consulta de sesiones

**Descripción:** Muestra todas las sesiones del usuario. Útil para cerrar sesión en dispositivos.

---

### 🔸 Revocar sesión específica

- **Método:** `DELETE`
- **Nombre:** Revocar sesión
- **Ruta:** `/auth/sessions/{session_id}`
- **Headers:** `Authorization: Bearer <access_token>`

**Respuesta:**

```json
{
  "message": "Session revoked successfully",
  "session_id": "abc123",
  "revoked_tokens": 2
}
```

**Modelos afectados:**

- **`Session`** - Actualiza status a "revoked"
- **`Token`** - Revoca tokens asociados a la sesión
- **`SessionActivityLog`** - Registra revocación de sesión
- **`TokenActivityLog`** - Registra revocación de tokens

---

## Auditoría y Seguridad

### 🔸 Registrar intento manual

- **Método:** `POST`
- **Nombre:** Registrar intento de autenticación
- **Ruta:** `/auth/attempts`

**Body:**

```json
{
  "method": "credentials", // credentials | token
  "email": "user@example.com",
  "token": "JWT...", // si method = token
  "application_id": "app-uuid",
  "device_info": { ... },
  "status": "failed", // pending | success | failed | expired
  "error_code": "INVALID_PASSWORD",
  "notes": "Suspicious activity detected"
}
```

**Respuesta:**

```json
{
  "auth_attempt_id": "ObjectID",
  "message": "Auth attempt registered successfully"
}
```

**Modelos afectados:**

- **`AuthAttempt`** - Crea registro manual del intento
- **`AuthLog`** - Registra el evento

**Descripción:** Registra un intento de login por token o credenciales, independientemente del éxito. Útil para registrar accesos inválidos.

---

### 🔸 Ver logs del usuario actual

- **Método:** `GET`
- **Nombre:** Ver logs personales
- **Ruta:** `/auth/logs`
- **Headers:** `Authorization: Bearer <access_token>`

**Query params:**

- `action` - Filtrar por acción: login, logout, token_validation, etc.
- `success` - Filtrar por éxito: true/false
- `application_id` - Filtrar por aplicación
- `limit` - Límite de resultados (default: 50)
- `from_date` - Fecha desde (ISO 8601)
- `to_date` - Fecha hasta (ISO 8601)

**Respuesta:**

```json
{
  "logs": [
    {
      "id": "ObjectID",
      "action": "login",
      "success": true,
      "timestamp": "2025-07-21T13:00:00Z",
      "application_id": "app-uuid",
      "session_id": "abc123",
      "device_info": {
        "browser_name": "Chrome",
        "os": "Windows",
        "ip": "192.168.1.1",
        "location": {
          "country": "Peru",
          "city": "Lima"
        }
      },
      "details": {
        "credential_status": "correct",
        "validation_time": 150,
        "processing_time": 200
      }
    }
  ],
  "total": 25,
  "summary": {
    "successful_logins": 20,
    "failed_attempts": 5,
    "unique_applications": 3,
    "unique_devices": 2
  }
}
```

**Modelos afectados:**

- **`AuthLog`** - Lee logs de autenticación del usuario
- **`Token`** - Verifica access_token

**Descripción:** Devuelve el historial reciente de actividades de autenticación del usuario actual.

---

### 🔸 Ver estadísticas de tokens

- **Método:** `GET`
- **Nombre:** Estadísticas de tokens
- **Ruta:** `/auth/tokens/stats`
- **Headers:** `Authorization: Bearer <access_token>`

**Query params:**

- `from_date` - Fecha desde
- `to_date` - Fecha hasta
- `group_by` - Agrupar por: day, week, month

**Respuesta:**

```json
{
  "period": {
    "from": "2025-07-01T00:00:00Z",
    "to": "2025-07-31T23:59:59Z"
  },
  "stats": [
    {
      "date": "2025-07-21",
      "total_tokens_issued": 15,
      "access_tokens_issued": 8,
      "refresh_tokens_issued": 7,
      "tokens_revoked": 2,
      "token_usages": 120,
      "token_refreshes": 5,
      "unique_applications": 3,
      "unique_devices": 2
    }
  ]
}
```

**Modelos afectados:**

- **`TokenStats`** - Lee estadísticas agregadas de tokens
- **`Token`** - Verifica access_token

---

## Salud del servicio

### 🔸 Endpoint de salud

- **Método:** `GET`
- **Nombre:** Health check
- **Ruta:** `/auth/health`

**Respuesta:**

```json
{
  "status": "ok",
  "timestamp": "2025-07-21T13:00:00Z",
  "version": "1.0.0",
  "dependencies": {
    "mongodb": "connected",
    "postgres": "connected"
  },
  "metrics": {
    "active_sessions": 1250,
    "active_tokens": 2500,
    "total_attempts_today": 5000
  }
}
```

**Descripción:** Endpoint simple para verificar que el servicio está activo. Usado para readiness/liveness probes.

---

## Reglas de autorización

### Endpoints públicos (sin token)

- `POST /auth/verify` - Verificar credenciales
- `POST /auth/login` - Iniciar sesión
- `POST /auth/token/refresh` - Refrescar token
- `POST /auth/token/validate` - Validar token
- `POST /auth/attempts` - Registrar intento
- `GET /auth/health` - Health check

### Endpoints protegidos (requieren `Authorization: Bearer <access_token>`)

- `POST /auth/logout` - Cerrar sesión
- `GET /auth/session/me` - Sesión actual
- `GET /auth/sessions` - Listar sesiones
- `DELETE /auth/sessions/{session_id}` - Revocar sesión
- `GET /auth/logs` - Ver logs personales
- `GET /auth/tokens/stats` - Estadísticas de tokens

---

## Estructura de respuestas de error

```json
{
  "error": "Mensaje de error legible",
  "code": "ERROR_CODE",
  "timestamp": "2025-07-21T13:00:00Z",
  "request_id": "uuid",
  "details": {
    "field": "validation error details"
  }
}
```

## 🔗 Relaciones entre modelos

```txt
User (PostgreSQL)
    ↓
AuthAttempt → ValidationResponse
    ↓
Session → SessionActivityLog
    ↓
Token (access/refresh pareados) → TokenActivityLog
    ↓
AuthLog ← CaptchaLog
    ↓
TokenStats (agregados)
```
