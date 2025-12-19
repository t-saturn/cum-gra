**1. GET - Listar Restricciones de Usuarios (sin filtros)**
```http
GET http://localhost:8080/api/user-restrictions?page=1&page_size=10
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**2. GET - Listar Restricciones por Usuario**
```http
GET http://localhost:8080/api/user-restrictions?page=1&page_size=10&user_id=a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Listar Restricciones por Aplicación**
```http
GET http://localhost:8080/api/user-restrictions?page=1&page_size=10&application_id=550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. GET - Listar Restricciones por Usuario y Aplicación**
```http
GET http://localhost:8080/api/user-restrictions?page=1&page_size=10&user_id=a54e954d-0c61-4861-ab94-d49eaf516672&application_id=550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**5. GET - Listar Restricciones eliminadas**
```http
GET http://localhost:8080/api/user-restrictions?page=1&page_size=10&is_deleted=true
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**6. GET - Estadísticas de Restricciones**
```http
GET http://localhost:8080/api/user-restrictions/stats
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**7. GET - Obtener Restricción por ID**
```http
GET http://localhost:8080/api/user-restrictions/dd0e8400-e29b-41d4-a716-446655440009
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**8. POST - Crear Restricción (tipo block - bloqueo total)**
```http
POST http://localhost:8080/api/user-restrictions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "restriction_type": "block",
  "reason": "Usuario suspendido temporalmente por violación de políticas"
}
```

**9. POST - Crear Restricción (tipo limit - limitar permisos)**
```http
POST http://localhost:8080/api/user-restrictions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "restriction_type": "limit",
  "max_permission_level": "read",
  "reason": "Usuario en periodo de prueba, solo lectura permitida"
}
```

**10. POST - Crear Restricción (tipo read_only - solo lectura)**
```http
POST http://localhost:8080/api/user-restrictions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "restriction_type": "read_only",
  "reason": "Auditoría en curso, permisos de escritura restringidos"
}
```

**11. POST - Crear Restricción con Fecha de Expiración**
```http
POST http://localhost:8080/api/user-restrictions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "restriction_type": "block",
  "reason": "Suspensión temporal de 7 días",
  "expires_at": "2025-12-20T23:59:59Z"
}
```

**12. POST - Crear Restricción con Límite de Permisos (write)**
```http
POST http://localhost:8080/api/user-restrictions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "restriction_type": "limit",
  "max_permission_level": "write",
  "reason": "Usuario junior, sin permisos de eliminación"
}
```

**13. POST - Crear Restricciones Masivas**
```http
POST http://localhost:8080/api/user-restrictions/bulk-create
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "module_ids": [
    "770e8400-e29b-41d4-a716-446655440003",
    "880e8400-e29b-41d4-a716-446655440004",
    "990e8400-e29b-41d4-a716-446655440005"
  ],
  "restriction_type": "read_only",
  "reason": "Usuario en capacitación, acceso limitado a módulos críticos"
}
```

**14. POST - Restricción Masiva con Bloqueo**
```http
POST http://localhost:8080/api/user-restrictions/bulk-create
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "module_ids": [
    "770e8400-e29b-41d4-a716-446655440003",
    "880e8400-e29b-41d4-a716-446655440004"
  ],
  "restriction_type": "block",
  "reason": "Investigación de seguridad en curso"
}
```

**15. PUT - Actualizar Restricción (cambiar tipo)**
```http
PUT http://localhost:8080/api/user-restrictions/dd0e8400-e29b-41d4-a716-446655440009
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "restriction_type": "read_only"
}
```

**16. PUT - Actualizar Restricción (cambiar nivel de permisos)**
```http
PUT http://localhost:8080/api/user-restrictions/dd0e8400-e29b-41d4-a716-446655440009
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "restriction_type": "limit",
  "max_permission_level": "execute"
}
```

**17. PUT - Actualizar Restricción (agregar fecha de expiración)**
```http
PUT http://localhost:8080/api/user-restrictions/dd0e8400-e29b-41d4-a716-446655440009
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "expires_at": "2025-12-31T23:59:59Z",
  "reason": "Restricción temporal extendida hasta fin de año"
}
```

**18. PUT - Actualizar Restricción (quitar expiración)**
```http
PUT http://localhost:8080/api/user-restrictions/dd0e8400-e29b-41d4-a716-446655440009
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "expires_at": "",
  "reason": "Restricción convertida a permanente"
}
```

**19. PUT - Actualizar solo la razón**
```http
PUT http://localhost:8080/api/user-restrictions/dd0e8400-e29b-41d4-a716-446655440009
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "reason": "Actualización: Usuario completó capacitación pero aún requiere supervisión"
}
```

**20. DELETE - Eliminar Restricción (Soft Delete)**
```http
DELETE http://localhost:8080/api/user-restrictions/dd0e8400-e29b-41d4-a716-446655440009
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**21. PATCH - Restaurar Restricción**
```http
PATCH http://localhost:8080/api/user-restrictions/dd0e8400-e29b-41d4-a716-446655440009/restore
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuestas esperadas:**

**GET /api/user-restrictions (200 OK):**
```json
{
  "data": [
    {
      "id": "dd0e8400-e29b-41d4-a716-446655440009",
      "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
      "module_id": "770e8400-e29b-41d4-a716-446655440003",
      "application_id": "550e8400-e29b-41d4-a716-446655440000",
      "restriction_type": "block",
      "max_permission_level": null,
      "reason": "Usuario suspendido temporalmente por violación de políticas",
      "expires_at": "2025-12-20T23:59:59Z",
      "created_at": "2024-12-12T10:00:00Z",
      "created_by": "b54e954d-0c61-4861-ab94-d49eaf516673",
      "updated_at": "2024-12-12T10:00:00Z",
      "updated_by": null,
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "user_email": "usuario@regionayacucho.gob.pe",
      "user_full_name": "Juan Pérez",
      "module_name": "Dashboard Principal",
      "module_route": "/dashboard",
      "application_name": "Mi Aplicación",
      "application_client_id": "mi-app-client"
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10
}
```

**GET /api/user-restrictions/stats (200 OK):**
```json
{
  "total_restrictions": 150,
  "active_restrictions": 120,
  "restricted_users": 45,
  "deleted_restrictions": 30
}
```

**POST /api/user-restrictions (201 Created):**
```json
{
  "id": "ee0e8400-e29b-41d4-a716-446655440010",
  "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "module_id": "770e8400-e29b-41d4-a716-446655440003",
  "application_id": "550e8400-e29b-41d4-a716-446655440000",
  "restriction_type": "block",
  "max_permission_level": null,
  "reason": "Usuario suspendido temporalmente por violación de políticas",
  "expires_at": null,
  "created_at": "2024-12-12T15:30:00Z",
  "created_by": "b54e954d-0c61-4861-ab94-d49eaf516673",
  "updated_at": "2024-12-12T15:30:00Z",
  "updated_by": null,
  "is_deleted": false,
  "deleted_at": null,
  "deleted_by": null,
  "user_email": "usuario@regionayacucho.gob.pe",
  "user_full_name": "Juan Pérez",
  "module_name": "Dashboard Principal",
  "module_route": "/dashboard",
  "application_name": "Mi Aplicación",
  "application_client_id": "mi-app-client"
}
```

**POST /api/user-restrictions/bulk-create (201 Created):**
```json
{
  "created": 2,
  "skipped": 1,
  "failed": 0,
  "details": [
    {
      "id": "ff0e8400-e29b-41d4-a716-446655440011",
      "user_id": "a54e954d-0c61-4861-ab94-d49eaf516672",
      "module_id": "770e8400-e29b-41d4-a716-446655440003",
      "application_id": "550e8400-e29b-41d4-a716-446655440000",
      "restriction_type": "read_only",
      "max_permission_level": null,
      "reason": "Usuario en capacitación",
      "expires_at": null,
      "created_at": "2024-12-12T15:35:00Z",
      "created_by": "b54e954d-0c61-4861-ab94-d49eaf516673",
      "updated_at": "2024-12-12T15:35:00Z",
      "updated_by": null,
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "user_email": "usuario@regionayacucho.gob.pe",
      "user_full_name": "Juan Pérez",
      "module_name": "Dashboard Principal",
      "module_route": "/dashboard",
      "application_name": "Mi Aplicación",
      "application_client_id": "mi-app-client"
    }
  ]
}
```

**DELETE /api/user-restrictions/:id (200 OK):**
```json
{
  "message": "Restricción eliminada correctamente"
}
```

**PATCH /api/user-restrictions/:id/restore (200 OK):**
```json
{
  "message": "Restricción restaurada correctamente"
}
```

**Errores comunes:**

**400 Bad Request (módulo no pertenece a la aplicación):**
```json
{
  "error": "el módulo no pertenece a la aplicación especificada"
}
```

**404 Not Found:**
```json
{
  "error": "restricción no encontrada"
}
```

**409 Conflict (restricción duplicada):**
```json
{
  "error": "ya existe una restricción activa para este usuario y módulo"
}
```

**Notas sobre tipos de restricciones:**
- **block**: Bloqueo total del acceso al módulo
- **limit**: Limita el nivel máximo de permisos (requiere `max_permission_level`)
- **read_only**: Solo permite lectura, sin permisos de escritura/modificación

**Niveles de permisos disponibles:**
- **read**: Solo lectura
- **write**: Lectura y escritura
- **execute**: Lectura, escritura y ejecución
- **delete**: Todos los anteriores más eliminación
- **admin**: Control total