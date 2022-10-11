package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, _ := excelize.OpenFile("./test_info.xlsx")
	sheet1 := f.GetSheetName(0)
	rows, _ := f.GetRows(sheet1)
	for _, row := range rows {
		for _, cell := range row {
			fmt.Println(cell)
		}
	}
}
