# Lista de Endpoints API - Sistema de Gestión de Usuarios

## AUTENTICACIÓN Y SESIONES

### Autenticación Básica

```
POST   /api/v1/auth/login              # Iniciar sesión
POST   /api/v1/auth/logout             # Cerrar sesión
POST   /api/v1/auth/refresh            # Renovar token
POST   /api/v1/auth/forgot-password    # Solicitar reset de contraseña
POST   /api/v1/auth/reset-password     # Confirmar reset de contraseña
POST   /api/v1/auth/verify-email       # Verificar email
POST   /api/v1/auth/resend-verification # Reenviar verificación
```

### Gestión de Sesiones

```
GET    /api/v1/sessions                # Listar sesiones activas del usuario
DELETE /api/v1/sessions/:session_id    # Cerrar sesión específica
DELETE /api/v1/sessions/all           # Cerrar todas las sesiones
GET    /api/v1/sessions/history        # Historial de sesiones
```

### Autenticación Multifactor (MFA)

```
POST   /api/v1/mfa/enable              # Habilitar MFA
POST   /api/v1/mfa/disable             # Deshabilitar MFA
POST   /api/v1/mfa/verify              # Verificar código MFA
GET    /api/v1/mfa/devices             # Listar dispositivos MFA
DELETE /api/v1/mfa/devices/:mfa_id     # Eliminar dispositivo MFA
```

---

## GESTIÓN DE USUARIOS

### CRUD Usuarios

```
GET    /api/v1/users                   # Listar usuarios (con filtros y paginación)
GET    /api/v1/users/:id               # Obtener usuario por ID
POST   /api/v1/users                   # Crear nuevo usuario
PUT    /api/v1/users/:id               # Actualizar usuario completo
PATCH  /api/v1/users/:id               # Actualizar campos específicos
DELETE /api/v1/users/:id               # Eliminación lógica de usuario
```

### Perfil de Usuario

```
GET    /api/v1/users/profile           # Obtener perfil del usuario logueado
PUT    /api/v1/users/profile           # Actualizar perfil propio
POST   /api/v1/users/change-password   # Cambiar contraseña
GET    /api/v1/users/password-history  # Historial de contraseñas
```

### Historial y Auditoría de Usuarios

```
GET    /api/v1/users/:id/history       # Historial de cambios de usuario
GET    /api/v1/users/:id/audit         # Log de auditoría del usuario
```

---

## UNIDADES ORGÁNICAS

### CRUD Unidades Orgánicas

```
GET    /api/v1/organic-units           # Listar unidades (jerarquía)
GET    /api/v1/organic-units/:id       # Obtener unidad por ID
POST   /api/v1/organic-units           # Crear nueva unidad
PUT    /api/v1/organic-units/:id       # Actualizar unidad
DELETE /api/v1/organic-units/:id       # Eliminación lógica
```

### Jerarquía y Estructura

```
GET    /api/v1/organic-units/hierarchy # Obtener jerarquía completa
GET    /api/v1/organic-units/:id/children # Obtener unidades hijas
GET    /api/v1/organic-units/:id/parent   # Obtener unidad padre
GET    /api/v1/organic-units/tree      # Árbol completo de unidades
```

### Historial de Unidades

```
GET    /api/v1/organic-units/:id/history # Historial de cambios
```

---

## POSICIONES ESTRUCTURALES

### CRUD Posiciones

```
GET    /api/v1/positions               # Listar posiciones
GET    /api/v1/positions/:id           # Obtener posición por ID
POST   /api/v1/positions               # Crear nueva posición
PUT    /api/v1/positions/:id           # Actualizar posición
DELETE /api/v1/positions/:id           # Eliminación lógica
```

### Posiciones por Unidad Orgánica

```
GET    /api/v1/organic-units/:id/positions # Posiciones de una unidad
GET    /api/v1/positions/heads         # Posiciones de jefatura
```

### Asignaciones de Posiciones

```
GET    /api/v1/user-positions          # Listar asignaciones usuario-posición
POST   /api/v1/user-positions          # Asignar usuario a posición
PUT    /api/v1/user-positions/:id      # Actualizar asignación
DELETE /api/v1/user-positions/:id      # Finalizar asignación
GET    /api/v1/users/:id/positions     # Posiciones de un usuario
GET    /api/v1/positions/:id/users     # Usuarios en una posición
```

---

## MOVIMIENTOS DE PERSONAL

### Gestión de Movimientos

