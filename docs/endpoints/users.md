**1. GET - Listar Usuarios (sin filtros)**
```http
GET http://localhost:8080/api/users?page=1&page_size=10
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**2. GET - Listar Usuarios Activos**
```http
GET http://localhost:8080/api/users?page=1&page_size=10&status=active
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Listar Usuarios Suspendidos**
```http
GET http://localhost:8080/api/users?page=1&page_size=10&status=suspended
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. GET - Listar Usuarios Inactivos**
```http
GET http://localhost:8080/api/users?page=1&page_size=10&status=inactive
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**5. GET - Listar Usuarios por Unidad Orgánica**
```http
GET http://localhost:8080/api/users?page=1&page_size=10&organic_unit_id=5
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**6. GET - Listar Usuarios por Posición Estructural**
```http
GET http://localhost:8080/api/users?page=1&page_size=10&position_id=3
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**7. GET - Listar Usuarios por Unidad Orgánica y Posición**
```http
GET http://localhost:8080/api/users?page=1&page_size=10&organic_unit_id=5&position_id=3
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**8. GET - Listar Usuarios eliminados**
```http
GET http://localhost:8080/api/users?page=1&page_size=10&is_deleted=true
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**9. GET - Estadísticas de Usuarios**
```http
GET http://localhost:8080/api/users/stats
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**10. GET - Obtener Usuario por ID**
```http
GET http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**11. POST - Crear Usuario (completo)**
```http
POST http://localhost:8080/api/users
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "email": "juan.perez@regionayacucho.gob.pe",
  "dni": "12345678",
  "first_name": "Juan",
  "last_name": "Pérez García",
  "phone": "+51987654321",
  "status": "active",
  "cod_emp_sgd": "EMP01",
  "structural_position_id": "3",
  "organic_unit_id": "5",
  "ubigeo_id": "1"
}
```

**12. POST - Crear Usuario (mínimo requerido)**
```http
POST http://localhost:8080/api/users
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "email": "maria.lopez@regionayacucho.gob.pe",
  "dni": "87654321",
  "first_name": "María",
  "last_name": "López Quispe"
}
```

**13. POST - Crear Usuario con Posición y Unidad**
```http
POST http://localhost:8080/api/users
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "email": "carlos.garcia@regionayacucho.gob.pe",
  "dni": "11223344",
  "first_name": "Carlos",
  "last_name": "García Rojas",
  "phone": "+51912345678",
  "structural_position_id": "2",
  "organic_unit_id": "3"
}
```

**14. POST - Crear Usuario Suspendido**
```http
POST http://localhost:8080/api/users
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "email": "temporal@regionayacucho.gob.pe",
  "dni": "55667788",
  "first_name": "Temporal",
  "last_name": "Usuario",
  "status": "suspended"
}
```

**15. PUT - Actualizar Usuario (completo)**
```http
PUT http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "email": "juan.perez.updated@regionayacucho.gob.pe",
  "first_name": "Juan Carlos",
  "last_name": "Pérez García",
  "phone": "+51999888777",
  "status": "active",
  "structural_position_id": "5",
  "organic_unit_id": "8",
  "cod_emp_sgd": "EMP02"
}
```

**16. PUT - Actualizar solo Email**
```http
PUT http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "email": "nuevo.email@regionayacucho.gob.pe"
}
```

**17. PUT - Actualizar solo Status (suspender usuario)**
```http
PUT http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "status": "suspended"
}
```

**18. PUT - Actualizar solo Status (activar usuario)**
```http
PUT http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "status": "active"
}
```

**19. PUT - Cambiar Posición Estructural**
```http
PUT http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "structural_position_id": "7"
}
```

**20. PUT - Cambiar Unidad Orgánica**
```http
PUT http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "organic_unit_id": "10"
}
```

**21. PUT - Quitar Posición Estructural**
```http
PUT http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "structural_position_id": ""
}
```

**22. PUT - Actualizar datos personales**
```http
PUT http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "first_name": "Juan Alberto",
  "last_name": "Pérez García Mendoza",
  "phone": "+51987123456"
}
```

**23. DELETE - Eliminar Usuario (Soft Delete)**
```http
DELETE http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**24. PATCH - Restaurar Usuario**
```http
PATCH http://localhost:8080/api/users/a54e954d-0c61-4861-ab94-d49eaf516672/restore
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuestas esperadas:**

