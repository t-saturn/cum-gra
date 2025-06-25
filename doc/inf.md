COLECCIÓN: user_credentials
Estado de validación de credenciales con expiración automática

```json
{
  _id: ObjectId,
  email: "usuario@empresa.com",
  passwordHash: "hash_bcrypt", // hash de la contraseña enviada

  // Estado de validación
  status: "pending", // "pending", "correct", "invalid", "expired"

  // Información de la aplicación desde donde se envían las credenciales
  applicationId: "app1",
  applicationUrl: "https://ventas.empresa.com",

  // Información del dispositivo/navegador
  deviceInfo: {
    userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
    ip: "192.168.1.100",
    deviceId: "device_fingerprint_abc123",
    browserName: "Chrome",
    browserVersion: "120.0.0.0",
    os: "Windows 10",
    location: {
      country: "Peru",
      city: "Lima",
      coordinates: [-12.0464, -77.0428]
    }
  },

  // Timestamps y expiración
  createdAt: ISODate,
  expiresAt: ISODate, // 5 minutos desde createdAt
  validatedAt: null, // cuando se cambió de pending a correct/invalid

  // Respuesta del servicio de validación
  validationResponse: {
    userId: ObjectId, // se llena cuando status = "correct"
    serviceResponse: "User validated successfully",
    validatedBy: "user-service-api"
  }
}
```

---

COLECCIÓN: tokens_activos
Tokens JWT válidos y activos

```json
{
  _id: ObjectId,
  tokenId: "jwt_jti_unique", // jti del JWT
  userId: ObjectId, // referencia a users
  sessionId: ObjectId, // referencia a sessions

  // Token info
  tokenHash: "sha256_hash_of_token", // hash del token para búsquedas rápidas
  tokenType: "access", // "access", "refresh"

  // Aplicación donde se usa el token
  applicationId: "app1",
  applicationUrl: "https://ventas.empresa.com",

  // Información del dispositivo
  deviceInfo: {
    userAgent: "Mozilla/5.0...",
    ip: "192.168.1.100",
    deviceId: "device_fingerprint_abc123",
    browserName: "Chrome",
    browserVersion: "120.0.0.0",
    os: "Windows 10",
    location: {
      country: "Peru",
      city: "Lima"
    }
  },

  // Timestamps
  issuedAt: ISODate,
  expiresAt: ISODate,
  lastUsed: ISODate, // se actualiza cada vez que se valida el token
  createdAt: ISODate,

  // Refresh info
  refreshCount: 0, // cuántas veces se ha refrescado
  maxRefreshCount: 10 // límite de refreshes
}
```

---

COLECCIÓN: tokens_invalid
Tokens que fueron invalidados (por seguridad, logout, etc.)

```json
{
  _id: ObjectId,
  tokenId: "jwt_jti_unique",
  tokenHash: "sha256_hash_of_token",
  userId: ObjectId,

  // Información de invalidación
  invalidatedAt: ISODate,
  invalidationReason: "user_logout", // "user_logout", "security_breach", "suspicious_activity", "admin_revoke", "invalid_token"
  invalidatedBy: "user", // "user", "admin", "system"

  // Información de contexto
  applicationId: "app1",
  deviceInfo: {
    userAgent: "Mozilla/5.0...",
    ip: "192.168.1.100",
    deviceId: "device_fingerprint_abc123"
  },

  // Info original del token
  originalIssuedAt: ISODate,
  originalExpiresAt: ISODate,

  // Para limpieza automática
  cleanupAt: ISODate // se elimina después de X tiempo
}
```

---

COLECCIÓN: tokens_expired
Tokens que expiraron automáticamente

```json
{
  _id: ObjectId,
  tokenId: "jwt_jti_unique",
  tokenHash: "sha256_hash_of_token",
  userId: ObjectId,

  // Información de expiración
  expiredAt: ISODate, // cuando se detectó que expiró
  originalExpiresAt: ISODate, // tiempo original de expiración

  // Contexto de la última actividad
  lastUsedAt: ISODate,
  applicationId: "app1",
  deviceInfo: {
    userAgent: "Mozilla/5.0...",
    ip: "192.168.1.100",
    deviceId: "device_fingerprint_abc123"
  },

  // Info del token
  tokenType: "access",
  issuedAt: ISODate,
  refreshCount: 3,

  // Para limpieza automática
  cleanupAt: ISODate // se elimina después de X tiempo para auditoría
}
```

