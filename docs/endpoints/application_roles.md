**1. GET - Listar Application Roles (sin filtros)**
```http
GET http://localhost:8080/api/application-roles?page=1&page_size=10
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**2. GET - Listar Application Roles (filtrado por aplicación)**
```http
GET http://localhost:8080/api/application-roles?page=1&page_size=10&application_id=550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Listar Application Roles eliminados**
```http
GET http://localhost:8080/api/application-roles?page=1&page_size=10&is_deleted=true
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. GET - Estadísticas de Application Roles**
```http
GET http://localhost:8080/api/application-roles/stats
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**5. GET - Obtener Application Role por ID**
```http
GET http://localhost:8080/api/application-roles/660e8400-e29b-41d4-a716-446655440001
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**6. POST - Crear Application Role**
```http
POST http://localhost:8080/api/application-roles
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Administrator",
  "description": "Administrador con acceso completo al sistema",
  "application_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

**7. POST - Crear Application Role (sin descripción)**
```http
POST http://localhost:8080/api/application-roles
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Viewer",
  "application_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

**8. PUT - Actualizar Application Role (completo)**
```http
PUT http://localhost:8080/api/application-roles/660e8400-e29b-41d4-a716-446655440001
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Super Administrator",
  "description": "Administrador con permisos especiales y acceso total"
}
```

**9. PUT - Actualizar Application Role (solo nombre)**
```http
PUT http://localhost:8080/api/application-roles/660e8400-e29b-41d4-a716-446655440001
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Editor"
}
```

**10. PUT - Actualizar Application Role (solo descripción)**
```http
PUT http://localhost:8080/api/application-roles/660e8400-e29b-41d4-a716-446655440001
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "description": "Rol con permisos de edición de contenido"
}
```

**11. DELETE - Eliminar Application Role (Soft Delete)**
```http
DELETE http://localhost:8080/api/application-roles/660e8400-e29b-41d4-a716-446655440001
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**12. PATCH - Restaurar Application Role**
```http
PATCH http://localhost:8080/api/application-roles/660e8400-e29b-41d4-a716-446655440001/restore
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuestas esperadas:**

**GET /api/application-roles (200 OK):**
```json
{
  "data": [
    {
      "id": "660e8400-e29b-41d4-a716-446655440001",
      "name": "Administrator",
      "description": "Administrador con acceso completo",
      "application_id": "550e8400-e29b-41d4-a716-446655440000",
      "application": {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "name": "Mi Aplicación",
        "client_id": "mi-app-client"
      },
      "created_at": "2024-12-12T10:00:00Z",
      "updated_at": "2024-12-12T10:00:00Z",
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "modules_count": 5,
      "users_count": 12
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10
}
```

**GET /api/application-roles/stats (200 OK):**
```json
{
  "total_roles": 25,
  "active_roles": 23,
  "deleted_roles": 2,
  "roles_with_modules": 18,
  "roles_with_users": 20
}
```

**POST /api/application-roles (201 Created):**
```json
{
  "id": "770e8400-e29b-41d4-a716-446655440002",
  "name": "Administrator",
  "description": "Administrador con acceso completo al sistema",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "application": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Mi Aplicación",
    "client_id": "mi-app-client"
  },
  "created_at": "2024-12-12T15:30:00Z",
  "updated_at": "2024-12-12T15:30:00Z",
  "is_deleted": false,
  "deleted_at": null,
  "deleted_by": null,
  "modules_count": 0,
  "users_count": 0
}
```

**DELETE /api/application-roles/:id (200 OK):**
```json
{
  "message": "Rol eliminado correctamente"
}
```

**PATCH /api/application-roles/:id/restore (200 OK):**
```json
{
  "message": "Rol restaurado correctamente"
}
```

**Errores comunes:**

**400 Bad Request (application_id inválido):**
```json
{
  "error": "application_id inválido"
}
```

**404 Not Found:**
```json
{
  "error": "rol no encontrado"
}
```

**409 Conflict (nombre duplicado):**
```json
{
  "error": "ya existe un rol con este nombre en esta aplicación"
}
```