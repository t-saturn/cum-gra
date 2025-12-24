package services

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func GenerateUbigeosTemplateExcel() (*excelize.File, error) {
	f := excelize.NewFile()
	sheetName := "Ubigeos"
	f.SetSheetName("Sheet1", sheetName)

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

	// Encabezados
	headers := []string{"ubigeo_code", "inei_code", "department", "province", "district"}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}
	f.SetCellStyle(sheetName, "A1", "E1", headerStyle)

	// Ejemplo
	exampleData := []interface{}{"050101", "050101", "AYACUCHO", "HUAMANGA", "AYACUCHO"}
	for i, val := range exampleData {
		cell, _ := excelize.CoordinatesToCellName(i+1, 2)
		f.SetCellValue(sheetName, cell, val)
	}

	// Tabla
	err := f.AddTable(sheetName, &excelize.Table{
		Range:          "A1:E1001",
		Name:           "TablaUbigeos",
		StyleName:      "TableStyleMedium2",
		ShowRowStripes: boolPtr(true),
	})
	if err != nil {
		return nil, fmt.Errorf("error creando tabla: %w", err)
	}

	// Anchos de columna
	f.SetColWidth(sheetName, "A", "A", 15)
	f.SetColWidth(sheetName, "B", "B", 15)
	f.SetColWidth(sheetName, "C", "C", 25)
	f.SetColWidth(sheetName, "D", "D", 25)
	f.SetColWidth(sheetName, "E", "E", 30)

	// Instrucciones
	f.SetCellValue(sheetName, "A1003", "INSTRUCCIONES:")
	f.SetCellValue(sheetName, "A1004", "1. ubigeo_code: Código de 6 dígitos (obligatorio)")
	f.SetCellValue(sheetName, "A1005", "2. inei_code: Código INEI (opcional)")
	f.SetCellValue(sheetName, "A1006", "3. department, province, district: Nombres (obligatorios)")

	return f, nil
}

func boolPtr(b bool) *bool {
	return &b
}