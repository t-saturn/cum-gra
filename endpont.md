# Endpoints

## 1. Componentes globales

### 1.1 `utils/` (helpers genéricos)

| Función                                                          | Retorno              |
| ---------------------------------------------------------------- | -------------------- |
| \* `HashPassword(password string)`                               | `(string, error)`    |
| \* `CheckPasswordHash(password, encodedHash string)`             | `bool`               |
| \* `GenerateAccessToken(userID string)`                          | `(string, error)`    |
| \* `GenerateRefreshToken(userID string)`                         | `(string, error)`    |
| `NowUTC()`                                                       | `time.Time`          |
| `ParseISOTime(s string)`                                         | `(time.Time, error)` |
| `JSON(w http.ResponseWriter, status int, payload interface{})`   | `void`               |
| `JSONError(w http.ResponseWriter, status int, code, msg string)` | `void`               |

---

### 1.2 `repo/` (data‑access helpers)

| Repositorio               | Función                                                                                                              | Retorno                        |
| ------------------------- | -------------------------------------------------------------------------------------------------------------------- | ------------------------------ |
| **UserRepository**        | `FindActiveByEmailOrDNI(ctx context.Context, email, dni *string)`                                                    | `(*models.User, error)`        |
|                           | `CreateUser(ctx context.Context, u *models.User)`                                                                    | `(uuid.UUID, error)`           |
| **AuthAttemptRepository** | `Insert(ctx context.Context, a *models.AuthAttempt)`                                                                 | `error`                        |
| **SessionRepository**     | `Create(ctx context.Context, s *models.Session)`                                                                     | `(primitive.ObjectID, error)`  |
|                           | `FindByUUID(ctx context.Context, uuid string)`                                                                       | `(*models.Session, error)`     |
|                           | `UpdateStatus(ctx context.Context, id primitive.ObjectID, status string, revokedAt *time.Time)`                      | `error`                        |
| **TokenRepository**       | `Create(ctx context.Context, t *models.Token)`                                                                       | `(primitive.ObjectID, error)`  |
|                           | `FindByID(ctx context.Context, tokenID string)`                                                                      | `(*models.Token, error)`       |
|                           | `UpdateStatus(ctx context.Context, id primitive.ObjectID, status string, revokedAt *time.Time, lastUsed *time.Time)` | `error`                        |
|                           | `IncrementRefreshCount(ctx context.Context, id primitive.ObjectID)`                                                  | `error`                        |
| **AuthLogRepository**     | `Insert(ctx context.Context, l *models.AuthLog)`                                                                     | `error`                        |
| **TokenActivityRepo**     | `Insert(ctx context.Context, l *models.TokenActivityLog)`                                                            | `error`                        |
| **SessionActivityRepo**   | `Insert(ctx context.Context, l *models.SessionActivityLog)`                                                          | `error`                        |
| **CaptchaLogRepository**  | `Insert(ctx context.Context, c *models.CaptchaLog)`                                                                  | `error`                        |
| **TokenStatsRepository**  | `Query(ctx context.Context, userID primitive.ObjectID, from, to time.Time, groupBy string)`                          | `([]models.TokenStats, error)` |

---

### 1.3 `services/` (casos de uso reutilizables)

| Servicio           | Método                                                      | Retorno                          |
| ------------------ | ----------------------------------------------------------- | -------------------------------- |
| **AuthService**    | `VerifyCredentials(ctx, input dto.AuthVerifyRequest)`       | `(*dto.VerifyResponse, error)`   |
|                    | `Login(ctx, input dto.AuthLoginRequest)`                    | `(*dto.LoginResponse, error)`    |
|                    | `Logout(ctx, token string, input dto.LogoutRequest)`        | `(*dto.LogoutResponse, error)`   |
|                    | `RefreshToken(ctx, input dto.RefreshRequest)`               | `(*dto.RefreshResponse, error)`  |
|                    | `ValidateToken(ctx, input dto.ValidateRequest)`             | `(*dto.ValidateResponse, error)` |
| **SessionService** | `GetCurrent(ctx, token string)`                             | `(*dto.SessionResponse, error)`  |
|                    | `List(ctx, userID string, params dto.ListSessionsParams)`   | `([]dto.SessionInfo, error)`     |
|                    | `Revoke(ctx, userID, sessionID string)`                     | `(*dto.RevokeResponse, error)`   |
| **AuditService**   | `RegisterAttempt(ctx, input dto.ManualAttemptRequest)`      | `(*dto.AttemptResponse, error)`  |
|                    | `GetLogs(ctx, userID string, params dto.LogsParams)`        | `([]models.AuthLog, error)`      |
| **StatsService**   | `GetTokenStats(ctx, userID string, params dto.StatsParams)` | `(*dto.StatsResponse, error)`    |
| **HealthService**  | `Check(ctx context.Context)`                                | `(*dto.HealthResponse, error)`   |

---

### 1.4 `handlers/` (adaptadores HTTP genéricos)

| Handler            | Método                            | Retorno |
| ------------------ | --------------------------------- | ------- |
| **AuthHandler**    | `Verify(c *gin.Context)`          | `void`  |
|                    | `Login(c *gin.Context)`           | `void`  |
|                    | `Logout(c *gin.Context)`          | `void`  |
|                    | `Refresh(c *gin.Context)`         | `void`  |
|                    | `Validate(c *gin.Context)`        | `void`  |
| **SessionHandler** | `Me(c *gin.Context)`              | `void`  |
|                    | `List(c *gin.Context)`            | `void`  |
|                    | `Revoke(c *gin.Context)`          | `void`  |
| **AuditHandler**   | `RegisterAttempt(c *gin.Context)` | `void`  |
|                    | `GetLogs(c *gin.Context)`         | `void`  |
| **StatsHandler**   | `TokenStats(c *gin.Context)`      | `void`  |
| **HealthHandler**  | `Health(c *gin.Context)`          | `void`  |

