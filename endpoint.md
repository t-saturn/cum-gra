# Endpoints

## 1. Componentes globales

### 1.1 `utils/` (helpers genéricos)

| Función                                                             | Retorno              |
| ------------------------------------------------------------------- | -------------------- |
| 🔸 `HashPassword(password string)`                                  | `(string, error)`    |
| 🔸 `CheckPasswordHash(password, encodedHash string)`                | `bool`               |
| 🔸 `GenerateAccessToken(userID string)`                             | `(string, error)`    |
| 🔸 `GenerateRefreshToken(userID string)`                            | `(string, error)`    |
| 🔸 `NowUTC()`                                                       | `time.Time`          |
| 🔸 `ParseISOTime(s string)`                                         | `(time.Time, error)` |
| 🔸 `JSON(w http.ResponseWriter, status int, payload interface{})`   | `void`               |
| 🔸 `JSONError(w http.ResponseWriter, status int, code, msg string)` | `void`               |

---

### 1.2 `repo/` (data‑access helpers)

| Repositorio               | Función                                                                                                                 | Retorno                       |
| ------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------- |
| **UserRepository**        | 🔸 `FindActiveByEmailOrDNI(ctx context.Context, email, dni *string)`                                                    | `(*UserData, error)`          |
| **AuthAttemptRepository** | 🔸 `Insert(ctx context.Context, a *models.AuthAttempt)`                                                                 | `error`                       |
| **SessionRepository**     | 🔸 `Create(ctx context.Context, s *models.Session)`                                                                     | `(primitive.ObjectID, error)` |
|                           | 🔸 `FindByUUID(ctx context.Context, uuid string)`                                                                       | `(*models.Session, error)`    |
|                           | 🔸 `FindByUserID(ctx context.Context, userID string, params dto.ListSessionsParams)`                                    | `([]models.Session, error)`   |
|                           | 🔸 `UpdateStatus(ctx context.Context, id primitive.ObjectID, status string, revokedAt *time.Time)`                      | `error`                       |
| **TokenRepository**       | 🔸 `Create(ctx context.Context, t *models.Token)`                                                                       | `(primitive.ObjectID, error)` |
|                           | 🔸 `FindByID(ctx context.Context, tokenID string)`                                                                      | `(*models.Token, error)`      |
|                           | 🔸 `UpdateStatus(ctx context.Context, id primitive.ObjectID, status string, revokedAt *time.Time, lastUsed *time.Time)` | `error`                       |
|                           | 🔸 `IncrementRefreshCount(ctx context.Context, id primitive.ObjectID)`                                                  | `error`                       |

---

### 1.3 `services/` (casos de uso reutilizables)

| Servicio           | Método                                                      | Retorno                                 |
| ------------------ | ----------------------------------------------------------- | --------------------------------------- |
| **AuthService**    | 🔸 `VerifyCredentials(ctx, input dto.AuthVerifyRequestDTO)` | `(*dto.AuthVerifyResponseDTO, error)`   |
|                    | `Login(ctx, input dto.AuthLoginRequestDTO)`                 | `(*dto.AuthLoginResponseDTO, error)`    |
|                    | `Logout(ctx, token string, input dto.AuthLogoutRequestDTO)` | `(*dto.AuthLogoutResponseDTO, error)`   |
|                    | `RefreshToken(ctx, input dto.AuthRefreshRequestDTO)`        | `(*dto.AuthRefreshResponseDTO, error)`  |
|                    | `ValidateToken(ctx, input dto.AuthValidateRequestDTO)`      | `(*dto.AuthValidateResponseDTO, error)` |
| **SessionService** | `GetCurrent(ctx, token string)`                             | `(*dto.SessionResponseDTO, error)`      |
|                    | `List(ctx, userID string, params dto.ListSessionsParams)`   | `([]dto.SessionInfoDTO, error)`         |
|                    | `Revoke(ctx, userID, sessionID string)`                     | `(*dto.RevokeResponseDTO, error)`       |
| **HealthService**  | `Check(ctx context.Context)`                                | `(*dto.HealthResponseDTO, error)`       |

---

### 1.4 `handlers/` (adaptadores HTTP genéricos)

