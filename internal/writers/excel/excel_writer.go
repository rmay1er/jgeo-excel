package excel

import (
	"fmt"
	"log"

	"github.com/rmay1er/jgeo-excel/internal/models"
	"github.com/xuri/excelize/v2"
)

type ExcelWriter struct {
	file *excelize.File
}

func NewExcelWriter() *ExcelWriter {
	return &ExcelWriter{}
}

func (w *ExcelWriter) Write(data *[]models.CordsData, color ...string) error {
	if data == nil || len(*data) == 0 {
		log.Printf("No data provided for writing")
		return fmt.Errorf("нет данных для записи")
	}

	w.file = excelize.NewFile()
	// Create new sheet with name "geojson"
	sheetIndex, err := w.file.NewSheet("geojson")
	if err != nil {
		log.Printf("Error creating new sheet 'geojson': %v", err)
		return err
	}

	// Delete the default sheet
	w.file.DeleteSheet("Sheet1")

	// Set "geojson" as active sheet
	w.file.SetActiveSheet(sheetIndex)

	header := []any{"Тип", "Имя", "Описание", "Координаты"}

	if err := w.file.SetSheetRow("geojson", "A1", &header); err != nil {
		log.Printf("Error setting header row: %v", err)
		return err
	}

	for i, item := range *data {
		row := []any{item.Type, item.IconCaption, item.Description, item.Cords}
		cell := fmt.Sprintf("A%d", i+2)
		if err := w.file.SetSheetRow("geojson", cell, &row); err != nil {
			log.Printf("Error setting row %d: %v", i+2, err)
			return err
		}
	}

	return nil
}

func (w *ExcelWriter) Save(path string) error {
	if err := w.file.SaveAs(path); err != nil {
		log.Printf("Error saving file to path '%s': %v", path, err)
		return err
	}
	return nil
}

func (w *ExcelWriter) Close() error {
	if w.file != nil {
		err := w.file.Close()
		if err != nil {
			log.Printf("Error closing Excel file: %v", err)
			return err
		}
		w.file = nil
	}
	return nil
}
