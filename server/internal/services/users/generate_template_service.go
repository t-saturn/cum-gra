package services

import (
	"fmt"
	"server/internal/config"
	"server/internal/models"

	"github.com/xuri/excelize/v2"
)

func GenerateUsersTemplateExcel() (*excelize.File, error) {
	db := config.DB
	f := excelize.NewFile()

	// Crear las 4 hojas
	sheetNames := []string{"Usuarios", "Posiciones", "Unidades Orgánicas", "Ubigeos"}

	// Crear hojas
	f.SetSheetName("Sheet1", sheetNames[0])
	for i := 1; i < len(sheetNames); i++ {
		_, err := f.NewSheet(sheetNames[i])
		if err != nil {
			return nil, fmt.Errorf("error creando hoja %s: %w", sheetNames[i], err)
		}
	}

	// Estilo para encabezados
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Size:  11,
			Color: "FFFFFF",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"4472C4"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})

	// Estilo para texto (para preservar ceros a la izquierda)
	textStyle, _ := f.NewStyle(&excelize.Style{
		CustomNumFmt: strPtr("@"), // Formato de texto
	})

	// Estilo para números
	numberStyle, _ := f.NewStyle(&excelize.Style{
		CustomNumFmt: strPtr("0"),
	})

	// === HOJA 1: PLANTILLA DE USUARIOS ===
	usersSheet := sheetNames[0]

	// Encabezados
	headers := []string{"email", "dni", "first_name", "last_name", "password", "phone", "status", "cod_emp_sgd", "structural_position_id", "organic_unit_id", "ubigeo_id"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(usersSheet, cell, header)
	}
	f.SetCellStyle(usersSheet, "A1", "K1", headerStyle)

	// Ejemplo de fila
	exampleData := []interface{}{
		"juan.perez@regionayacucho.gob.pe",
		"12345678",
		"Juan",
		"Pérez García",
		"MiClave123",
		"+51987654321",
		"active",
		"EMP01",
		"1",
		"1",
		"1",
	}
	for i, val := range exampleData {
		cell, _ := excelize.CoordinatesToCellName(i+1, 2)
		f.SetCellValue(usersSheet, cell, val)
	}

	// Aplicar formato de TEXTO a columnas que necesitan preservar ceros
	// DNI (B), Password (E), Phone (F), cod_emp_sgd (H)
	f.SetCellStyle(usersSheet, "B2", "B1001", textStyle) // DNI
	f.SetCellStyle(usersSheet, "E2", "E1001", textStyle) // Password
	f.SetCellStyle(usersSheet, "F2", "F1001", textStyle) // Phone
	f.SetCellStyle(usersSheet, "H2", "H1001", textStyle) // cod_emp_sgd

	// Aplicar formato de NÚMERO a columnas de IDs
	f.SetCellStyle(usersSheet, "I2", "I1001", numberStyle) // structural_position_id
	f.SetCellStyle(usersSheet, "J2", "J1001", numberStyle) // organic_unit_id
	f.SetCellStyle(usersSheet, "K2", "K1001", numberStyle) // ubigeo_id

	// Crear tabla de Excel para la hoja de Usuarios
	err := f.AddTable(usersSheet, &excelize.Table{
		Range:             "A1:K1001",
		Name:              "TablaUsuarios",
		StyleName:         "TableStyleMedium2",
		ShowFirstColumn:   false,
		ShowLastColumn:    false,
		ShowRowStripes:    boolPtr(true),
		ShowColumnStripes: false,
	})
	if err != nil {
		return nil, fmt.Errorf("error creando tabla de usuarios: %w", err)
	}

	// Ajustar anchos de columnas
	columnWidths := map[string]float64{
		"A": 35, // email
		"B": 12, // dni
		"C": 20, // first_name
		"D": 25, // last_name
		"E": 15, // password
		"F": 18, // phone
		"G": 12, // status
		"H": 15, // cod_emp_sgd
		"I": 25, // structural_position_id
		"J": 23, // organic_unit_id
		"K": 15, // ubigeo_id
	}
	for col, width := range columnWidths {
		f.SetColWidth(usersSheet, col, col, width)
	}

	// Agregar nota explicativa
	f.SetCellValue(usersSheet, "A1002", "INSTRUCCIONES:")
	f.SetCellValue(usersSheet, "A1003", "1. Completa los datos en las filas superiores (máximo 1000 usuarios)")
	f.SetCellValue(usersSheet, "A1004", "2. El DNI debe ser de 8 dígitos")
	f.SetCellValue(usersSheet, "A1005", "3. Si no ingresas contraseña, se usará el DNI como contraseña")
	f.SetCellValue(usersSheet, "A1006", "4. Status válidos: active, suspended, inactive")
	f.SetCellValue(usersSheet, "A1007", "5. Consulta las otras hojas para ver los IDs disponibles")

	// === HOJA 2: POSICIONES ESTRUCTURALES ===
	positionsSheet := sheetNames[1]

	var positions []models.StructuralPosition
	if err := db.Where("is_deleted = FALSE AND is_active = TRUE").
		Order("id ASC"). // CAMBIADO: Ordenar por ID ascendente
		Find(&positions).Error; err != nil {
		return nil, fmt.Errorf("error obteniendo posiciones: %w", err)
	}

	// Encabezados
	posHeaders := []string{"ID", "Código", "Nombre", "Nivel"}
	for i, header := range posHeaders {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(positionsSheet, cell, header)
	}
	f.SetCellStyle(positionsSheet, "A1", "D1", headerStyle)

	// Datos
	for idx, pos := range positions {
		row := idx + 2
		f.SetCellValue(positionsSheet, fmt.Sprintf("A%d", row), pos.ID)
		f.SetCellValue(positionsSheet, fmt.Sprintf("B%d", row), pos.Code)
		f.SetCellValue(positionsSheet, fmt.Sprintf("C%d", row), pos.Name)
		if pos.Level != nil {
			f.SetCellValue(positionsSheet, fmt.Sprintf("D%d", row), *pos.Level)
		}
	}

	// Crear tabla para Posiciones
	if len(positions) > 0 {
		lastRow := len(positions) + 1
		tableRange := fmt.Sprintf("A1:D%d", lastRow)
		err := f.AddTable(positionsSheet, &excelize.Table{
			Range:             tableRange,
			Name:              "TablaPosiciones",
			StyleName:         "TableStyleMedium9",
			ShowFirstColumn:   false,
			ShowLastColumn:    false,
			ShowRowStripes:    boolPtr(true),
			ShowColumnStripes: false,
		})
		if err != nil {
			return nil, fmt.Errorf("error creando tabla de posiciones: %w", err)
		}
	}

	// Ajustar anchos
	f.SetColWidth(positionsSheet, "A", "A", 8)
	f.SetColWidth(positionsSheet, "B", "B", 20)
	f.SetColWidth(positionsSheet, "C", "C", 50)
	f.SetColWidth(positionsSheet, "D", "D", 10)

	// === HOJA 3: UNIDADES ORGÁNICAS ===
	unitsSheet := sheetNames[2]

	var organicUnits []models.OrganicUnit
	if err := db.Where("is_deleted = FALSE AND is_active = TRUE").
		Order("id ASC"). // CAMBIADO: Ordenar por ID ascendente
		Find(&organicUnits).Error; err != nil {
		return nil, fmt.Errorf("error obteniendo unidades orgánicas: %w", err)
	}

	// Encabezados
	unitHeaders := []string{"ID", "Nombre", "Acrónimo", "ID Padre"}
	for i, header := range unitHeaders {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(unitsSheet, cell, header)
	}
	f.SetCellStyle(unitsSheet, "A1", "D1", headerStyle)

	// Datos
	for idx, unit := range organicUnits {
		row := idx + 2
		f.SetCellValue(unitsSheet, fmt.Sprintf("A%d", row), unit.ID)
		f.SetCellValue(unitsSheet, fmt.Sprintf("B%d", row), unit.Name)
		f.SetCellValue(unitsSheet, fmt.Sprintf("C%d", row), unit.Acronym)
		if unit.ParentID != nil {
			f.SetCellValue(unitsSheet, fmt.Sprintf("D%d", row), *unit.ParentID)
		}
	}

	// Crear tabla para Unidades
	if len(organicUnits) > 0 {
		lastRow := len(organicUnits) + 1
		tableRange := fmt.Sprintf("A1:D%d", lastRow)
		err := f.AddTable(unitsSheet, &excelize.Table{
			Range:             tableRange,
			Name:              "TablaUnidades",
			StyleName:         "TableStyleMedium13",
			ShowFirstColumn:   false,
			ShowLastColumn:    false,
			ShowRowStripes:    boolPtr(true),
			ShowColumnStripes: false,
		})
		if err != nil {
			return nil, fmt.Errorf("error creando tabla de unidades: %w", err)
		}
	}

	// Ajustar anchos
	f.SetColWidth(unitsSheet, "A", "A", 8)
	f.SetColWidth(unitsSheet, "B", "B", 55)
	f.SetColWidth(unitsSheet, "C", "C", 15)
	f.SetColWidth(unitsSheet, "D", "D", 12)

	// === HOJA 4: UBIGEOS ===
	ubigeosSheet := sheetNames[3]

	var ubigeos []models.Ubigeo
	if err := db.Order("id ASC"). // CAMBIADO: Ordenar por ID ascendente
		Find(&ubigeos).Error; err != nil {
		return nil, fmt.Errorf("error obteniendo ubigeos: %w", err)
	}

	// Encabezados
	ubigeoHeaders := []string{"ID", "Código Ubigeo", "Departamento", "Provincia", "Distrito"}
	for i, header := range ubigeoHeaders {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(ubigeosSheet, cell, header)
	}
	f.SetCellStyle(ubigeosSheet, "A1", "E1", headerStyle)

	// Datos
	for idx, ubigeo := range ubigeos {
		row := idx + 2
		f.SetCellValue(ubigeosSheet, fmt.Sprintf("A%d", row), ubigeo.ID)
		f.SetCellValue(ubigeosSheet, fmt.Sprintf("B%d", row), ubigeo.UbigeoCode)
		f.SetCellValue(ubigeosSheet, fmt.Sprintf("C%d", row), ubigeo.Department)
		f.SetCellValue(ubigeosSheet, fmt.Sprintf("D%d", row), ubigeo.Province)
		f.SetCellValue(ubigeosSheet, fmt.Sprintf("E%d", row), ubigeo.District)
	}

	// Crear tabla para Ubigeos
	if len(ubigeos) > 0 {
		lastRow := len(ubigeos) + 1
		tableRange := fmt.Sprintf("A1:E%d", lastRow)
		err := f.AddTable(ubigeosSheet, &excelize.Table{
			Range:             tableRange,
			Name:              "TablaUbigeos",
			StyleName:         "TableStyleMedium6",
			ShowFirstColumn:   false,
			ShowLastColumn:    false,
			ShowRowStripes:    boolPtr(true),
			ShowColumnStripes: false,
		})
		if err != nil {
			return nil, fmt.Errorf("error creando tabla de ubigeos: %w", err)
		}
	}

	// Ajustar anchos
	f.SetColWidth(ubigeosSheet, "A", "A", 8)
	f.SetColWidth(ubigeosSheet, "B", "B", 15)
	f.SetColWidth(ubigeosSheet, "C", "C", 20)
	f.SetColWidth(ubigeosSheet, "D", "D", 20)
	f.SetColWidth(ubigeosSheet, "E", "E", 25)

	// Activar primera hoja
	f.SetActiveSheet(0)

	return f, nil
}

// Helper functions
func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}