package stringHelper

import (
	"strings"

	"github.com/spf13/cast"
)

func Int64ArrToStr(arr []int64, delimiter string) string {
	strArr := []string{}
	for _, v := range arr {
		strArr = append(strArr, cast.ToString(v))
	}
	return strings.Join(strArr, delimiter)
}

func ComaStrToInt64Array(str string) []int64 {
	strArr := strings.Split(str, ",")
	intArr := []int64{}
	for _, v := range strArr {
		intArr = append(intArr, cast.ToInt64(v))
	}
	return intArr
}
