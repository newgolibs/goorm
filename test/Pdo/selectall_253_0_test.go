package Pdo

import (
	"encoding/json"
	"fmt"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/selectall_253_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Selectall_253_0 struct {
}

func TestPdo_selectall_253_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range selectall_253_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Selectall_253_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Selectall_253_0) run(input, arg interface{}) interface{} {
	// 配置还原成对象
	var pdoconfig goorm.Pdoconfig
	pdoconfig.SqldbPoolFromBytes(input.([]byte))
	// 生成链接对象
	pdo := goorm.Pdo{TX: pdoconfig.NewTX()}
	defer pdo.Commit()
	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	v, _ := pdo.SelectAll(arg2["sql"].(string), arg2["binds"].([]interface{}))
	marshal, _ := json.Marshal(v)
	fmt.Printf("%+v\n", []interface{}{string(marshal)})
	return string(marshal)
}
