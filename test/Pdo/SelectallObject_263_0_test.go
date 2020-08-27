package Pdo

import (
	"encoding/json"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/SelectallObject_263_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type SelectallObject_263_0 struct {
}

func TestPdo_SelectallObject_263_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range SelectallObject_263_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], SelectallObject_263_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (SelectallObject_263_0) run(input, arg interface{}) interface{} {
	// 配置还原成对象
	var pdoconfig goorm.Pdoconfig
	json.Unmarshal(input.([]byte), &pdoconfig)
	// 生成链接对象
	pdo := goorm.Pdo{Pdoconfig: &pdoconfig}
	defer pdo.Commit()

	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	var bormx []borm
	pdo.SelectallObject(arg2["sql"].(string), arg2["binds"].([]interface{}), &bormx)
	marshal, _ := json.Marshal(bormx)
	return string(marshal)
}