```
GET    /api/v1/personnel-movements     # Listar movimientos
GET    /api/v1/personnel-movements/:id # Obtener movimiento por ID
POST   /api/v1/personnel-movements     # Crear nuevo movimiento
PUT    /api/v1/personnel-movements/:id # Actualizar movimiento
DELETE /api/v1/personnel-movements/:id # Eliminar movimiento
```

### Movimientos por Usuario/Unidad

```
GET    /api/v1/users/:id/movements     # Movimientos de un usuario
GET    /api/v1/organic-units/:id/movements # Movimientos de una unidad
```

---

## SISTEMAS

### CRUD Sistemas

```
GET    /api/v1/systems                 # Listar sistemas
GET    /api/v1/systems/:id             # Obtener sistema por ID
POST   /api/v1/systems                 # Crear nuevo sistema
PUT    /api/v1/systems/:id             # Actualizar sistema
DELETE /api/v1/systems/:id             # Eliminación lógica
```

### Configuración de Sistemas

```
GET    /api/v1/systems/:id/config      # Obtener configuración
PUT    /api/v1/systems/:id/config      # Actualizar configuración
GET    /api/v1/systems/:id/history     # Historial de cambios
```

---

## ROLES Y PERMISOS

### Gestión de Roles

```
GET    /api/v1/roles                   # Listar roles
GET    /api/v1/roles/:id               # Obtener rol por ID
POST   /api/v1/roles                   # Crear nuevo rol
PUT    /api/v1/roles/:id               # Actualizar rol
DELETE /api/v1/roles/:id               # Eliminación lógica
GET    /api/v1/systems/:id/roles       # Roles de un sistema
```

### Gestión de Permisos

```
GET    /api/v1/permissions             # Listar permisos
GET    /api/v1/permissions/:id         # Obtener permiso por ID
POST   /api/v1/permissions             # Crear nuevo permiso
PUT    /api/v1/permissions/:id         # Actualizar permiso
DELETE /api/v1/permissions/:id         # Eliminación lógica
GET    /api/v1/systems/:id/permissions # Permisos de un sistema
```

### Asignación Rol-Permisos

```
GET    /api/v1/role-permissions        # Listar asignaciones rol-permiso
POST   /api/v1/role-permissions        # Asignar permiso a rol
DELETE /api/v1/role-permissions/:id    # Eliminar asignación
GET    /api/v1/roles/:id/permissions   # Permisos de un rol
GET    /api/v1/permissions/:id/roles   # Roles que tienen un permiso
```

### Asignación Usuario-Roles

```
GET    /api/v1/user-roles              # Listar asignaciones usuario-rol
POST   /api/v1/user-roles              # Asignar rol a usuario
PUT    /api/v1/user-roles/:id          # Actualizar asignación
DELETE /api/v1/user-roles/:id          # Eliminar asignación
GET    /api/v1/users/:id/roles         # Roles de un usuario
GET    /api/v1/roles/:id/users         # Usuarios con un rol
```

---

## MÓDULOS

### Gestión de Módulos

```
GET    /api/v1/modules                 # Listar módulos
GET    /api/v1/modules/:id             # Obtener módulo por ID
POST   /api/v1/modules                 # Crear nuevo módulo
PUT    /api/v1/modules/:id             # Actualizar módulo
DELETE /api/v1/modules/:id             # Eliminación lógica
GET    /api/v1/systems/:id/modules     # Módulos de un sistema
```

### Permisos de Módulos

```
GET    /api/v1/module-permissions      # Listar permisos de módulos
POST   /api/v1/module-permissions      # Crear permiso de módulo
PUT    /api/v1/module-permissions/:id  # Actualizar permiso
DELETE /api/v1/module-permissions/:id  # Eliminar permiso
GET    /api/v1/modules/:id/permissions # Permisos de un módulo
```

---

## GRUPOS

### Gestión de Grupos

```
GET    /api/v1/groups                  # Listar grupos
GET    /api/v1/groups/:id              # Obtener grupo por ID
POST   /api/v1/groups                  # Crear nuevo grupo
PUT    /api/v1/groups/:id              # Actualizar grupo
DELETE /api/v1/groups/:id              # Eliminación lógica
```

### Asignación Usuario-Grupos

```
GET    /api/v1/user-groups             # Listar asignaciones usuario-grupo
POST   /api/v1/user-groups             # Asignar usuario a grupo
DELETE /api/v1/user-groups/:id         # Eliminar asignación
GET    /api/v1/users/:id/groups        # Grupos de un usuario
GET    /api/v1/groups/:id/users        # Usuarios de un grupo
```

