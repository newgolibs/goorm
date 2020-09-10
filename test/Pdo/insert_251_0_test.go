package Pdo

import (
	"fmt"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/insert_251_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Insert_251_0 struct {
}

func TestPdo_insert_251_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range insert_251_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Insert_251_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Insert_251_0) run(input, arg interface{}) interface{} {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%+v\n", []interface{}{err, "recover()"})
		}
	}()

	// 配置还原成对象
	var pdoconfig goorm.Pdoconfig
	pdoconfig.SqldbPoolFromBytes(input.([]byte))
	// 生成链接对象
	pdo := goorm.Pdo{TX: pdoconfig.NewTX()}
	defer pdo.Commit()
	var arg2 = arg.(map[string]interface{})
	num, _ := pdo.Insert(arg2["sql"].(string), arg2["binds"].([]interface{}))
	return num
}
