package Pdo

import (
	"encoding/json"
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
		p:=recover()
		fmt.Printf("%+v\n", []interface{}{p})
	}()

	// 配置还原成对象
	var pdoconfig goorm.Pdoconfig
	json.Unmarshal(input.([]byte), &pdoconfig)
	// 生成链接对象
	pdo := goorm.Pdo{Pdoconfig: pdoconfig}
	pdo.Begin()
	var arg2 = arg.(map[string]interface{})
	num := pdo.Insert(arg2["sql"].(string), arg2["binds"].([]interface{}))
	pdo.Commit()
	return num
}
