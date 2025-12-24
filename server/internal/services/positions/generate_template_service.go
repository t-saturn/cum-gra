package services

import (
	"github.com/xuri/excelize/v2"
)

func GeneratePositionsTemplateExcel() (*excelize.File, error) {
	f := excelize.NewFile()
	sheetName := "Posiciones"
	f.SetSheetName("Sheet1", sheetName)

	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 11, Color: "FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"4472C4"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})

	headers := []string{"name", "code", "level", "description", "is_active", "cod_car_sgd"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, h)
	}
	f.SetCellStyle(sheetName, "A1", "F1", headerStyle)

	example := []interface{}{"Director General", "DG-001", 1, "Cargo directivo principal", "true", "0001"}
	for i, val := range example {
		cell, _ := excelize.CoordinatesToCellName(i+1, 2)
		f.SetCellValue(sheetName, cell, val)
	}

	f.AddTable(sheetName, &excelize.Table{
		Range:          "A1:F1001",
		Name:           "TablaPosiciones",
		StyleName:      "TableStyleMedium2",
		ShowRowStripes: boolPtr(true),
	})

	f.SetColWidth(sheetName, "A", "A", 45)
	f.SetColWidth(sheetName, "B", "B", 15)
	f.SetColWidth(sheetName, "C", "C", 10)
	f.SetColWidth(sheetName, "D", "D", 40)
	f.SetColWidth(sheetName, "E", "E", 12)
	f.SetColWidth(sheetName, "F", "F", 15)

	f.SetCellValue(sheetName, "A1003", "INSTRUCCIONES:")
	f.SetCellValue(sheetName, "A1004", "1. name y code: Obligatorios")
	f.SetCellValue(sheetName, "A1005", "2. level: Nivel jerárquico (1-10)")
	f.SetCellValue(sheetName, "A1006", "3. is_active: true o false")

	return f, nil
}

func boolPtr(b bool) *bool {
	return &b
}