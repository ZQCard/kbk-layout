package typeConvert

import (
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.Indirect(reflect.ValueOf(obj))
	objType := objValue.Type()

	// 如果 obj 不是结构体类型，则返回空的 map
	if objType.Kind() != reflect.Struct {
		return nil
	}

	resultMap := make(map[string]interface{})

	for i := 0; i < objValue.NumField(); i++ {
		// 获取结构体字段名和字段值
		field := objType.Field(i)
		fieldValue := objValue.Field(i)
		if field.Name == "state" || field.Name == "sizeCache" || field.Name == "unknownFields" {
			continue
		}
		resultMap[field.Name] = fieldValue.Interface()
	}

	return resultMap
}
