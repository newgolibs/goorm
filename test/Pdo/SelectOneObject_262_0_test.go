package Pdo

import (
	"encoding/json"
	"fmt"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/SelectOneObject_262_0_test"
	"github.com/stretchr/testify/assert"
	"log"
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
	pdoconfig.SqldbPoolFromBytes(input.([]byte))
	// 生成链接对象
	pdo := goorm.Pdo{TX: pdoconfig.NewTX()}
	defer pdo.Commit()

	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	var borm = borm{}
	pdo.SelectOneObject(arg2["sql"].(string), arg2["binds"].([]interface{}), &borm)
	fmt.Printf("%+v\n", []interface{}{"borm", borm})
	marshal, _ := json.Marshal(borm)
	return string(marshal)

}
