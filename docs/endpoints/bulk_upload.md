## Ubigeos

### 1. GET - Descargar Plantilla Excel de Ubigeos
```http
GET http://localhost:8080/api/ubigeos/template
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuesta (200 OK):**
- Content-Type: `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`
- Content-Disposition: `attachment; filename=plantilla_ubigeos_20241224_153045.xlsx`
- Archivo Excel con columnas: `ubigeo_code`, `inei_code`, `department`, `province`, `district`

---

### 2. POST - Carga Masiva de Ubigeos
```http
POST http://localhost:8080/api/ubigeos/bulk
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "ubigeos": [
    {
      "ubigeo_code": "050101",
      "inei_code": "050101",
      "department": "AYACUCHO",
      "province": "HUAMANGA",
      "district": "AYACUCHO"
    },
    {
      "ubigeo_code": "050102",
      "inei_code": "050102",
      "department": "AYACUCHO",
      "province": "HUAMANGA",
      "district": "ACOCRO"
    },
    {
      "ubigeo_code": "050103",
      "inei_code": "050103",
      "department": "AYACUCHO",
      "province": "HUAMANGA",
      "district": "ACOS VINCHOS"
    }
  ]
}
```

**Respuesta (201 Created):**
```json
{
  "created": 2,
  "updated": 0,
  "skipped": 1,
  "failed": 0,
  "errors": [
    {
      "row": 2,
      "field": "ubigeo_code",
      "message": "Ubigeo 050101 ya existe"
    }
  ],
  "details": [
    {
      "id": 1874,
      "ubigeo_code": "050102",
      "district": "ACOCRO"
    },
    {
      "id": 1875,
      "ubigeo_code": "050103",
      "district": "ACOS VINCHOS"
    }
  ]
}
```

**Errores comunes:**

**400 Bad Request (sin datos):**
```json
{
  "error": "No se proporcionaron ubigeos"
}
```

**400 Bad Request (excede límite):**
```json
{
  "error": "Máximo 1000 registros por carga"
}
```

---

## Unidades Orgánicas

### 3. GET - Descargar Plantilla Excel de Unidades Orgánicas
```http
GET http://localhost:8080/api/organic-units/template
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuesta (200 OK):**
- Content-Type: `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`
- Content-Disposition: `attachment; filename=plantilla_unidades_organicas_20241224_153045.xlsx`
- Archivo Excel con 2 hojas:
  - **Unidades Orgánicas**: Plantilla con columnas `name`, `acronym`, `brand`, `description`, `parent_id`, `is_active`, `cod_dep_sgd`
  - **Unidades Existentes**: Referencia con `ID`, `Nombre`, `Acrónimo`, `ID Padre`

---

### 4. POST - Carga Masiva de Unidades Orgánicas
```http
POST http://localhost:8080/api/organic-units/bulk
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "organic_units": [
    {
      "name": "Gerencia Regional de Desarrollo Social",
      "acronym": "GRDS",
      "description": "Gerencia encargada de políticas sociales",
      "is_active": true,
      "cod_dep_sgd": "001"
    },
    {
      "name": "Sub Gerencia de Programas Sociales",
      "acronym": "SGPS",
      "description": "Sub gerencia dependiente de GRDS",
      "parent_id": 1,
      "is_active": true,
      "cod_dep_sgd": "002"
    },
    {
      "name": "Gerencia Regional de Infraestructura",
      "acronym": "GRI",
      "brand": "Obras Ayacucho",
      "is_active": true
    }
  ]
}
```

**Respuesta (201 Created):**
```json
{
  "created": 3,
  "updated": 0,
  "skipped": 0,
  "failed": 0,
  "errors": [],
  "details": [
    {
      "id": 45,
      "name": "Gerencia Regional de Desarrollo Social",
      "acronym": "GRDS"
    },
    {
      "id": 46,
      "name": "Sub Gerencia de Programas Sociales",
      "acronym": "SGPS"
    },
    {
      "id": 47,
      "name": "Gerencia Regional de Infraestructura",
      "acronym": "GRI"
    }
  ]
}
```

**Respuesta con errores (201 Created):**
```json
{
  "created": 1,
  "updated": 0,
  "skipped": 1,
  "failed": 1,
  "errors": [
    {
      "row": 2,
      "field": "name",
      "message": "Unidad 'Gerencia Regional de Desarrollo Social' ya existe"
    },
    {
      "row": 3,
      "field": "parent_id",
      "message": "Unidad padre 999 no existe"
    }
  ],
  "details": [
    {
      "id": 48,
      "name": "Gerencia Regional de Infraestructura",
      "acronym": "GRI"
    }
  ]
}
```

