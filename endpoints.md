# auth-service-server – API Documentation

Este servicio maneja la autenticación centralizada del sistema SSO usando MongoDB y PostgreSQL. Registra intentos de acceso, emite tokens, gestiona sesiones, valida autenticación y expone logs de seguridad.

---

## Flujo General por Endpoint

Cada endpoint sigue esta estructura:

```text
HTTP Request → Handler → Service → Repository → Modelos
```

Por ejemplo:

```text
POST /auth/login
 → AuthHandler.Login
   → AuthService.Login
     → UserRepository.FindActiveByEmailOrDNI
     → SessionRepository.Create
     → TokenRepository.Create
     → AuthAttemptRepository.Insert
 → Modelos: AuthAttempt, Session, Token, AuthLog, etc.
```

---

## Endpoints disponibles

| Ruta                   | Método | Descripción               | Handler                 | Service               | Repos / Modelos clave                                                             |
| ---------------------- | ------ | ------------------------- | ----------------------- | --------------------- | --------------------------------------------------------------------------------- |
| `/auth/verify`         | POST   | Verificar credenciales    | `AuthHandler.Verify`    | `VerifyCredentials`   | `UserRepository`, `VerifyAttemp`                                                  |
| `/auth/login`          | POST   | Login completo            | `AuthHandler.Login`     | `Login`               | `UserRepository`, `SessionRepository`, `TokenRepository`, `AuthAttemptRepository` |
| `/auth/logout`         | POST   | Cerrar sesión             | `AuthHandler.Logout`    | `Logout`              | `TokenRepository`, `SessionRepository`                                            |
| `/auth/token/refresh`  | POST   | Refrescar token           | `AuthHandler.Refresh`   | `RefreshToken`        | `TokenRepository`, `SessionRepository`                                            |
| `/auth/token/validate` | POST   | Validar access token      | `AuthHandler.Validate`  | `ValidateToken`       | `TokenRepository`, `SessionRepository`                                            |
| `/auth/session/me`     | GET    | Ver sesión actual         | `SessionHandler.Me`     | `GetCurrent`          | `SessionRepository`                                                               |
| `/auth/sessions`       | GET    | Listar sesiones           | `SessionHandler.List`   | `List`                | `SessionRepository`                                                               |
| `/auth/sessions/{id}`  | DELETE | Revocar sesión específica | `SessionHandler.Revoke` | `Revoke`              | `SessionRepository`, `TokenRepository`                                            |
| `/auth/attempts`       | POST   | Registrar intento manual  | -                       | -                     | `AuthAttemptRepository`                                                           |
| `/auth/logs`           | GET    | Ver logs de usuario       | -                       | -                     | `AuthLog`, `Token`                                                                |
| `/auth/tokens/stats`   | GET    | Estadísticas de tokens    | -                       | -                     | `TokenStats`, `Token`                                                             |
| `/auth/health`         | GET    | Ver estado del servicio   | `HealthHandler.Health`  | `HealthService.Check` | _(sin repos)_                                                                     |

---

## Relación Endpoints ↔ Repos ↔ Servicios

### `/auth/verify`

- **Handler:** `AuthHandler.Verify`
- **Servicio:** `AuthService.VerifyCredentials`
- **Repositorios:**

  - `UserRepository.FindActiveByEmailOrDNI`
  - `AuthAttemptRepository.Insert`

- **Modelos:** `User`, `AuthAttempt`, `AuthLog`, `CaptchaLog`

---

### `/auth/login`

- **Handler:** `AuthHandler.Login`
- **Servicio:** `AuthService.Login`
- **Repositorios:**

  - `UserRepository.FindActiveByEmailOrDNI`
  - `AuthAttemptRepository.Insert`
  - `SessionRepository.Create`
  - `TokenRepository.Create`

- **Modelos:** `AuthAttempt`, `Session`, `Token`

---

### `/auth/logout`

- **Handler:** `AuthHandler.Logout`
- **Servicio:** `AuthService.Logout`
- **Repositorios:**

  - `TokenRepository.UpdateStatus`
  - `SessionRepository.UpdateStatus`

- **Modelos:** `Token`, `Session`, `SessionActivityLog`

---

### `/auth/token/refresh`

- **Handler:** `AuthHandler.Refresh`
- **Servicio:** `AuthService.RefreshToken`
- **Repositorios:**

  - `TokenRepository.FindByID`
  - `TokenRepository.IncrementRefreshCount`
  - `TokenRepository.Create`
  - `SessionRepository.UpdateStatus`

- **Modelos:** `Token`, `Session`

---

### `/auth/token/validate`

- **Handler:** `AuthHandler.Validate`
- **Servicio:** `AuthService.ValidateToken`
- **Repositorios:**

  - `TokenRepository.FindByID`
  - `TokenRepository.UpdateStatus`
  - `SessionRepository.UpdateStatus`

- **Modelos:** `Token`, `Session`

---

### `/auth/session/me`

- **Handler:** `SessionHandler.Me`
- **Servicio:** `SessionService.GetCurrent`
- **Repositorios:**

  - `SessionRepository.FindByUUID`

- **Modelos:** `Session`, `Token`, `AuthLog`

---

### `/auth/sessions`

- **Handler:** `SessionHandler.List`
- **Servicio:** `SessionService.List`
- **Repositorios:**

  - `SessionRepository.ListByUser`

- **Modelos:** `Session`, `Token`, `AuthLog`

---

### `/auth/sessions/{id}`

- **Handler:** `SessionHandler.Revoke`
- **Servicio:** `SessionService.Revoke`
- **Repositorios:**

  - `SessionRepository.UpdateStatus`
  - `TokenRepository.UpdateStatusBySession`

- **Modelos:** `Session`, `Token`, `SessionActivityLog`, `TokenActivityLog`

---

### 👤 **Rutas para el Usuario (sesiones personales)**

| Método | Ruta                  | Descripción                                                 |
| ------ | --------------------- | ----------------------------------------------------------- |
| `GET`  | `/session/me`         | Obtener **todas** las sesiones propias (activas/inactivas). |
| `GET`  | `/session/me/current` | Obtener **la sesión actual** del usuario.                   |

> Puedes agregar filtros opcionales por query param si deseas algo como `/session/me?status=active`.

---

### 🛡️ **Rutas para el Administrador (gestiona cualquier usuario)**

| Método | Ruta                             | Descripción                                                     |
| ------ | -------------------------------- | --------------------------------------------------------------- |
| `GET`  | `/session/user/:user_id`         | Ver todas las sesiones de un usuario específico.                |
| `GET`  | `/session/user/:user_id/current` | Ver la sesión actual (si está activa) de un usuario.            |
| `GET`  | `/session`                       | Listar sesiones de **todos los usuarios** (opcional, paginada). |

> Estas rutas deben requerir permisos especiales en middleware para asegurar que solo los administradores puedan accederlas.

---

```js
GET /session/me?status=active
GET /session/user/abc123?status=inactive&limit=10&page=2
```
