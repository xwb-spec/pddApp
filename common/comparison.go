package common

import (
	"strings"
)

type Comparison struct {
	PicDir *string
	Brand  *string
}

func GetComparison(excelPath, excelSheet string) (ComparisonMap map[string]*Comparison) {
	rows := GetExcelRows(excelPath, excelSheet)
	ComparisonMap = make(map[string]*Comparison)
	for i, row := range rows {
		if i > 0 && len(row) != 0 {
			key := strings.Trim(row[0], " ")
			if key != "" {
				picDir := strings.Trim(row[1], " ")
				brand := strings.Trim(row[2], " ")
				ComparisonMap[strings.Trim(row[0], " ")] = &Comparison{
					PicDir: &picDir,
					Brand:  &brand,
				}
			}
		}
	}
	return
}
