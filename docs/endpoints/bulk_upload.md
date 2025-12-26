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

### 2. POST - Carga Masiva de Ubigeos (Excel)
```http
POST http://localhost:8080/api/ubigeos/bulk
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: multipart/form-data

file: [archivo.xlsx]
```

**Respuesta (200 OK):**
```json
{
  "created": 2,
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

**400 Bad Request (sin archivo):**
```json
{
  "error": "No se proporcionó archivo Excel"
}
```

**400 Bad Request (formato inválido):**
```json
{
  "error": "El archivo debe ser formato Excel (.xlsx o .xls)"
}
```

**400 Bad Request (sin datos):**
```json
{
  "error": "El archivo debe contener al menos una fila de datos"
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

### 4. POST - Carga Masiva de Unidades Orgánicas (Excel)
```http
POST http://localhost:8080/api/organic-units/bulk
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: multipart/form-data

file: [archivo.xlsx]
```

**Respuesta (200 OK):**
```json
{
  "created": 3,
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

**Respuesta con errores (200 OK):**
```json
{
  "created": 1,
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

**400 Bad Request (sin archivo):**
```json
{
  "error": "No se proporcionó archivo Excel"
}
```

**400 Bad Request (formato inválido):**
```json
{
  "error": "El archivo debe ser formato Excel (.xlsx o .xls)"
}
```

**400 Bad Request (sin datos):**
```json
{
  "error": "No se encontraron unidades válidas en el archivo"
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

### 6. POST - Carga Masiva de Posiciones Estructurales (Excel)
```http
POST http://localhost:8080/api/positions/bulk
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: multipart/form-data

file: [archivo.xlsx]
```

**Respuesta (200 OK):**
```json
{
  "created": 4,
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

**Respuesta con errores (200 OK):**
```json
{
  "created": 2,
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

**400 Bad Request (sin archivo):**
```json
{
  "error": "No se proporcionó archivo Excel"
}
```

**400 Bad Request (formato inválido):**
```json
{
  "error": "El archivo debe ser formato Excel (.xlsx o .xls)"
}
```

**400 Bad Request (sin datos):**
```json
{
  "error": "No se encontraron posiciones válidas en el archivo"
}
```

**400 Bad Request (excede límite):**
```json
{
  "error": "Máximo 500 registros por carga"
}
```

---

## Resumen de Endpoints

| Recurso | Método | Endpoint | Content-Type | Descripción | Límite |
|---------|--------|----------|--------------|-------------|--------|
| Ubigeos | GET | `/api/ubigeos/template` | - | Descargar plantilla Excel | - |
| Ubigeos | POST | `/api/ubigeos/bulk` | `multipart/form-data` | Carga masiva Excel | 1000 |
| Unidades Orgánicas | GET | `/api/organic-units/template` | - | Descargar plantilla Excel | - |
| Unidades Orgánicas | POST | `/api/organic-units/bulk` | `multipart/form-data` | Carga masiva Excel | 500 |
| Posiciones | GET | `/api/positions/template` | - | Descargar plantilla Excel | - |
| Posiciones | POST | `/api/positions/bulk` | `multipart/form-data` | Carga masiva Excel | 500 |

---

## Estructura del Archivo Excel

### Ubigeos
| Columna | Campo | Obligatorio | Descripción |
|---------|-------|-------------|-------------|
| A | ubigeo_code | Sí | Código ubigeo (6 dígitos) |
| B | inei_code | No | Código INEI |
| C | department | Sí | Nombre del departamento |
| D | province | Sí | Nombre de la provincia |
| E | district | Sí | Nombre del distrito |

### Unidades Orgánicas
| Columna | Campo | Obligatorio | Descripción |
|---------|-------|-------------|-------------|
| A | name | Sí | Nombre de la unidad |
| B | acronym | Sí | Acrónimo (ej: GTI) |
| C | brand | No | Marca o imagen |
| D | description | No | Descripción |
| E | parent_id | No | ID de unidad padre |
| F | is_active | No | true/false (default: true) |
| G | cod_dep_sgd | No | Código SGD |

### Posiciones Estructurales
| Columna | Campo | Obligatorio | Descripción |
|---------|-------|-------------|-------------|
| A | name | Sí | Nombre de la posición |
| B | code | Sí | Código único |
| C | level | No | Nivel jerárquico (1-10) |
| D | description | No | Descripción |
| E | is_active | No | true/false (default: true) |
| F | cod_car_sgd | No | Código SGD |

---

## Estructura de Respuesta de Carga Masiva

```json
{
  "created": 0,      // Registros creados exitosamente
  "skipped": 0,      // Registros omitidos (duplicados)
  "failed": 0,       // Registros con errores de validación
  "errors": [        // Detalle de errores y omisiones
    {
      "row": 2,      // Número de fila en Excel (header = 1)
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