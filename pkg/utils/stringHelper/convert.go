package stringHelper

import (
	"strings"

	"github.com/spf13/cast"
)

func int64ArrayToComaStr(arr []int64) string {
	strArr := []string{}
	for _, v := range arr {
		strArr = append(strArr, cast.ToString(v))
	}
	return strings.Join(strArr, ",")
}
