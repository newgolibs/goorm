package goorm

import (
	"reflect"
	"strconv"
)

func Map_to_struct(data map[string]string, struct_ptr interface{}) {
	TypeOfthis := reflect.TypeOf(struct_ptr).Elem()
	ValueOfthis_Elem := reflect.ValueOf(struct_ptr).Elem() // 为了改变对象的内部值，需使用引用

	for i := 0; i < TypeOfthis.NumField(); i++ {
		fieldName := TypeOfthis.Field(i).Tag.Get("db")
		// 没有配置db标签的，略过
		if fieldName == "" {
			continue
		}
		f := ValueOfthis_Elem.FieldByName(TypeOfthis.Field(i).Name)
		if f.Kind() == reflect.Int {
			val, _ := strconv.Atoi(data[fieldName]) // 通过tag获取列数据
			f.SetInt(int64(val))
		} else if f.Kind() == reflect.String {
			f.SetString(data[fieldName])
		} else {
			// panic("未知类型：" + f.Kind().String())
		}
	}

}
