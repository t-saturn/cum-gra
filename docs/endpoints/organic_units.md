**1. GET - Listar Unidades Orgánicas (sin filtros)**
```http
GET http://localhost:8080/api/organic-units?page=1&page_size=10
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

### Unidades Orgánicas - Listado completo para Selects
```http
GET http://localhost:8080/api/organic-units/all
GET http://localhost:8080/api/organic-units/all?only_active=false
Authorization: Bearer TOKEN
```

**2. GET - Listar Unidades Orgánicas raíz (sin padre)**
```http
GET http://localhost:8080/api/organic-units?page=1&page_size=10&parent_id=null
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Listar Sub-unidades de una unidad específica**
```http
GET http://localhost:8080/api/organic-units?page=1&page_size=10&parent_id=5
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. GET - Listar Unidades Orgánicas eliminadas**
```http
GET http://localhost:8080/api/organic-units?page=1&page_size=10&is_deleted=true
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**5. GET - Estadísticas de Unidades Orgánicas**
```http
GET http://localhost:8080/api/organic-units/stats
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**6. GET - Obtener Unidad Orgánica por ID**
```http
GET http://localhost:8080/api/organic-units/5
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**7. POST - Crear Unidad Orgánica (completa con padre)**
```http
POST http://localhost:8080/api/organic-units
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Gerencia de Tecnologías de la Información",
  "acronym": "GTI",
  "brand": "TI Ayacucho",
  "description": "Unidad encargada de la gestión de tecnologías de información",
  "parent_id": "1",
  "is_active": true,
  "cod_dep_sgd": "GTI01"
}
```

**8. POST - Crear Unidad Orgánica raíz (sin padre)**
```http
POST http://localhost:8080/api/organic-units
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Gerencia Regional",
  "acronym": "GR",
  "description": "Gerencia Regional de Ayacucho",
  "is_active": true
}
```

**9. POST - Crear Unidad Orgánica (mínimo requerido)**
```http
POST http://localhost:8080/api/organic-units
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Oficina de Recursos Humanos",
  "acronym": "ORH"
}
```

**10. POST - Crear Sub-unidad**
```http
POST http://localhost:8080/api/organic-units
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Área de Desarrollo de Sistemas",
  "acronym": "ADS",
  "description": "Desarrollo y mantenimiento de sistemas informáticos",
  "parent_id": "5",
  "is_active": true,
  "cod_dep_sgd": "ADS01"
}
```

**11. PUT - Actualizar Unidad Orgánica (completo)**
```http
PUT http://localhost:8080/api/organic-units/5
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Gerencia de Tecnologías de la Información y Comunicaciones",
  "acronym": "GTIC",
  "brand": "TIC Ayacucho",
  "description": "Unidad encargada de la gestión de TI y comunicaciones",
  "is_active": true,
  "cod_dep_sgd": "GTIC1"
}
```

**12. PUT - Actualizar Unidad Orgánica (solo nombre)**
```http
PUT http://localhost:8080/api/organic-units/5
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Gerencia de Sistemas e Informática"
}
```

**13. PUT - Actualizar Unidad Orgánica (cambiar padre)**
```http
PUT http://localhost:8080/api/organic-units/5
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "parent_id": "3"
}
```

**14. PUT - Actualizar Unidad Orgánica (quitar padre - convertir en raíz)**
```http
PUT http://localhost:8080/api/organic-units/5
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "parent_id": ""
}
```

**15. PUT - Actualizar Unidad Orgánica (desactivar)**
```http
PUT http://localhost:8080/api/organic-units/5
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "is_active": false
}
```

**16. PUT - Actualizar código SGD**
```http
PUT http://localhost:8080/api/organic-units/5
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "cod_dep_sgd": "GTIC2"
}
```

**17. DELETE - Eliminar Unidad Orgánica (Soft Delete)**
```http
DELETE http://localhost:8080/api/organic-units/5
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**18. PATCH - Restaurar Unidad Orgánica**
```http
PATCH http://localhost:8080/api/organic-units/5/restore
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuestas esperadas:**

**GET /api/organic-units (200 OK):**
```json
{
  "data": [
    {
      "id": "5",
      "name": "Gerencia de Tecnologías de la Información",
      "acronym": "GTI",
      "brand": "TI Ayacucho",
      "description": "Unidad encargada de la gestión de tecnologías de información",
      "parent_id": "1",
      "is_active": true,
      "created_at": "2024-12-12T10:00:00Z",
      "updated_at": "2024-12-12T10:00:00Z",
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "cod_dep_sgd": "GTI01",
      "users_count": 25
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10
}
```

**GET /api/organic-units/stats (200 OK):**
```json
{
  "total_organic_units": 50,
  "active_organic_units": 45,
  "deleted_organic_units": 5,
  "total_employees": 350
}
```

**POST /api/organic-units (201 Created):**
```json
{
  "id": "15",
  "name": "Gerencia de Tecnologías de la Información",
  "acronym": "GTI",
  "brand": "TI Ayacucho",
  "description": "Unidad encargada de la gestión de tecnologías de información",
  "parent_id": "1",
  "is_active": true,
  "created_at": "2024-12-12T15:30:00Z",
  "updated_at": "2024-12-12T15:30:00Z",
  "is_deleted": false,
  "deleted_at": null,
  "deleted_by": null,
  "cod_dep_sgd": "GTI01",
  "users_count": 0
}
```

**DELETE /api/organic-units/:id (200 OK):**
```json
{
  "message": "Unidad orgánica eliminada correctamente"
}
```

**PATCH /api/organic-units/:id/restore (200 OK):**
```json
{
  "message": "Unidad orgánica restaurada correctamente"
}
```

**Errores comunes:**

**400 Bad Request (tiene sub-unidades):**
```json
{
  "error": "no se puede eliminar una unidad orgánica que tiene sub-unidades"
}
```

**400 Bad Request (tiene usuarios asignados):**
```json
{
  "error": "no se puede eliminar una unidad orgánica que tiene usuarios asignados"
}
```

**400 Bad Request (padre inválido):**
```json
{
  "error": "una unidad orgánica no puede ser su propio padre"
}
```

**404 Not Found:**
```json
{
  "error": "unidad orgánica no encontrada"
}
```

**404 Not Found (padre no existe):**
```json
{
  "error": "unidad orgánica padre no encontrada"
}
```

**409 Conflict (nombre duplicado):**
```json
{
  "error": "ya existe una unidad orgánica con este nombre"
}
```

**409 Conflict (acrónimo duplicado):**
```json
{
  "error": "ya existe una unidad orgánica con este acrónimo"
}
```