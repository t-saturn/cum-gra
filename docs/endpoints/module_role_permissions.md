**1. GET - Listar Module Role Permissions (sin filtros)**
```http
GET http://localhost:8080/api/module-role-permissions?page=1&page_size=10
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**2. GET - Listar Permisos por Módulo**
```http
GET http://localhost:8080/api/module-role-permissions?page=1&page_size=10&module_id=770e8400-e29b-41d4-a716-446655440003
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Listar Permisos por Rol**
```http
GET http://localhost:8080/api/module-role-permissions?page=1&page_size=10&role_id=660e8400-e29b-41d4-a716-446655440001
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. GET - Listar Permisos por Módulo Y Rol**
```http
GET http://localhost:8080/api/module-role-permissions?page=1&page_size=10&module_id=770e8400-e29b-41d4-a716-446655440003&role_id=660e8400-e29b-41d4-a716-446655440001
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**5. GET - Listar Permisos eliminados**
```http
GET http://localhost:8080/api/module-role-permissions?page=1&page_size=10&is_deleted=true
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**6. GET - Estadísticas de Permisos**
```http
GET http://localhost:8080/api/module-role-permissions/stats
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**7. GET - Obtener Permiso por ID**
```http
GET http://localhost:8080/api/module-role-permissions/990e8400-e29b-41d4-a716-446655440005
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**8. POST - Crear Permiso (tipo read)**
```http
POST http://localhost:8080/api/module-role-permissions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
  "permission_type": "read"
}
```

**9. POST - Crear Permiso (tipo write)**
```http
POST http://localhost:8080/api/module-role-permissions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
  "permission_type": "write"
}
```

**10. POST - Crear Permiso (tipo execute)**
```http
POST http://localhost:8080/api/module-role-permissions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
  "permission_type": "execute"
}
```

**11. POST - Crear Permiso (tipo delete)**
```http
POST http://localhost:8080/api/module-role-permissions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
  "permission_type": "delete"
}
```

**12. POST - Crear Permiso (tipo admin)**
```http
POST http://localhost:8080/api/module-role-permissions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
  "permission_type": "admin"
}
```

**13. POST - Asignar Permisos Masivamente**
```http
POST http://localhost:8080/api/module-role-permissions/bulk-assign
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
  "module_ids": [
    "770e8400-e29b-41d4-a716-446655440003",
    "880e8400-e29b-41d4-a716-446655440004",
    "990e8400-e29b-41d4-a716-446655440005"
  ],
  "permission_type": "read"
}
```

**14. POST - Asignar Permisos Masivamente (tipo write)**
```http
POST http://localhost:8080/api/module-role-permissions/bulk-assign
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
  "module_ids": [
    "770e8400-e29b-41d4-a716-446655440003",
    "880e8400-e29b-41d4-a716-446655440004"
  ],
  "permission_type": "write"
}
```

**15. PUT - Actualizar Permiso (cambiar tipo)**
```http
PUT http://localhost:8080/api/module-role-permissions/990e8400-e29b-41d4-a716-446655440005
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "permission_type": "admin"
}
```

**16. DELETE - Eliminar Permiso (Soft Delete)**
```http
DELETE http://localhost:8080/api/module-role-permissions/990e8400-e29b-41d4-a716-446655440005
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**17. PATCH - Restaurar Permiso**
```http
PATCH http://localhost:8080/api/module-role-permissions/990e8400-e29b-41d4-a716-446655440005/restore
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuestas esperadas:**

**GET /api/module-role-permissions (200 OK):**
```json
{
  "data": [
    {
      "id": "990e8400-e29b-41d4-a716-446655440005",
      "module_id": "770e8400-e29b-41d4-a716-446655440003",
      "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
      "permission_type": "read",
      "created_at": "2024-12-12T10:00:00Z",
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "module_name": "Dashboard Principal",
      "module_route": "/dashboard",
      "role_name": "Administrator",
      "application_name": "Mi Aplicación",
      "application_client_id": "mi-app-client"
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10
}
```

**GET /api/module-role-permissions/stats (200 OK):**
```json
{
  "total_permissions": 150,
  "active_permissions": 145,
  "deleted_permissions": 5,
  "unique_modules": 25,
  "unique_roles": 12,
  "permissions_by_type": {
    "read": 50,
    "write": 40,
    "execute": 30,
    "delete": 15,
    "admin": 10
  }
}
```

**POST /api/module-role-permissions (201 Created):**
```json
{
  "id": "aa0e8400-e29b-41d4-a716-446655440006",
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
  "permission_type": "read",
  "created_at": "2024-12-12T15:30:00Z",
  "is_deleted": false,
  "deleted_at": null,
  "deleted_by": null,
  "module_name": "Dashboard Principal",
  "module_route": "/dashboard",
  "role_name": "Administrator",
  "application_name": "Mi Aplicación",
  "application_client_id": "mi-app-client"
}
```

**POST /api/module-role-permissions/bulk-assign (201 Created):**
```json
{
  "created": 2,
  "skipped": 1,
  "failed": 0,
  "details": [
    {
      "id": "bb0e8400-e29b-41d4-a716-446655440007",
      "module_id": "770e8400-e29b-41d4-a716-446655440003",
      "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
      "permission_type": "read",
      "created_at": "2024-12-12T15:35:00Z",
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "module_name": "Dashboard Principal",
      "module_route": "/dashboard",
      "role_name": "Administrator",
      "application_name": "Mi Aplicación",
      "application_client_id": "mi-app-client"
    },
    {
      "id": "cc0e8400-e29b-41d4-a716-446655440008",
      "module_id": "880e8400-e29b-41d4-a716-446655440004",
      "application_role_id": "660e8400-e29b-41d4-a716-446655440001",
      "permission_type": "read",
      "created_at": "2024-12-12T15:35:01Z",
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "module_name": "Reportes",
      "module_route": "/reports",
      "role_name": "Administrator",
      "application_name": "Mi Aplicación",
      "application_client_id": "mi-app-client"
    }
  ]
}
```

**DELETE /api/module-role-permissions/:id (200 OK):**
```json
{
  "message": "Permiso eliminado correctamente"
}
```

**PATCH /api/module-role-permissions/:id/restore (200 OK):**
```json
{
  "message": "Permiso restaurado correctamente"
}
```

**Errores comunes:**

**400 Bad Request (módulo y rol de diferentes aplicaciones):**
```json
{
  "error": "el módulo y el rol deben pertenecer a la misma aplicación"
}
```

**404 Not Found:**
```json
{
  "error": "permiso no encontrado"
}
```

**404 Not Found (módulo no existe):**
```json
{
  "error": "módulo no encontrado"
}
```

**404 Not Found (rol no existe):**
```json
{
  "error": "rol no encontrado"
}
```

**409 Conflict (permiso duplicado):**
```json
{
  "error": "ya existe un permiso para este módulo y rol"
}
```