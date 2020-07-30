package Pdo

import (
	"encoding/json"
	"fmt"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/SelectOneObject_262_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"testing"
)

type SelectOneObject_262_0 struct {
}

func TestPdo_SelectOneObject_262_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range SelectOneObject_262_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], SelectOneObject_262_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (SelectOneObject_262_0) run(input, arg interface{}) interface{} {
	// 配置还原成对象
	var pdoconfig goorm.Pdoconfig
	json.Unmarshal(input.([]byte), &pdoconfig)
	// 生成链接对象
	pdo := goorm.Pdo{Pdoconfig: pdoconfig}
	defer pdo.Commit()

	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	var borm = borm{}
	pdo.SelectOneObject(arg2["sql"].(string), arg2["binds"].([]interface{}), &borm)
	fmt.Printf("%+v\n", []interface{}{"borm", borm})
	marshal, _ := json.Marshal(borm)
	return string(marshal)

}

type borm struct {
	A1 string `db:"a1"`
	Dd string `db:"dd"`
	Id int    `db:"id"`
	goorm.PdoOrm
}

/**
具体的struct对象里面执行反射
*/
func (this *borm) Data_to_struct(data map[string]string) {
	t := reflect.TypeOf(*this)
	vv := reflect.ValueOf(this).Elem() // 为了改变对象的内部值，需使用引用
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Tag.Get("db")
		if fieldName == "" {
			continue
		}
		this.FixValue(vv.FieldByName(t.Field(i).Name), data[fieldName])
	}
}
