package common

import (
	"github.com/xuri/excelize/v2"
	"os"
	"sort"
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