---

## 2. Componentes específicos por endpoint

> **Formato:** `Función (paquete) — Retorno`

### 🔸 `/auth/verify`

- **Repo**:

  - `UserRepository.FindActiveByEmailOrDNI` — `(*models.User, error)`
  - `AuthAttemptRepository.Insert` — `error`
  - `CaptchaLogRepository.Insert` — `error`
  - `AuthLogRepository.Insert` — `error`

- **Service**:

  - `AuthService.VerifyCredentials` — `(*dto.VerifyResponse, error)`

- **Handler**:

  - `AuthHandler.Verify` — `void`

---

### 🔸 `/auth/login`

- **Repo**:

  - `UserRepository.FindActiveByEmailOrDNI` — `(*models.User, error)`
  - `AuthAttemptRepository.Insert` — `error`
  - `SessionRepository.Create` — `(primitive.ObjectID, error)`
  - `TokenRepository.Create` — `(primitive.ObjectID, error)`
  - `CaptchaLogRepository.Insert` — `error`
  - `TokenActivityRepo.Insert` — `error`
  - `SessionActivityRepo.Insert` — `error`
  - `AuthLogRepository.Insert` — `error`

- **Service**:

  - `AuthService.Login` — `(*dto.LoginResponse, error)`

- **Handler**:

  - `AuthHandler.Login` — `void`

---

### 🔸 `/auth/logout`

- **Repo**:

  - `TokenRepository.UpdateStatus` — `error`
  - `SessionRepository.UpdateStatus` — `error`
  - `AuthLogRepository.Insert` — `error`
  - `TokenActivityRepo.Insert` — `error`
  - `SessionActivityRepo.Insert` — `error`

- **Service**:

  - `AuthService.Logout` — `(*dto.LogoutResponse, error)`

- **Handler**:

  - `AuthHandler.Logout` — `void`

---

### 🔸 `/auth/token/refresh`

- **Repo**:

  - `TokenRepository.FindByID` — `(*models.Token, error)`
  - `TokenRepository.IncrementRefreshCount` — `error`
  - `TokenRepository.Create` — `(primitive.ObjectID, error)`
  - `SessionRepository.UpdateStatus` — `error`
  - `AuthLogRepository.Insert` — `error`
  - `TokenActivityRepo.Insert` — `error`

- **Service**:

  - `AuthService.RefreshToken` — `(*dto.RefreshResponse, error)`

- **Handler**:

  - `AuthHandler.Refresh` — `void`

---

### 🔸 `/auth/token/validate`

- **Repo**:

  - `TokenRepository.FindByID` — `(*models.Token, error)`
  - `TokenRepository.UpdateStatus` — `error`
  - `SessionRepository.UpdateStatus` — `error`
  - `AuthLogRepository.Insert` — `error`
  - `TokenActivityRepo.Insert` — `error`

- **Service**:

  - `AuthService.ValidateToken` — `(*dto.ValidateResponse, error)`

- **Handler**:

  - `AuthHandler.Validate` — `void`

---

### 🔸 `/auth/session/me`

- **Repo**:

  - `SessionRepository.FindByUUID` — `(*models.Session, error)`

- **Service**:

  - `SessionService.GetCurrent` — `(*dto.SessionResponse, error)`

- **Handler**:

  - `SessionHandler.Me` — `void`

---

### 🔸 `/auth/sessions`

- **Repo**:

  - `SessionRepository.ListByUser` — `([]models.Session, error)`

- **Service**:

  - `SessionService.List` — `([]dto.SessionInfo, error)`

- **Handler**:

  - `SessionHandler.List` — `void`

---

### 🔸 `DELETE /auth/sessions/{id}`

- **Repo**:

  - `SessionRepository.UpdateStatus` — `error`
  - `TokenRepository.UpdateStatusBySession` — `error`
  - `SessionActivityRepo.Insert` — `error`
  - `TokenActivityRepo.Insert` — `error`

- **Service**:

  - `SessionService.Revoke` — `(*dto.RevokeResponse, error)`

- **Handler**:

  - `SessionHandler.Revoke` — `void`

---

### 🔸 `POST /auth/attempts`

- **Repo**:

  - `AuthAttemptRepository.Insert` — `error`
  - `AuthLogRepository.Insert` — `error`

- **Service**:

  - `AuditService.RegisterAttempt` — `(*dto.AttemptResponse, error)`

- **Handler**:

  - `AuditHandler.RegisterAttempt` — `void`

---

### 🔸 `GET /auth/logs`

- **Repo**:

  - `AuthLogRepository.Query` — `([]models.AuthLog, error)`

- **Service**:

  - `AuditService.GetLogs` — `([]models.AuthLog, error)`

- **Handler**:

  - `AuditHandler.GetLogs` — `void`

---

### 🔸 `GET /auth/tokens/stats`

- **Repo**:

  - `TokenStatsRepository.Query` — `([]models.TokenStats, error)`

- **Service**:

  - `StatsService.GetTokenStats` — `(*dto.StatsResponse, error)`

- **Handler**:

  - `StatsHandler.TokenStats` — `void`

---

### 🔸 `GET /auth/health`

- **Repo**: _(ninguno)_

- **Service**:

  - `HealthService.Check` — `(*dto.HealthResponse, error)`

- **Handler**:

  - `HealthHandler.Health` — `void`
