package cmd

import (
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

var (
	excelFile *excelize.File
	excelLock sync.Mutex // Ensure goroutine-safe operations on the Excel file
	rowIndex  = 2        // Start from row 2 assuming row 1 will have headers
)

func initExcel() {
	excelFile = excelize.NewFile()
	// Create a new sheet.
	sheetName := "Results"
	excelFile.NewSheet(sheetName)
	// Set headers.
	excelFile.SetCellValue(sheetName, "A1", "Status")
	excelFile.SetCellValue(sheetName, "B1", "Target")
	excelFile.SetCellValue(sheetName, "C1", "Method")
	excelFile.SetCellValue(sheetName, "D1", "Message")

	// Correct handling of sheet activation.
	sheetIndex, err := excelFile.GetSheetIndex(sheetName)
	if err != nil {
		log.Fatal("Failed to get sheet index for Results:", err)
	}
	excelFile.SetActiveSheet(sheetIndex)
}

func saveExcel(fileName string) {
	if err := excelFile.SaveAs(fileName); err != nil {
		log.Fatal("Failed to save Excel file:", err)
	}
}

func writeToExcel(status interface{}, target, method, message string) {
	excelLock.Lock()
	defer excelLock.Unlock()

	sheetName := excelFile.GetSheetName(excelFile.GetActiveSheetIndex())
	cellRef := fmt.Sprintf("A%d", rowIndex)
	excelFile.SetCellValue(sheetName, cellRef, status)
	excelFile.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), target)
	excelFile.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), method)
	excelFile.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), message)

	rowIndex++
}