| Handler            | Método                      | Retorno |
| ------------------ | --------------------------- | ------- |
| **AuthHandler**    | 🔸 `Verify(c *gin.Context)` | `void`  |
|                    | `Login(c *gin.Context)`     | `void`  |
|                    | `Logout(c *gin.Context)`    | `void`  |
|                    | `Refresh(c *gin.Context)`   | `void`  |
|                    | `Validate(c *gin.Context)`  | `void`  |
| **SessionHandler** | `Me(c *gin.Context)`        | `void`  |
|                    | `List(c *gin.Context)`      | `void`  |
|                    | `Revoke(c *gin.Context)`    | `void`  |
| **HealthHandler**  | `Health(c *gin.Context)`    | `void`  |

---

## `/auth/verify`

- **Repo**:

  - 🔸 `UserRepository.FindActiveByEmailOrDNI` — `(*UserData, error)`
  - 🔸 `AuthAttemptRepository.Insert` — `error`

- **Service**:

  - 🔸 `AuthService.VerifyCredentials` — `(*dto.AuthVerifyResponseDTO, error)`

- **Handler**:

  - 🔸 `AuthHandler.Verify` — `void`

---

## `/auth/login`

- **Repo**:

  - `UserRepository.FindActiveByEmailOrDNI` — `(*UserData, error)`
  - `AuthAttemptRepository.Insert` — `error`
  - `SessionRepository.Create` — `(primitive.ObjectID, error)`
  - `TokenRepository.Create` — `(primitive.ObjectID, error)`

- **Service**:

  - `AuthService.Login` — `(*dto.AuthLoginResponseDTO, error)`

- **Handler**:

  - `AuthHandler.Login` — `void`

---

## `/auth/logout`

- **Repo**:

  - `TokenRepository.UpdateStatus` — `error`
  - `SessionRepository.UpdateStatus` — `error`

- **Service**:

  - `AuthService.Logout` — `(*dto.AuthLogoutResponseDTO, error)`

- **Handler**:

  - `AuthHandler.Logout` — `void`

---

## `/auth/token/refresh`

- **Repo**:

  - `TokenRepository.FindByID` — `(*models.Token, error)`
  - `TokenRepository.IncrementRefreshCount` — `error`
  - `TokenRepository.Create` — `(primitive.ObjectID, error)`
  - `SessionRepository.UpdateStatus` — `error`

- **Service**:

  - `AuthService.RefreshToken` — `(*dto.AuthRefreshResponseDTO, error)`

- **Handler**:

  - `AuthHandler.Refresh` — `void`

---

## `/auth/token/validate`

- **Repo**:

  - `TokenRepository.FindByID` — `(*models.Token, error)`
  - `TokenRepository.UpdateStatus` — `error`
  - `SessionRepository.UpdateStatus` — `error`

- **Service**:

  - `AuthService.ValidateToken` — `(*dto.AuthValidateResponseDTO, error)`

- **Handler**:

  - `AuthHandler.Validate` — `void`

---

## `/auth/session/me`

- **Repo**:

  - `SessionRepository.FindByUUID` — `(*models.Session, error)`

- **Service**:

  - `SessionService.GetCurrent` — `(*dto.SessionResponseDTO, error)`

- **Handler**:

  - `SessionHandler.Me` — `void`

---

## `/auth/sessions`

- **Repo**:

  - `SessionRepository.ListByUser` — `([]models.Session, error)`

- **Service**:

  - `SessionService.List` — `([]dto.SessionInfoDTO, error)`

- **Handler**:

  - `SessionHandler.List` — `void`

---

## `DELETE /auth/sessions/{id}`

- **Repo**:

  - `SessionRepository.UpdateStatus` — `error`
  - `TokenRepository.UpdateStatusBySession` — `error`

- **Service**:

  - `SessionService.Revoke` — `(*dto.RevokeResponseDTO, error)`

- **Handler**:

  - `SessionHandler.Revoke` — `void`

---

## `GET /auth/health`

- **Repo**: _(ninguno)_

- **Service**:

  - `HealthService.Check` — `(*dto.HealthResponseDTO, error)`

- **Handler**:

  - `HealthHandler.Health` — `void`
