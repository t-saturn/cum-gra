**1. GET - Listar Asignaciones de Roles (sin filtros)**
```http
GET http://localhost:8080/api/user-application-roles?page=1&page_size=10
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**2. GET - Listar Asignaciones por Usuario**
```http
GET http://localhost:8080/api/user-application-roles?page=1&page_size=10&user_id=a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Listar Asignaciones por Aplicación**
```http
GET http://localhost:8080/api/user-application-roles?page=1&page_size=10&application_id=550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. GET - Listar Asignaciones Activas (no revocadas)**
```http
GET http://localhost:8080/api/user-application-roles?page=1&page_size=10&is_revoked=false
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**5. GET - Listar Asignaciones Revocadas**
```http
GET http://localhost:8080/api/user-application-roles?page=1&page_size=10&is_revoked=true
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**6. GET - Listar Asignaciones por Usuario y Aplicación**
```http
GET http://localhost:8080/api/user-application-roles?page=1&page_size=10&user_id=a54e954d-0c61-4861-ab94-d49eaf516672&application_id=550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**7. GET - Listar Asignaciones eliminadas**
```http
GET http://localhost:8080/api/user-application-roles?page=1&page_size=10&is_deleted=true
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**8. GET - Estadísticas de Asignaciones**
```http
GET http://localhost:8080/api/user-application-roles/stats
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**9. GET - Obtener Asignación por ID**
```http
GET http://localhost:8080/api/user-application-roles/cc0e8400-e29b-41d4-a716-446655440012
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**10. POST - Asignar Rol a Usuario**
```http
POST http://localhost:8080/api/user-application-roles
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001"
}
```

**11. POST - Asignar Rol de Administrador**
```http
POST http://localhost:8080/api/user-application-roles
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "application_role_id": "770e8400-e29b-41d4-a716-446655440002"
}
```

**12. POST - Asignar múltiples roles a un usuario**
```http
POST http://localhost:8080/api/user-application-roles/bulk-assign-roles-to-user
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "role_ids": [
    "660e8400-e29b-41d4-a716-446655440001",
    "770e8400-e29b-41d4-a716-446655440002",
    "880e8400-e29b-41d4-a716-446655440003"
  ]
}
```

**13. POST - Asignar rol a múltiples usuarios**
```http
POST http://localhost:8080/api/user-application-roles/bulk-assign-role-to-users
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_ids": [
    "a54e954d-0c61-4861-ab94-d49eaf516672",
    "b54e954d-0c61-4861-ab94-d49eaf516673",
    "c54e954d-0c61-4861-ab94-d49eaf516674"
  ],
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001"
}
```

**14. POST - Asignar rol de "Viewer" a múltiples usuarios**
```http
POST http://localhost:8080/api/user-application-roles/bulk-assign-role-to-users
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_ids": [
    "a54e954d-0c61-4861-ab94-d49eaf516672",
    "b54e954d-0c61-4861-ab94-d49eaf516673"
  ],
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "application_role_id": "990e8400-e29b-41d4-a716-446655440004"
}
```

**15. PATCH - Revocar Rol**
```http
PATCH http://localhost:8080/api/user-application-roles/cc0e8400-e29b-41d4-a716-446655440012/revoke
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "reason": "Usuario cambió de departamento"
}
```

**16. PATCH - Revocar Rol (sin razón)**
```http
PATCH http://localhost:8080/api/user-application-roles/cc0e8400-e29b-41d4-a716-446655440012/revoke
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{}
```

**17. PATCH - Restaurar Rol Revocado**
```http
PATCH http://localhost:8080/api/user-application-roles/cc0e8400-e29b-41d4-a716-446655440012/restore
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**18. DELETE - Eliminar Asignación (Soft Delete)**
```http
DELETE http://localhost:8080/api/user-application-roles/cc0e8400-e29b-41d4-a716-446655440012
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**19. PATCH - Recuperar Asignación Eliminada**
```http
PATCH http://localhost:8080/api/user-application-roles/cc0e8400-e29b-41d4-a716-446655440012/undelete
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuestas esperadas:**

**GET /api/user-application-roles (200 OK):**
```json
{
  "data": [
    {
      "id": "cc0e8400-e29b-41d4-a716-446655440012",
      "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
      "application_id": "550e8400-e29b-41d4-a716-446655440000",
      "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
      "granted_at": "2024-12-12T10:00:00Z",
      "granted_by": "b54e954d-0c61-4861-ab94-d49eaf516673",
      "revoked_at": null,
      "revoked_by": null,
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "created_at": "2024-12-12T10:00:00Z",
      "updated_at": "2024-12-12T10:00:00Z",
      "user_email": "usuario@regionayacucho.gob.pe",
      "user_full_name": "Juan Pérez",
      "application_name": "Mi Aplicación",
      "application_client_id": "mi-app-client",
      "role_name": "Administrator",
      "granted_by_email": "admin@regionayacucho.gob.pe",
      "revoked_by_email": null
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10
}
```

