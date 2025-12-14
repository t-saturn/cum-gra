**1. GET - Listar Módulos (sin filtros)**
```http
GET http://localhost:8080/api/modules?page=1&page_size=10
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**2. GET - Listar Módulos (filtrado por aplicación)**
```http
GET http://localhost:8080/api/modules?page=1&page_size=10&application_id=550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Listar Módulos eliminados**
```http
GET http://localhost:8080/api/modules?page=1&page_size=10&is_deleted=true
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. GET - Estadísticas de Módulos**
```http
GET http://localhost:8080/api/modules/stats
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**5. GET - Obtener Módulo por ID**
```http
GET http://localhost:8080/api/modules/770e8400-e29b-41d4-a716-446655440003
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**6. POST - Crear Módulo (completo con aplicación y padre)**
```http
POST http://localhost:8080/api/modules
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "item": "dashboard",
  "name": "Dashboard Principal",
  "route": "/dashboard",
  "icon": "dashboard-icon",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "parent_id": null,
  "sort_order": 1,
  "status": "active"
}
```

**7. POST - Crear Submódulo (con padre)**
```http
POST http://localhost:8080/api/modules
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Reportes de Ventas",
  "route": "/dashboard/sales-reports",
  "icon": "report-icon",
  "parent_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "sort_order": 1,
  "status": "active"
}
```

**8. POST - Crear Módulo simple (mínimo requerido)**
```http
POST http://localhost:8080/api/modules
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Configuración",
  "route": "/settings"
}
```

**9. POST - Crear Módulo sin aplicación (módulo global)**
```http
POST http://localhost:8080/api/modules
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Perfil de Usuario",
  "route": "/profile",
  "icon": "user-icon",
  "sort_order": 100,
  "status": "active"
}
```

**10. PUT - Actualizar Módulo (completo)**
```http
PUT http://localhost:8080/api/modules/770e8400-e29b-41d4-a716-446655440003
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "item": "dashboard-updated",
  "name": "Dashboard Actualizado",
  "route": "/dashboard/v2",
  "icon": "new-dashboard-icon",
  "sort_order": 2,
  "status": "active"
}
```

**11. PUT - Actualizar Módulo (solo nombre)**
```http
PUT http://localhost:8080/api/modules/770e8400-e29b-41d4-a716-446655440003
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Panel de Control"
}
```

**12. PUT - Actualizar Módulo (cambiar padre)**
```http
PUT http://localhost:8080/api/modules/770e8400-e29b-41d4-a716-446655440003
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "parent_id": "880e8400-e29b-41d4-a716-446655440004"
}
```

**13. PUT - Actualizar Módulo (quitar padre - convertir en raíz)**
```http
PUT http://localhost:8080/api/modules/770e8400-e29b-41d4-a716-446655440003
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "parent_id": ""
}
```

**14. PUT - Actualizar Módulo (cambiar orden y estado)**
```http
PUT http://localhost:8080/api/modules/770e8400-e29b-41d4-a716-446655440003
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "sort_order": 5,
  "status": "inactive"
}
```

**15. DELETE - Eliminar Módulo (Soft Delete)**
```http
DELETE http://localhost:8080/api/modules/770e8400-e29b-41d4-a716-446655440003
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**16. PATCH - Restaurar Módulo**
```http
PATCH http://localhost:8080/api/modules/770e8400-e29b-41d4-a716-446655440003/restore
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuestas esperadas:**

**GET /api/modules (200 OK):**
```json
{
  "data": [
    {
      "id": "770e8400-e29b-41d4-a716-446655440003",
      "item": "dashboard",
      "name": "Dashboard Principal",
      "route": "/dashboard",
      "icon": "dashboard-icon",
      "parent_id": null,
      "application_id": "550e8400-e29b-41d4-a716-446655440000",
      "sort_order": 1,
      "status": "active",
      "created_at": "2024-12-12T10:00:00Z",
      "updated_at": "2024-12-12T10:00:00Z",
      "deleted_at": null,
      "deleted_by": null,
      "application_name": "Mi Aplicación",
      "application_client_id": "mi-app-client",
      "users_count": 25
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10
}
```

**GET /api/modules/stats (200 OK):**
```json
{
  "total_modules": 45,
  "active_modules": 40,
  "deleted_modules": 5,
  "total_users": 150
}
```

**POST /api/modules (201 Created):**
```json
{
  "id": "880e8400-e29b-41d4-a716-446655440004",
  "item": "dashboard",
  "name": "Dashboard Principal",
  "route": "/dashboard",
  "icon": "dashboard-icon",
  "parent_id": null,
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "sort_order": 1,
  "status": "active",
  "created_at": "2024-12-12T15:30:00Z",
  "updated_at": "2024-12-12T15:30:00Z",
  "deleted_at": null,
  "deleted_by": null,
  "application_name": "Mi Aplicación",
  "application_client_id": "mi-app-client",
  "users_count": 0
}
```

**DELETE /api/modules/:id (200 OK):**
```json
{
  "message": "Módulo eliminado correctamente"
}
```

**PATCH /api/modules/:id/restore (200 OK):**
```json
{
  "message": "Módulo restaurado correctamente"
}
```

**Errores comunes:**

**400 Bad Request (módulo con hijos):**
```json
{
  "error": "no se puede eliminar un módulo que tiene submódulos"
}
```

**400 Bad Request (padre inválido):**
```json
{
  "error": "un módulo no puede ser su propio padre"
}
```

**404 Not Found:**
```json
{
  "error": "módulo no encontrado"
}
```

**404 Not Found (padre no existe):**
```json
{
  "error": "módulo padre no encontrado"
}
```

**409 Conflict (nombre duplicado):**
```json
{
  "error": "ya existe un módulo con este nombre en esta aplicación"
}
```