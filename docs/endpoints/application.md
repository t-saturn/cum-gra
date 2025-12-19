**1. GET - Listar Aplicaciones (con paginación)**
```http
GET http://localhost:8000/api/applications?page=1&page_size=10&is_deleted=false
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**2. GET - Obtener Aplicación por ID**
```http
GET http://localhost:8000/api/applications/550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Estadísticas de Aplicaciones**
```http
GET http://localhost:8000/api/applications/stats
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. POST - Crear Aplicación**
```http
POST http://localhost:8000/api/applications
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Mi Nueva Aplicación",
  "client_id": "mi-app-client-2024",
  "client_secret": "super-secret-password-123",
  "domain": "https://mi-app.regionayacucho.gob.pe",
  "logo": "https://mi-app.regionayacucho.gob.pe/logo.png",
  "description": "Esta es una aplicación de prueba para el sistema",
  "status": "active"
}
```

**5. PUT - Actualizar Aplicación**
```http
PUT http://localhost:8000/api/applications/550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Aplicación Actualizada",
  "description": "Descripción actualizada de la aplicación",
  "status": "inactive"
}
```

**Actualización parcial (solo un campo):**
```http
PUT http://localhost:8000/api/applications/550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "status": "active"
}
```

**6. DELETE - Eliminar Aplicación (Soft Delete)**
```http
DELETE http://localhost:8000/api/applications/550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**7. PATCH - Restaurar Aplicación Eliminada**
```http
PATCH http://localhost:8000/api/applications/550e8400-e29b-41d4-a716-446655440000/restore
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**8. GET - Listar Aplicaciones Eliminadas**
```http
GET http://localhost:8000/api/applications?page=1&page_size=10&is_deleted=true
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuestas esperadas:**

**GET /applications (200 OK):**
```json
{
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Mi Aplicación",
      "client_id": "mi-app-client",
      "domain": "https://mi-app.regionayacucho.gob.pe",
      "logo": "https://mi-app.regionayacucho.gob.pe/logo.png",
      "description": "Descripción de la app",
      "status": "active",
      "created_at": "2024-12-12T10:00:00Z",
      "updated_at": "2024-12-12T10:00:00Z",
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "admins": [
        {
          "full_name": "Admin User",
          "dni": "12345678",
          "email": "admin@regionayacucho.gob.pe"
        }
      ],
      "users_count": 25
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10
}
```

**POST /applications (201 Created):**
```json
{
  "id": "660e8400-e29b-41d4-a716-446655440001",
  "name": "Mi Nueva Aplicación",
  "client_id": "mi-app-client-2024",
  "domain": "https://mi-app.regionayacucho.gob.pe",
  "logo": "https://mi-app.regionayacucho.gob.pe/logo.png",
  "description": "Esta es una aplicación de prueba",
  "status": "active",
  "created_at": "2024-12-12T15:30:00Z",
  "updated_at": "2024-12-12T15:30:00Z",
  "is_deleted": false,
  "deleted_at": null,
  "deleted_by": null,
  "admins": [],
  "users_count": 0
}
```

**Errores comunes:**

**401 Unauthorized (sin token o token inválido):**
```json
{
  "error": "Token de autorización requerido"
}
```

**403 Forbidden (sin permisos):**
```json
{
  "error": "Rol requerido en realm-management: manage-clients"
}
```

**404 Not Found:**
```json
{
  "error": "aplicación no encontrada"
}
```

**409 Conflict (client_id duplicado):**
```json
{
  "error": "ya existe una aplicación con este client_id"
}
```