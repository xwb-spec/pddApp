package common

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
)

type Comparison struct {
	PicDir string
	Brand  string
}

type Comparisons struct {
	ExcelPath      string
	ExcelSheetName string
	ComparisonMap  map[string]*Comparison
}

func (c *Comparisons) Read() {
	f, err := excelize.OpenFile(c.ExcelPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	rows, err := f.GetRows(c.ExcelSheetName)
	if err != nil {
		log.Println(err)
		return
	}
	c.ComparisonMap = make(map[string]*Comparison)
	for i, row := range rows {
		if i > 0 {
			key := strings.Trim(row[0], " ")
			if key != "" {
				c.ComparisonMap[strings.Trim(row[0], " ")] = &Comparison{
					PicDir: strings.Trim(row[1], " "),
					Brand:  strings.Trim(row[2], " "),
				}
			}
		}
	}
}