---

COLECCIÓN: sessions
Sesiones activas de usuarios con información detallada

```json
{
  _id: ObjectId,
  sessionId: "session_uuid_456",
  userId: ObjectId,

  // Estado de la sesión
  isActive: true,
  createdAt: ISODate,
  lastActivity: ISODate,
  expiresAt: ISODate,

  // Información completa del dispositivo
  deviceInfo: {
    userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
    ip: "192.168.1.100",
    deviceId: "device_fingerprint_abc123", // fingerprint único del dispositivo
    browserName: "Chrome",
    browserVersion: "120.0.0.0",
    os: "Windows 10",
    osVersion: "10.0.19045",
    deviceType: "desktop", // "desktop", "mobile", "tablet"
    screenResolution: "1920x1080",
    timezone: "America/Lima",
    language: "es-PE",

    // Información de geolocalización
    location: {
      country: "Peru",
      countryCode: "PE",
      region: "Lima",
      city: "Lima",
      coordinates: [-12.0464, -77.0428],
      isp: "Telefonica del Peru",
      organization: "Movistar"
    }
  },

  // Métricas de la sesión
  metrics: {
    totalRequests: 147,
    totalTokenRefreshes: 5,
    applicationsAccessed: 3,
    lastApplicationUsed: "app1",
    sessionDuration: 7200 // segundos desde el inicio
  }
}
```

---

COLECCIÓN: auth_logs
Logs detallados de todas las operaciones de autenticación

```json
{
  _id: ObjectId,

  // Identificadores
  userId: ObjectId,
  sessionId: ObjectId,
  credentialId: "cred_uuid_123", // si viene de validación de credenciales
  tokenId: "jwt_jti_unique", // si está relacionado con un token

  // Información de la operación
  action: "credential_validation", // Ver lista completa abajo
  success: true,

  // Contexto de la aplicación
  applicationId: "app1",
  applicationUrl: "https://ventas.empresa.com",

  // Detalles específicos de la acción
  details: {
    // Para credential_validation
    credentialStatus: "correct", // "pending", "correct", "invalid", "expired"
    validationTime: 1500, // milisegundos que tomó validar

    // Para token_validation
    tokenType: "access",
    refreshed: true,
    refreshCount: 3,

    // Para errores
    errorCode: "INVALID_TOKEN",
    errorMessage: "Token has expired",

    // Para logout
    logoutType: "user_initiated", // "user_initiated", "session_expired", "admin_forced"

    // Información adicional
    userAgent: "Mozilla/5.0...",
    previousIP: "192.168.1.99", // si cambió de IP
    suspiciousActivity: false
  },

  // Información completa del dispositivo (igual que en sessions)
  deviceInfo: {
    userAgent: "Mozilla/5.0...",
    ip: "192.168.1.100",
    deviceId: "device_fingerprint_abc123",
    browserName: "Chrome",
    browserVersion: "120.0.0.0",
    os: "Windows 10",
    location: {
      country: "Peru",
      city: "Lima",
      coordinates: [-12.0464, -77.0428]
    }
  },

  // Timestamps
  timestamp: ISODate,
  processingTime: 150 // milisegundos que tomó procesar la operación
}
```

---

ACCIONES PARA AUTH_LOGS

```js
/*
Posibles valores para el campo "action":

Validación de credenciales:
- "credential_submit"      // Usuario envía credenciales
- "credential_validation"  // Validación con servicio externo
- "credential_expired"     // Credenciales expiraron (5 min)

Gestión de tokens:
- "token_generate"         // Se genera nuevo token
- "token_validate"         // Se valida token existente  
- "token_refresh"          // Se refresca token
- "token_invalidate"       // Token se marca como inválido
- "token_expire"           // Token expira automáticamente

Sesiones:
- "session_create"         // Nueva sesión
- "session_extend"         // Sesión se extiende por actividad
- "session_terminate"      // Sesión termina
- "sso_login"             // Login SSO en nueva aplicación

Autenticación:
- "login_attempt"         // Intento de login
- "login_success"         // Login exitoso
- "logout"                // Logout del usuario
- "forced_logout"         // Logout forzado por admin/sistema

Seguridad:
- "suspicious_activity"   // Actividad sospechosa detectada
- "security_violation"    // Violación de políticas de seguridad
- "device_change"         // Cambio de dispositivo detectado
- "location_change"       // Cambio de ubicación detectado
*/
```
