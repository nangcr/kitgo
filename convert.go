package kitgo

import (
	"strconv"
)

//-------------------------------------
//
//
//
//-------------------------------------
func ConvertFromUnicode(src string) (string, bool) {
	str, err := strconv.Unquote(`"` + src + `"`)
	if err != nil {
		Log(src, " 转换到中文失败 => ", err)
	}
	return str, err != nil
}

//-------------------------------------
//
// 
//
//-------------------------------------
func ConvertToInt(src string) (int, bool) {
	if src == "" {
		return 0, false
	} else {
		r, err := strconv.Atoi(src)
		if err != nil {
			return 0, false
		} else {
			return r, true
		}
	}
}

//-------------------------------------
//
// 
//
//-------------------------------------
func ConvertMustInt(src string) int {
	r, _ := ConvertToInt(src)
	return r
}