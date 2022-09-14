package common

import (
	"github.com/xuri/excelize/v2"
	"os"
	"sort"
	"strings"
)

func IsPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsEleExistsSlice(ele string, strSlice []string) bool {
	sort.Strings(strSlice)
	index := sort.SearchStrings(strSlice, ele)
	//index的取值：[0,len(str_array)]
	if index < len(strSlice) && strSlice[index] == ele { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}
func IsSheetExists(excelPath, sheetName string) bool {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		return false
	}
	defer f.Close()
	idx := f.GetSheetIndex(sheetName)
	if idx < 0 {
		return false
	}
	return true
}

func SearchSheet(excelPath, sheetName, key string) (values []string, err error) {
	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	locList, err := f.SearchSheet(sheetName, key, true)
	if err != nil {
		return nil, err
	}
	for _, l := range locList {
		if strings.HasPrefix(l, "A") {
			lb := strings.Replace(l, "A", "B", 1)
			lc := strings.Replace(l, "A", "C", 1)
			bv, err := f.GetCellValue(sheetName, lb)
			if err != nil {
				return nil, err
			}
			cv, err := f.GetCellValue(sheetName, lc)
			if err != nil {
				return nil, err
			}
			return []string{bv, cv}, nil
		}
	}
	return nil, err
}
