# API Error Codes

Lista de valores válidos para `error.code` en las respuestas de la API.

| Código | Descripción |
|-------|-------------|
| `ACCOUNT_INACTIVE` | Cuenta inactiva |
| `AUTH_FAILED` | Error interno al autenticar |
| `BAD_AUTH_HEADER` | Formato de cabecera Authorization inválido |
| `BAD_FORMAT` | Datos mal formateados |
| `BAD_QUERY` | Parámetros de consulta inválidos |
| `HEALTH_CHECK_FAILED` | Error interno al verificar salud |
| `INACTIVE_ACCOUNT` | Cuenta inactiva |
| `INVALID_AUTH_HEADER` | Formato de encabezado inválido |
| `INVALID_CREDENTIALS` | Credenciales inválidas |
| `INVALID_REFRESH_TOKEN` | Refresh token inválido o inactivo |
| `INVALID_TOKEN` | Token inválido o inactivo |
| `LOGIN_FAILED` | Error interno al realizar login |
| `LOGOUT_FAILED` | Error interno al cerrar sesión |
| `MISSING_SESSION_ID` | Falta session_id en la ruta |
| `MISSING_TOKEN` | Falta cabecera Authorization |
| `NO_AUTH_HEADER` | Encabezado de autorización requerido |
| `REFRESH_EXPIRED` | Refresh token expirado |
| `REFRESH_FAILED` | Error interno al refrescar token |
| `REVOKE_FAILED` | No se pudo revocar la sesión |
| `SESSION_ALREADY_REVOKED` | La sesión ya está revocada |
| `SESSION_INACTIVE` | Sesión inactiva o expirada |
| `SESSION_MISMATCH` | Token no pertenece a la sesión proporcionada |
| `SESSION_NOT_FOUND` | Sesión no encontrada |
| `VALIDATION_ERROR` | Error interno al validar token |
