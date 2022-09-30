package common

import (
	"encoding/json"
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

func CreateJson(filename string, data interface{}) (err error) {
	// 创建文件
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(f)
	if err = encoder.Encode(data); err != nil {
		return err
	}
	return nil
}

func LoadJson(filename string, v interface{}) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// 创建json解码器
	decoder := json.NewDecoder(f)
	if err = decoder.Decode(&v); err != nil {
		return err
	}
	return nil
}

func GetExec() string {
	exePath, _ := os.Getwd()
	return exePath
}
