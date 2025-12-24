package services

import (
	"fmt"
	"server/internal/config"
	"server/internal/models"

	"github.com/xuri/excelize/v2"
)

func GenerateOrganicUnitsTemplateExcel() (*excelize.File, error) {
	db := config.DB
	f := excelize.NewFile()

	sheetNames := []string{"Unidades Orgánicas", "Unidades Existentes"}
	f.SetSheetName("Sheet1", sheetNames[0])
	f.NewSheet(sheetNames[1])

	// Estilo encabezados
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 11, Color: "FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"4472C4"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center", Vertical: "center"},
	})

	// === HOJA 1: PLANTILLA ===
	mainSheet := sheetNames[0]
	headers := []string{"name", "acronym", "brand", "description", "parent_id", "is_active", "cod_dep_sgd"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(mainSheet, cell, h)
	}
	f.SetCellStyle(mainSheet, "A1", "G1", headerStyle)

	// Ejemplo
	example := []interface{}{"Gerencia Regional de Desarrollo Social", "GRDS", "", "Gerencia encargada de...", "", "true", "001"}
	for i, val := range example {
		cell, _ := excelize.CoordinatesToCellName(i+1, 2)
		f.SetCellValue(mainSheet, cell, val)
	}

	f.AddTable(mainSheet, &excelize.Table{
		Range:          "A1:G1001",
		Name:           "TablaUnidades",
		StyleName:      "TableStyleMedium2",
		ShowRowStripes: boolPtr(true),
	})

	f.SetColWidth(mainSheet, "A", "A", 50)
	f.SetColWidth(mainSheet, "B", "B", 15)
	f.SetColWidth(mainSheet, "C", "C", 20)
	f.SetColWidth(mainSheet, "D", "D", 40)
	f.SetColWidth(mainSheet, "E", "E", 12)
	f.SetColWidth(mainSheet, "F", "F", 12)
	f.SetColWidth(mainSheet, "G", "G", 15)

	// === HOJA 2: UNIDADES EXISTENTES ===
	refSheet := sheetNames[1]
	var units []models.OrganicUnit
	db.Where("is_deleted = FALSE").Order("id ASC").Find(&units)

	refHeaders := []string{"ID", "Nombre", "Acrónimo", "ID Padre"}
	for i, h := range refHeaders {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(refSheet, cell, h)
	}
	f.SetCellStyle(refSheet, "A1", "D1", headerStyle)

	for idx, unit := range units {
		row := idx + 2
		f.SetCellValue(refSheet, fmt.Sprintf("A%d", row), unit.ID)
		f.SetCellValue(refSheet, fmt.Sprintf("B%d", row), unit.Name)
		f.SetCellValue(refSheet, fmt.Sprintf("C%d", row), unit.Acronym)
		if unit.ParentID != nil {
			f.SetCellValue(refSheet, fmt.Sprintf("D%d", row), *unit.ParentID)
		}
	}

	if len(units) > 0 {
		f.AddTable(refSheet, &excelize.Table{
			Range:          fmt.Sprintf("A1:D%d", len(units)+1),
			Name:           "TablaReferencia",
			StyleName:      "TableStyleMedium9",
			ShowRowStripes: boolPtr(true),
		})
	}

	f.SetColWidth(refSheet, "A", "A", 8)
	f.SetColWidth(refSheet, "B", "B", 50)
	f.SetColWidth(refSheet, "C", "C", 15)
	f.SetColWidth(refSheet, "D", "D", 12)

	f.SetActiveSheet(0)
	return f, nil
}

func boolPtr(b bool) *bool {
	return &b
}