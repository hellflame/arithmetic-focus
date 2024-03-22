package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func LowerCamelMarshal(i any) []byte {
	tp := reflect.TypeOf(i)
	switch tp.Kind() {
	case reflect.Pointer:
		return LowerCamelMarshal(reflect.ValueOf(i).Elem().Interface())
	case reflect.Struct:
		val := reflect.ValueOf(i)
		totalFields := tp.NumField()
		tmp := make([]string, totalFields)
		for i := 0; i < totalFields; i++ {
			name := []byte(tp.Field(i).Name)
			if name[0] >= 'A' && name[0] <= 'Z' {
				name[0] += 'a' - 'A'
			}
			tmp[i] = fmt.Sprintf("\"%s\":%s", string(name),
				LowerCamelMarshal(val.Field(i).Interface()))
		}
		return []byte("{" + strings.Join(tmp, ",") + "}")
	case reflect.Slice:
		val := reflect.ValueOf(i)
		size := val.Len()
		tmp := make([]string, size)
		for i := 0; i < size; i++ {
			tmp[i] = string(LowerCamelMarshal(val.Index(i).Interface()))
		}
		return []byte("[" + strings.Join(tmp, ",") + "]")
	default:
		r, e := json.Marshal(i)
		if e != nil {
			panic(e)
		}
		return r
	}
}
