package goorm
import (
        "reflect"
)

//对象必须实现的接口方法
type PdoOrmInterface interface {
    /**    把数值转换成正确的类型，写入到结构体    */
    FixValue(Value reflect.Value,data string)
    /**        */
    Data_to_struct(data map[string]string)

}

/**
后台生成的表格struct，通用的接口;
*/
type PdoOrm struct
{
}






//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PdoOrmInterface =new(PdoOrm)