---

## TOKENS Y POLÍTICAS

### Tokens de API

```
GET    /api/v1/api-tokens              # Listar tokens del usuario
POST   /api/v1/api-tokens              # Crear nuevo token
PUT    /api/v1/api-tokens/:id          # Actualizar token
DELETE /api/v1/api-tokens/:id          # Eliminar token
```

### Tokens de Verificación

```
POST   /api/v1/verification-tokens     # Crear token de verificación
GET    /api/v1/verification-tokens/:token # Validar token
DELETE /api/v1/verification-tokens/:id # Invalidar token
```

### Políticas de Sesión

```
GET    /api/v1/session-policies        # Obtener políticas de sesión
PUT    /api/v1/session-policies/:id    # Actualizar política
GET    /api/v1/systems/:id/session-policy # Política de un sistema
```

---

## 📊 AUDITORÍA Y REPORTES

### Logs de Auditoría

```
GET    /api/v1/audit-logs              # Listar logs de auditoría
GET    /api/v1/audit-logs/:id          # Obtener log específico
GET    /api/v1/users/:id/audit-logs    # Logs de un usuario
GET    /api/v1/audit-logs/search       # Búsqueda avanzada de logs
```

### Reportes

```
GET    /api/v1/reports/users           # Reporte de usuarios
GET    /api/v1/reports/sessions        # Reporte de sesiones
GET    /api/v1/reports/movements       # Reporte de movimientos
GET    /api/v1/reports/permissions     # Reporte de permisos
GET    /api/v1/reports/activity        # Reporte de actividad
```

### Historial General

```
GET    /api/v1/history/users           # Historial de usuarios
GET    /api/v1/history/organic-units   # Historial de unidades
GET    /api/v1/history/positions       # Historial de posiciones
GET    /api/v1/history/systems         # Historial de sistemas
```

---

## 🔍 BÚSQUEDA Y FILTROS

### Búsquedas Generales

```
GET    /api/v1/search/users            # Búsqueda de usuarios
GET    /api/v1/search/organic-units    # Búsqueda de unidades
GET    /api/v1/search/positions        # Búsqueda de posiciones
GET    /api/v1/search/global           # Búsqueda global
```

---

## ⚙️ ADMINISTRACIÓN

### Sistema General

```
GET    /api/v1/admin/stats             # Estadísticas del sistema
GET    /api/v1/admin/health            # Estado de salud del sistema
POST   /api/v1/admin/maintenance       # Modo mantenimiento
GET    /api/v1/admin/backup            # Crear backup
POST   /api/v1/admin/restore           # Restaurar backup
```

### Configuración

```
GET    /api/v1/config                  # Obtener configuración
PUT    /api/v1/config                  # Actualizar configuración
GET    /api/v1/config/features         # Features habilitados
```

---

## 📋 PARÁMETROS COMUNES

### Paginación (Query Parameters)

```
?page=1&limit=20&sort=created_at&order=desc
```

### Filtros Comunes

```
?status=active&is_deleted=false&organic_unit_id=123
?search=keyword&date_from=2024-01-01&date_to=2024-12-31
```

### Includes (Relaciones)

```
?include=organic_unit,roles,positions
?include=history,audit_logs
```

---

## 🔐 MIDDLEWARES REQUERIDOS

- **Autenticación**: JWT Token validation
- **Autorización**: Role-based access control (RBAC)
- **Rate Limiting**: Control de límites por usuario/IP
- **Audit Logging**: Registro automático de todas las acciones
- **CORS**: Cross-origin resource sharing
- **Validation**: Validación de datos de entrada
- **Error Handling**: Manejo centralizado de errores

---

## 📝 CÓDIGOS DE RESPUESTA HTTP

- **200**: OK - Operación exitosa
- **201**: Created - Recurso creado exitosamente
- **204**: No Content - Eliminación exitosa
- **400**: Bad Request - Datos de entrada inválidos
- **401**: Unauthorized - Token inválido/expirado
- **403**: Forbidden - Sin permisos suficientes
- **404**: Not Found - Recurso no encontrado
- **409**: Conflict - Conflicto (ej: email duplicado)
- **422**: Unprocessable Entity - Errores de validación
- **429**: Too Many Requests - Rate limit excedido
- **500**: Internal Server Error - Error del servidor