**GET /api/user-application-roles/stats (200 OK):**
```json
{
  "total_assignments": 500,
  "active_assignments": 450,
  "revoked_assignments": 30,
  "deleted_assignments": 20,
  "users_with_roles": 120
}
```

**POST /api/user-application-roles (201 Created):**
```json
{
  "id": "dd0e8400-e29b-41d4-a716-446655440013",
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
  "granted_at": "2024-12-12T15:30:00Z",
  "granted_by": "b54e954d-0c61-4861-ab94-d49eaf516673",
  "revoked_at": null,
  "revoked_by": null,
  "is_deleted": false,
  "deleted_at": null,
  "deleted_by": null,
  "created_at": "2024-12-12T15:30:00Z",
  "updated_at": "2024-12-12T15:30:00Z",
  "user_email": "usuario@regionayacucho.gob.pe",
  "user_full_name": "Juan Pérez",
  "application_name": "Mi Aplicación",
  "application_client_id": "mi-app-client",
  "role_name": "Administrator",
  "granted_by_email": "admin@regionayacucho.gob.pe",
  "revoked_by_email": null
}
```

**POST /api/user-application-roles/bulk-assign-roles-to-user (201 Created):**
```json
{
  "created": 2,
  "skipped": 1,
  "failed": 0,
  "details": [
    {
      "id": "ee0e8400-e29b-41d4-a716-446655440014",
      "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
      "application_id": "550e8400-e29b-41d4-a716-446655440000",
      "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
      "granted_at": "2024-12-12T15:35:00Z",
      "granted_by": "b54e954d-0c61-4861-ab94-d49eaf516673",
      "revoked_at": null,
      "revoked_by": null,
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "created_at": "2024-12-12T15:35:00Z",
      "updated_at": "2024-12-12T15:35:00Z",
      "user_email": "usuario@regionayacucho.gob.pe",
      "user_full_name": "Juan Pérez",
      "application_name": "Mi Aplicación",
      "application_client_id": "mi-app-client",
      "role_name": "Administrator",
      "granted_by_email": "admin@regionayacucho.gob.pe",
      "revoked_by_email": null
    }
  ]
}
```

**POST /api/user-application-roles/bulk-assign-role-to-users (201 Created):**
```json
{
  "created": 3,
  "skipped": 0,
  "failed": 0,
  "details": [
    {
      "id": "ff0e8400-e29b-41d4-a716-446655440015",
      "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
      "application_id": "550e8400-e29b-41d4-a716-446655440000",
      "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
      "granted_at": "2024-12-12T15:40:00Z",
      "granted_by": "b54e954d-0c61-4861-ab94-d49eaf516673",
      "revoked_at": null,
      "revoked_by": null,
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "created_at": "2024-12-12T15:40:00Z",
      "updated_at": "2024-12-12T15:40:00Z",
      "user_email": "usuario1@regionayacucho.gob.pe",
      "user_full_name": "Juan Pérez",
      "application_name": "Mi Aplicación",
      "application_client_id": "mi-app-client",
      "role_name": "Administrator",
      "granted_by_email": "admin@regionayacucho.gob.pe",
      "revoked_by_email": null
    }
  ]
}
```

**PATCH /api/user-application-roles/:id/revoke (200 OK):**
```json
{
  "message": "Rol revocado correctamente"
}
```

**PATCH /api/user-application-roles/:id/restore (200 OK):**
```json
{
  "message": "Rol restaurado correctamente"
}
```

**DELETE /api/user-application-roles/:id (200 OK):**
```json
{
  "message": "Asignación de rol eliminada correctamente"
}
```

**PATCH /api/user-application-roles/:id/undelete (200 OK):**
```json
{
  "message": "Asignación de rol recuperada correctamente"
}
```

**Errores comunes:**

**400 Bad Request (rol no pertenece a la aplicación):**
```json
{
  "error": "el rol no pertenece a la aplicación especificada"
}
```

**400 Bad Request (rol ya revocado):**
```json
{
  "error": "este rol ya fue revocado anteriormente"
}
```

**400 Bad Request (rol no revocado):**
```json
{
  "error": "este rol no ha sido revocado"
}
```

**404 Not Found:**
```json
{
  "error": "asignación de rol no encontrada"
}
```

**409 Conflict (asignación duplicada):**
```json
{
  "error": "el usuario ya tiene este rol asignado en esta aplicación"
}
```