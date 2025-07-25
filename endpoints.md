# auth-service-server – API Documentation

Este servicio maneja la autenticación centralizada del sistema SSO. Registra intentos de acceso, emite tokens, gestiona sesiones, valida autenticación y expone logs de seguridad.

---

## Autenticación por credenciales

### 🔸 Verificar credenciales

- **Método:** `POST`
- **Nombre:** Verificar usuario y contraseña
- **Ruta:** `/auth/verify`
- **Respuesta:**

```json
{
  "user_id": "ObjectID",
  "status": "success",
  "validated_by": "credentials"
}
```

- **Descripción:** Verifica si las credenciales del usuario son válidas. Registra un `AuthAttempt` (success o failed). No genera tokens ni sesión.

---

### 🔸 Iniciar sesión

- **Método:** `POST`
- **Nombre:** Login completo
- **Ruta:** `/auth/login`
- **Respuesta:**

  ```json
  {
    "access_token": "JWT...",
    "refresh_token": "JWT...",
    "session_id": "abc123",
    "expires_at": "2025-08-01T12:00:00Z"
  }
  ```

- **Descripción:** Valida credenciales, crea una sesión y emite tokens. Registra logs y device info. Requiere `application_id`.

---

### 🔸 Cerrar sesión

- **Método:** `POST`
- **Nombre:** Logout
- **Ruta:** `/auth/logout`
- **Respuesta:**

  ```json
  {
    "message": "Session and tokens revoked successfully"
  }
  ```

- **Descripción:** Revoca el `access_token` actual y cierra la sesión. También registra un `AuthLog`.

---

## Manejo de Tokens

### 🔸 Refrescar token

- **Método:** `POST`
- **Nombre:** Refresh token
- **Ruta:** `/auth/token/refresh`
- **Respuesta:**

  ```json
  {
    "access_token": "nuevo_token",
    "expires_at": "2025-08-01T12:00:00Z"
  }
  ```

- **Descripción:** Genera un nuevo `access_token` si el `refresh_token` es válido. Actualiza estadísticas y logs.

---

### 🔸 Validar token

- **Método:** `POST`
- **Nombre:** Validación de token
- **Ruta:** `/auth/token/validate`
- **Respuesta:**

  ```json
  {
    "valid": true,
    "user_id": "ObjectID",
    "application_id": "app-uuid"
  }
  ```

- **Descripción:** Valida si un `access_token` es válido y activo. Usado por gateway o WebSocket.

---

## Sesiones

### 🔸 Obtener sesión actual

- **Método:** `GET`
- **Nombre:** Obtener sesión activa
- **Ruta:** `/auth/session/me`
- **Respuesta:**

  ```json
  {
    "user_id": "ObjectID",
    "session_id": "abc123",
    "application_id": "app-uuid",
    "device_info": { ... }
  }
  ```

- **Descripción:** Devuelve la información de la sesión actual basada en el token recibido.

---

### 🔸 Listar todas las sesiones activas

- **Método:** `GET`
- **Nombre:** Listar sesiones activas
- **Ruta:** `/auth/sessions`
- **Respuesta:**

  ```json
  [
    {
      "session_id": "abc123",
      "created_at": "...",
      "device_info": { ... }
    }
  ]
  ```

- **Descripción:** Muestra todas las sesiones activas del usuario. Útil para cerrar sesión en dispositivos.

---

## Auditoría y Seguridad

### 🔸 Registrar intento manual

- **Método:** `POST`
- **Nombre:** Registrar intento de autenticación
- **Ruta:** `/auth/attempts`
- **Descripción:** Registra un intento de login por token o credenciales, independientemente del éxito. Útil para registrar accesos inválidos.

---

### 🔸 Ver logs del usuario actual

- **Método:** `GET`
- **Nombre:** Ver logs personales
- **Ruta:** `/auth/logs`
- **Respuesta:**

  ```json
  [
    {
      "action": "login",
      "success": true,
      "timestamp": "2025-07-21T13:00:00Z",
      "application_id": "app-uuid",
      "device_info": { ... }
    }
  ]
  ```

- **Descripción:** Devuelve el historial reciente de actividades de autenticación del usuario actual.

---

## Salud del servicio

### 🔸 Endpoint de salud

- **Método:** `GET`
- **Nombre:** Health check
- **Ruta:** `/auth/health`
- **Respuesta:**

  ```json
  {
    "status": "ok"
  }
  ```

- **Descripción:** Endpoint simple para verificar que el servicio está activo. Usado para readiness/liveness probes.

---

## 🔐 Reglas de autorización

- Los endpoints `/auth/login`, `/auth/verify`, `/auth/token/refresh` y `/auth/attempts` **no requieren token**.
- Los endpoints `/auth/logout`, `/auth/session/me`, `/auth/sessions`, `/auth/logs` y `/auth/token/validate` **requieren `access_token` válido en header `Authorization`**.
