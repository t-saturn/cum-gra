**1. GET - Listar Posiciones Estructurales (sin filtros)**
```http
GET http://localhost:8080/api/positions?page=1&page_size=10
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**2. GET - Listar Posiciones por Nivel**
```http
GET http://localhost:8080/api/positions?page=1&page_size=10&level=1
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Listar Posiciones Nivel Directivo (nivel 1)**
```http
GET http://localhost:8080/api/positions?page=1&page_size=10&level=1
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. GET - Listar Posiciones Nivel Ejecutivo (nivel 2)**
```http
GET http://localhost:8080/api/positions?page=1&page_size=10&level=2
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**5. GET - Listar Posiciones eliminadas**
```http
GET http://localhost:8080/api/positions?page=1&page_size=10&is_deleted=true
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**6. GET - Estadísticas de Posiciones**
```http
GET http://localhost:8080/api/positions/stats
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**7. GET - Obtener Posición por ID**
```http
GET http://localhost:8080/api/positions/3
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**8. POST - Crear Posición Estructural (completa)**
```http
POST http://localhost:8080/api/positions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Director Regional",
  "code": "DIR-REG-001",
  "level": 1,
  "description": "Máxima autoridad de la región",
  "is_active": true,
  "cod_car_sgd": "DR01"
}
```

**9. POST - Crear Posición Nivel Ejecutivo**
```http
POST http://localhost:8080/api/positions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Gerente de Tecnología",
  "code": "GER-TEC-001",
  "level": 2,
  "description": "Responsable de la gestión tecnológica",
  "is_active": true,
  "cod_car_sgd": "GT01"
}
```

**10. POST - Crear Posición Nivel Operativo**
```http
POST http://localhost:8080/api/positions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Especialista en Sistemas",
  "code": "ESP-SIS-001",
  "level": 3,
  "description": "Desarrollo y soporte de sistemas informáticos",
  "is_active": true
}
```

**11. POST - Crear Posición (mínimo requerido)**
```http
POST http://localhost:8080/api/positions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Asistente Administrativo",
  "code": "AST-ADM-001"
}
```

**12. POST - Crear Posición de Soporte**
```http
POST http://localhost:8080/api/positions
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Técnico de Soporte TI",
  "code": "TEC-SOP-001",
  "level": 4,
  "description": "Soporte técnico a usuarios finales",
  "is_active": true,
  "cod_car_sgd": "TS01"
}
```

**13. PUT - Actualizar Posición (completo)**
```http
PUT http://localhost:8080/api/positions/3
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Director Regional de Gobierno",
  "code": "DIR-REG-GOB-001",
  "level": 1,
  "description": "Máxima autoridad del gobierno regional",
  "is_active": true,
  "cod_car_sgd": "DRG1"
}
```

**14. PUT - Actualizar Posición (solo nombre)**
```http
PUT http://localhost:8080/api/positions/3
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "name": "Gerente General"
}
```

**15. PUT - Actualizar Posición (cambiar nivel)**
```http
PUT http://localhost:8080/api/positions/3
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "level": 2
}
```

**16. PUT - Actualizar Posición (desactivar)**
```http
PUT http://localhost:8080/api/positions/3
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "is_active": false
}
```

**17. PUT - Actualizar código SGD**
```http
PUT http://localhost:8080/api/positions/3
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "cod_car_sgd": "DR02"
}
```

**18. DELETE - Eliminar Posición (Soft Delete)**
```http
DELETE http://localhost:8080/api/positions/3
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YXhwbWJxS0ljIn0...
```

**19. PATCH - Restaurar Posición**
```http
PATCH http://localhost:8080/api/positions/3/restore
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuestas esperadas:**

**GET /api/positions (200 OK):**
```json
{
  "data": [
    {
      "id": "3",
      "name": "Director Regional",
      "code": "DIR-REG-001",
      "level": 1,
      "description": "Máxima autoridad de la región",
      "is_active": true,
      "created_at": "2024-12-12T10:00:00Z",
      "updated_at": "2024-12-12T10:00:00Z",
      "is_deleted": false,
      "deleted_at": null,
      "deleted_by": null,
      "cod_car_sgd": "DR01",
      "users_count": 5
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10
}
```

**GET /api/positions/stats (200 OK):**
```json
{
  "total_positions": 45,
  "active_positions": 40,
  "deleted_positions": 5,
  "assigned_employees": 320
}
```

**POST /api/positions (201 Created):**
```json
{
  "id": "15",
  "name": "Director Regional",
  "code": "DIR-REG-001",
  "level": 1,
  "description": "Máxima autoridad de la región",
  "is_active": true,
  "created_at": "2024-12-12T15:30:00Z",
  "updated_at": "2024-12-12T15:30:00Z",
  "is_deleted": false,
  "deleted_at": null,
  "deleted_by": null,
  "cod_car_sgd": "DR01",
  "users_count": 0
}
```

**DELETE /api/positions/:id (200 OK):**
```json
{
  "message": "Posición estructural eliminada correctamente"
}
```

**PATCH /api/positions/:id/restore (200 OK):**
```json
{
  "message": "Posición estructural restaurada correctamente"
}
```

**Errores comunes:**

**400 Bad Request (tiene usuarios asignados):**
```json
{
  "error": "no se puede eliminar una posición estructural que tiene usuarios asignados"
}
```

**404 Not Found:**
```json
{
  "error": "posición estructural no encontrada"
}
```

**409 Conflict (nombre duplicado):**
```json
{
  "error": "ya existe una posición estructural con este nombre"
}
```

**409 Conflict (código duplicado):**
```json
{
  "error": "ya existe una posición estructural con este código"
}
```

**Notas sobre niveles jerárquicos comunes:**
- **Nivel 1**: Directivo (Director, Gerente General)
- **Nivel 2**: Ejecutivo (Gerente, Jefe de Área)
- **Nivel 3**: Profesional (Especialista, Analista)
- **Nivel 4**: Técnico (Técnico, Asistente)
- **Nivel 5**: Apoyo (Auxiliar, Practicante)