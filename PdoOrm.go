package goorm

import (
	"reflect"
	"strconv"
)

/**    把数值转换成正确的类型，写入到结构体    */
func (PdoOrm) FixValue(f reflect.Value, data string) {
	if f.Kind() == reflect.Int {
		val, _ := strconv.Atoi(data) // 通过tag获取列数据
		f.SetInt(int64(val))
	} else if f.Kind() == reflect.String {
		f.SetString(data)
	} else {
		//panic("未知类型：" + f.Kind().String())
	}

}

/**        */
func (PdoOrm) Data_to_struct(data map[string]string) {
	panic("具体类去实现这个方法，不能直接调用")
}