**GET /api/users (200 OK):**
```json
{
  "data": [
    {
      "id": "a54e954d-0c61-4861-ab94-d49eaf516672",
      "email": "juan.perez@regionayacucho.gob.pe",
      "first_name": "Juan",
      "last_name": "Pérez García",
      "phone": "+51987654321",
      "dni": "12345678",
      "status": "active",
      "created_at": "2024-12-12T10:00:00Z",
      "updated_at": "2024-12-12T10:00:00Z",
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "organic_unit": {
        "id": "5",
        "name": "Gerencia de Tecnologías de la Información",
        "acronym": "GTI",
        "parent_id": "1"
      },
      "structural_position": {
        "id": "3",
        "name": "Especialista en Sistemas",
        "code": "ESP-SIS-001",
        "level": 3
      },
      "ubigeo": {
        "id": "1",
        "ubigeo_code": "050101",
        "department": "Ayacucho",
        "province": "Huamanga",
        "district": "Ayacucho"
      }
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10
}
```

**GET /api/users/stats (200 OK):**
```json
{
  "total_users": 350,
  "active_users": 320,
  "suspended_users": 25,
  "new_users_last_month": 15
}
```

**GET /api/users/:id (200 OK):**
```json
{
  "id": "a54e954d-0c61-4861-ab94-d49eaf516672",
  "email": "juan.perez@regionayacucho.gob.pe",
  "dni": "12345678",
  "status": "active",
  "first_name": "Juan",
  "last_name": "Pérez García",
  "phone": "+51987654321",
  "cod_emp_sgd": "EMP01",
  "structural_position_id": "3",
  "organic_unit_id": "5",
  "ubigeo_id": "1",
  "created_at": "2024-12-12T10:00:00Z",
  "updated_at": "2024-12-12T10:00:00Z",
  "is_deleted": false,
  "deleted_at": null,
  "deleted_by": null,
  "structural_position": {
    "id": "3",
    "name": "Especialista en Sistemas",
    "code": "ESP-SIS-001",
    "level": 3
  },
  "organic_unit": {
    "id": "5",
    "name": "Gerencia de Tecnologías de la Información",
    "acronym": "GTI",
    "parent_id": "1"
  },
  "ubigeo": {
    "id": "1",
    "ubigeo_code": "050101",
    "department": "Ayacucho",
    "province": "Huamanga",
    "district": "Ayacucho"
  }
}
```

**POST /api/users (201 Created):**
```json
{
  "id": "b54e954d-0c61-4861-ab94-d49eaf516673",
  "email": "juan.perez@regionayacucho.gob.pe",
  "dni": "12345678",
  "status": "active",
  "first_name": "Juan",
  "last_name": "Pérez García",
  "phone": "+51987654321",
  "cod_emp_sgd": "EMP01",
  "structural_position_id": "3",
  "organic_unit_id": "5",
  "ubigeo_id": "1",
  "created_at": "2024-12-12T15:30:00Z",
  "updated_at": "2024-12-12T15:30:00Z",
  "is_deleted": false,
  "deleted_at": null,
  "deleted_by": null,
  "structural_position": {
    "id": "3",
    "name": "Especialista en Sistemas",
    "code": "ESP-SIS-001",
    "level": 3
  },
  "organic_unit": {
    "id": "5",
    "name": "Gerencia de Tecnologías de la Información",
    "acronym": "GTI",
    "parent_id": "1"
  },
  "ubigeo": {
    "id": "1",
    "ubigeo_code": "050101",
    "department": "Ayacucho",
    "province": "Huamanga",
    "district": "Ayacucho"
  }
}
```

**DELETE /api/users/:id (200 OK):**
```json
{
  "message": "Usuario eliminado correctamente"
}
```

**PATCH /api/users/:id/restore (200 OK):**
```json
{
  "message": "Usuario restaurado correctamente"
}
```

**Errores comunes:**

**400 Bad Request (posición no encontrada):**
```json
{
  "error": "posición estructural no encontrada"
}
```

**400 Bad Request (unidad orgánica no encontrada):**
```json
{
  "error": "unidad orgánica no encontrada"
}
```

**404 Not Found:**
```json
{
  "error": "usuario no encontrado"
}
```

**409 Conflict (email duplicado):**
```json
{
  "error": "ya existe un usuario con este email"
}
```

**409 Conflict (DNI duplicado):**
```json
{
  "error": "ya existe un usuario con este DNI"
}
```

**Notas importantes:**
- **Status válidos**: `active`, `suspended`, `inactive`
- **DNI**: Debe ser exactamente 8 dígitos numéricos
- **Email**: Debe ser un email válido
- **Phone**: Formato internacional recomendado (ej: +51987654321)
- **Ubigeo**: Código de ubicación geográfica peruano (departamento-provincia-distrito)