**Errores comunes:**

**400 Bad Request (sin datos):**
```json
{
  "error": "No se proporcionaron unidades"
}
```

**400 Bad Request (excede límite):**
```json
{
  "error": "Máximo 500 registros por carga"
}
```

---

## Posiciones Estructurales

### 5. GET - Descargar Plantilla Excel de Posiciones
```http
GET http://localhost:8080/api/positions/template
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuesta (200 OK):**
- Content-Type: `application/vnd.openxmlformats-officedocument.spreadsheetml.sheet`
- Content-Disposition: `attachment; filename=plantilla_posiciones_20241224_153045.xlsx`
- Archivo Excel con columnas: `name`, `code`, `level`, `description`, `is_active`, `cod_car_sgd`

---

### 6. POST - Carga Masiva de Posiciones Estructurales
```http
POST http://localhost:8080/api/positions/bulk
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "positions": [
    {
      "name": "Director General",
      "code": "DG-001",
      "level": 1,
      "description": "Cargo directivo de mayor jerarquía",
      "is_active": true,
      "cod_car_sgd": "0001"
    },
    {
      "name": "Gerente Regional",
      "code": "GR-001",
      "level": 2,
      "description": "Cargo gerencial regional",
      "is_active": true,
      "cod_car_sgd": "0002"
    },
    {
      "name": "Sub Gerente",
      "code": "SG-001",
      "level": 3,
      "is_active": true
    },
    {
      "name": "Especialista Administrativo",
      "code": "EA-001",
      "level": 4,
      "description": "Profesional especializado",
      "is_active": true,
      "cod_car_sgd": "0004"
    }
  ]
}
```

**Respuesta (201 Created):**
```json
{
  "created": 4,
  "updated": 0,
  "skipped": 0,
  "failed": 0,
  "errors": [],
  "details": [
    {
      "id": 101,
      "name": "Director General",
      "code": "DG-001"
    },
    {
      "id": 102,
      "name": "Gerente Regional",
      "code": "GR-001"
    },
    {
      "id": 103,
      "name": "Sub Gerente",
      "code": "SG-001"
    },
    {
      "id": 104,
      "name": "Especialista Administrativo",
      "code": "EA-001"
    }
  ]
}
```

**Respuesta con errores (201 Created):**
```json
{
  "created": 2,
  "updated": 0,
  "skipped": 2,
  "failed": 0,
  "errors": [
    {
      "row": 2,
      "field": "name",
      "message": "Posición 'Director General' ya existe"
    },
    {
      "row": 4,
      "field": "code",
      "message": "Código 'GR-001' ya existe"
    }
  ],
  "details": [
    {
      "id": 105,
      "name": "Sub Gerente",
      "code": "SG-001"
    },
    {
      "id": 106,
      "name": "Especialista Administrativo",
      "code": "EA-001"
    }
  ]
}
```

**Errores comunes:**

**400 Bad Request (sin datos):**
```json
{
  "error": "No se proporcionaron posiciones"
}
```

**400 Bad Request (excede límite):**
```json
{
  "error": "Máximo 500 registros por carga"
}
```

**400 Bad Request (campos obligatorios):**
```json
{
  "error": "Datos inválidos"
}
```

---

## Resumen de Endpoints

| Recurso | Método | Endpoint | Descripción | Límite |
|---------|--------|----------|-------------|--------|
| Ubigeos | GET | `/api/ubigeos/template` | Descargar plantilla Excel | - |
| Ubigeos | POST | `/api/ubigeos/bulk` | Carga masiva JSON | 1000 |
| Unidades Orgánicas | GET | `/api/organic-units/template` | Descargar plantilla Excel | - |
| Unidades Orgánicas | POST | `/api/organic-units/bulk` | Carga masiva JSON | 500 |
| Posiciones | GET | `/api/positions/template` | Descargar plantilla Excel | - |
| Posiciones | POST | `/api/positions/bulk` | Carga masiva JSON | 500 |

---

## Estructura de Respuesta de Carga Masiva

```json
{
  "created": 0,      // Registros creados exitosamente
  "updated": 0,      // Registros actualizados (si aplica)
  "skipped": 0,      // Registros omitidos (duplicados)
  "failed": 0,       // Registros con errores
  "errors": [        // Detalle de errores
    {
      "row": 2,      // Fila del error (basado en Excel)
      "field": "",   // Campo con error (opcional)
      "message": ""  // Descripción del error
    }
  ],
  "details": [       // Registros creados exitosamente
    {
      "id": 0,
      "...": "..."
    }
  ]
}
```