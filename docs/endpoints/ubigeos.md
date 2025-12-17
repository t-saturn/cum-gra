**1. GET - Listar Ubigeos (sin filtros)**
```http
GET http://localhost:8080/api/ubigeos?page=1&page_size=10
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**2. GET - Listar Ubigeos por Departamento**
```http
GET http://localhost:8080/api/ubigeos?page=1&page_size=10&department=Ayacucho
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Listar Ubigeos por Provincia**
```http
GET http://localhost:8080/api/ubigeos?page=1&page_size=10&province=Huamanga
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. GET - Listar Ubigeos por Distrito**
```http
GET http://localhost:8080/api/ubigeos?page=1&page_size=10&district=Ayacucho
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**5. GET - Listar Ubigeos por Departamento y Provincia**
```http
GET http://localhost:8080/api/ubigeos?page=1&page_size=10&department=Ayacucho&province=Huamanga
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**6. GET - Búsqueda parcial por Departamento (contiene "Aya")**
```http
GET http://localhost:8080/api/ubigeos?page=1&page_size=10&department=Aya
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**7. GET - Estadísticas de Ubigeos**
```http
GET http://localhost:8080/api/ubigeos/stats
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**8. GET - Obtener Ubigeo por ID**
```http
GET http://localhost:8080/api/ubigeos/1
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**9. POST - Crear Ubigeo (Ayacucho - Huamanga - Ayacucho)**
```http
POST http://localhost:8080/api/ubigeos
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "ubigeo_code": "050101",
  "inei_code": "050101",
  "department": "Ayacucho",
  "province": "Huamanga",
  "district": "Ayacucho"
}
```

**10. POST - Crear Ubigeo (Ayacucho - Huamanga - Carmen Alto)**
```http
POST http://localhost:8080/api/ubigeos
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "ubigeo_code": "050102",
  "inei_code": "050102",
  "department": "Ayacucho",
  "province": "Huamanga",
  "district": "Carmen Alto"
}
```

**11. POST - Crear Ubigeo (Ayacucho - Huamanga - San Juan Bautista)**
```http
POST http://localhost:8080/api/ubigeos
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "ubigeo_code": "050103",
  "inei_code": "050103",
  "department": "Ayacucho",
  "province": "Huamanga",
  "district": "San Juan Bautista"
}
```

**12. POST - Crear Ubigeo (Lima - Lima - Lima)**
```http
POST http://localhost:8080/api/ubigeos
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "ubigeo_code": "150101",
  "inei_code": "150101",
  "department": "Lima",
  "province": "Lima",
  "district": "Lima"
}
```

**13. POST - Crear Ubigeo (Cusco - Cusco - Cusco)**
```http
POST http://localhost:8080/api/ubigeos
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "ubigeo_code": "080101",
  "inei_code": "080101",
  "department": "Cusco",
  "province": "Cusco",
  "district": "Cusco"
}
```

**14. PUT - Actualizar Ubigeo (completo)**
```http
PUT http://localhost:8080/api/ubigeos/1
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "ubigeo_code": "050101",
  "inei_code": "050101",
  "department": "Ayacucho",
  "province": "Huamanga",
  "district": "Ayacucho (Ciudad)"
}
```

**15. PUT - Actualizar solo código INEI**
```http
PUT http://localhost:8080/api/ubigeos/1
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "inei_code": "050101-2024"
}
```

**16. PUT - Actualizar solo distrito**
```http
PUT http://localhost:8080/api/ubigeos/1
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
Content-Type: application/json

{
  "district": "Ayacucho - Centro Histórico"
}
```

