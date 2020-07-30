package Pdo

import (
	"encoding/json"
	"fmt"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/SelectOne_252_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"strconv"
	"testing"
)

type SelectOne_252_0 struct {
}

func TestPdo_SelectOne_252_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range SelectOne_252_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], SelectOne_252_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}
func (SelectOne_252_0) run(input, arg interface{}) interface{} {

	// 配置还原成对象
	var pdoconfig goorm.Pdoconfig
	json.Unmarshal(input.([]byte), &pdoconfig)
	// 生成链接对象
	pdo := goorm.Pdo{Pdoconfig: pdoconfig}
	defer pdo.Commit()

	db := pdoconfig.SqldbPool()
	fmt.Printf("%+v\n", []interface{}{"pdoconfig.SqldbPool()", pdoconfig.SqldbPool(), db.Stats()})

	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	v := pdo.SelectOne(arg2["sql"].(string), arg2["binds"].([]interface{}))
	marshal, _ := json.Marshal(v)

	// 赋值
	aormvar := aorm{}
	aormvar.Data_to_struct(v)
	fmt.Printf("%+v\n", []interface{}{aormvar, "<-已经把数据变成结构体对象了"})

	return string(marshal)
}

type aorm struct {
	A1 string `db:"a1"`
	Dd string `db:"dd"`
	Id int    `db:"id"`
}

func (this *aorm) Data_to_struct(data map[string]string) {
	t := reflect.TypeOf(*this)
	vv := reflect.ValueOf(this).Elem() // 为了改变对象的内部值，需使用引用
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Tag.Get("db")
		f := vv.FieldByName(t.Field(i).Name)
		value := data[fieldName]
		if f.Kind() == reflect.Int {
			val, _ := strconv.Atoi(value) // 通过tag获取列数据
			f.SetInt(int64(val))
		} else if f.Kind() == reflect.String {
			f.SetString(value)
		}
	}
}