**17. DELETE - Eliminar Ubigeo**
```http
DELETE http://localhost:8080/api/ubigeos/1
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**Respuestas esperadas:**

**GET /api/ubigeos (200 OK):**
```json
{
  "data": [
    {
      "id": "1",
      "ubigeo_code": "050101",
      "inei_code": "050101",
      "department": "Ayacucho",
      "province": "Huamanga",
      "district": "Ayacucho",
      "created_at": "2024-12-12T10:00:00Z",
      "updated_at": "2024-12-12T10:00:00Z"
    },
    {
      "id": "2",
      "ubigeo_code": "050102",
      "inei_code": "050102",
      "department": "Ayacucho",
      "province": "Huamanga",
      "district": "Carmen Alto",
      "created_at": "2024-12-12T10:01:00Z",
      "updated_at": "2024-12-12T10:01:00Z"
    }
  ],
  "total": 2,
  "page": 1,
  "page_size": 10
}
```

**GET /api/ubigeos/stats (200 OK):**
```json
{
  "total_ubigeos": 1874,
  "total_departments": 25,
  "total_provinces": 196,
  "total_districts": 1874
}
```

**GET /api/ubigeos/:id (200 OK):**
```json
{
  "id": "1",
  "ubigeo_code": "050101",
  "inei_code": "050101",
  "department": "Ayacucho",
  "province": "Huamanga",
  "district": "Ayacucho",
  "created_at": "2024-12-12T10:00:00Z",
  "updated_at": "2024-12-12T10:00:00Z"
}
```

**POST /api/ubigeos (201 Created):**
```json
{
  "id": "15",
  "ubigeo_code": "050101",
  "inei_code": "050101",
  "department": "Ayacucho",
  "province": "Huamanga",
  "district": "Ayacucho",
  "created_at": "2024-12-12T15:30:00Z",
  "updated_at": "2024-12-12T15:30:00Z"
}
```

**DELETE /api/ubigeos/:id (200 OK):**
```json
{
  "message": "Ubigeo eliminado correctamente"
}
```

**Errores comunes:**

**400 Bad Request (tiene usuarios asignados):**
```json
{
  "error": "no se puede eliminar un ubigeo que tiene usuarios asignados"
}
```

**404 Not Found:**
```json
{
  "error": "ubigeo no encontrado"
}
```

**409 Conflict (código duplicado):**
```json
{
  "error": "ya existe un ubigeo con este código"
}
```

**Notas importantes sobre Ubigeos en Perú:**

**Estructura del código Ubigeo:**
- **Formato**: 6 dígitos (DDPPDD)
  - **DD** (2 dígitos): Código del departamento
  - **PP** (2 dígitos): Código de la provincia
  - **DD** (2 dígitos): Código del distrito

**Uso típico:**
Los ubigeos se utilizan para:
- Registro de direcciones de usuarios


**Nuevos Endpoints para Selects:**

**1. GET - Lista de Departamentos (para select)**
```http
GET http://localhost:8080/api/ubigeos/departments
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**2. GET - Lista de Provincias de Ayacucho**
```http
GET http://localhost:8080/api/ubigeos/provinces?department=Ayacucho
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**3. GET - Lista de Provincias de Lima**
```http
GET http://localhost:8080/api/ubigeos/provinces?department=Lima
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**4. GET - Lista de Distritos de Huamanga (Ayacucho)**
```http
GET http://localhost:8080/api/ubigeos/districts?department=Ayacucho&province=Huamanga
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**5. GET - Lista de Distritos de Lima (Lima)**
```http
GET http://localhost:8080/api/ubigeos/districts?department=Lima&province=Lima
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJpVTRFY3NZMUxtVkdMYmg2MHhzSWJHMEtQUTVhRGpqN2w3YVhwbWJxS0ljIn0...
```

**GET /api/ubigeos/departments (200 OK):**
```json
[
  {
    "name": "Amazonas"
  },
  {
    "name": "Áncash"
  },
  {
    "name": "Apurímac"
  },
  {
    "name": "Arequipa"
  },
  {
    "name": "Ayacucho"
  },
  {
    "name": "Cajamarca"
  },
  {
    "name": "Callao"
  },
  {
    "name": "Cusco"
  },
  {
    "name": "Huancavelica"
  },
  {
    "name": "Huánuco"
  },
  {
    "name": "Ica"
  },
  {
    "name": "Junín"
  },
  {
    "name": "La Libertad"
  },
  {
    "name": "Lambayeque"
  },
  {
    "name": "Lima"
  },
  {
    "name": "Loreto"
  },
  {
    "name": "Madre de Dios"
  },
  {
    "name": "Moquegua"
  },
  {
    "name": "Pasco"
  },
  {
    "name": "Piura"
  },
  {
    "name": "Puno"
  },
  {
    "name": "San Martín"
  },
  {
    "name": "Tacna"
  },
  {
    "name": "Tumbes"
  },
  {
    "name": "Ucayali"
  }
]
```

**GET /api/ubigeos/provinces?department=Ayacucho (200 OK):**
```json
[
  {
    "name": "Cangallo",
    "department": "Ayacucho"
  },
  {
    "name": "Huamanga",
    "department": "Ayacucho"
  },
  {
    "name": "Huanca Sancos",
    "department": "Ayacucho"
  },
  {
    "name": "Huanta",
    "department": "Ayacucho"
  },
  {
    "name": "La Mar",
    "department": "Ayacucho"
  },
  {
    "name": "Lucanas",
    "department": "Ayacucho"
  },
  {
    "name": "Parinacochas",
    "department": "Ayacucho"
  },
  {
    "name": "Páucar del Sara Sara",
    "department": "Ayacucho"
  },
  {
    "name": "Sucre",
    "department": "Ayacucho"
  },
  {
    "name": "Víctor Fajardo",
    "department": "Ayacucho"
  },
  {
    "name": "Vilcas Huamán",
    "department": "Ayacucho"
  }
]
```

**GET /api/ubigeos/districts?department=Ayacucho&province=Huamanga (200 OK):**
```json
[
  {
    "id": "1",
    "name": "Ayacucho",
    "province": "Huamanga",
    "department": "Ayacucho",
    "ubigeo_code": "050101"
  },
  {
    "id": "2",
    "name": "Acocro",
    "province": "Huamanga",
    "department": "Ayacucho",
    "ubigeo_code": "050102"
  },
  {
    "id": "3",
    "name": "Acos Vinchos",
    "province": "Huamanga",
    "department": "Ayacucho",
    "ubigeo_code": "050103"
  },
  {
    "id": "4",
    "name": "Carmen Alto",
    "province": "Huamanga",
    "department": "Ayacucho",
    "ubigeo_code": "050104"
  },
  {
    "id": "5",
    "name": "Chiara",
    "province": "Huamanga",
    "department": "Ayacucho",
    "ubigeo_code": "050105"
  },
  {
    "id": "6",
    "name": "Jesús Nazareno",
    "province": "Huamanga",
    "department": "Ayacucho",
    "ubigeo_code": "050113"
  },
  {
    "id": "7",
    "name": "San Juan Bautista",
    "province": "Huamanga",
    "department": "Ayacucho",
    "ubigeo_code": "050114"
  },
  {
    "id": "8",
    "name": "Santiago de Pischa",
    "province": "Huamanga",
    "department": "Ayacucho",
    "ubigeo_code": "050115"
  }
]