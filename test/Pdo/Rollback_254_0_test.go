package Pdo

import (
	"encoding/json"
	"fmt"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/Rollback_254_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type  Rollback_254_0 struct {

}

func TestPdo_Rollback_254_0(t *testing.T){

    log.SetFlags(log.Lshortfile | log.LstdFlags)

    for _, v := range Rollback_254_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Rollback_254_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Rollback_254_0) run(input, arg interface{}) interface{} {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%+v\n", []interface{}{err, "recover()"})
		}
	}()

	// 配置还原成对象
	var pdoconfig goorm.Pdoconfig
	json.Unmarshal(input.([]byte), &pdoconfig)
	// 生成链接对象
	pdo := goorm.Pdo{Pdoconfig: pdoconfig}
	defer pdo.Commit()
	var arg2 = arg.([]map[string]interface{})
	num := pdo.Insert(arg2[0]["sql"].(string), arg2[0]["binds"].([]interface{}))
	fmt.Printf("%+v\n", []interface{}{"11111111111111"})
	pdo.Insert(arg2[1]["sql"].(string), arg2[1]["binds"].([]interface{}))
	fmt.Printf("%+v\n", []interface{}{"22222222222222"})
	pdo.Insert(arg2[2]["sql"].(string), arg2[2]["binds"].([]interface{}))
	fmt.Printf("%+v\n", []interface{}{"333333333333333333"})
	return num
}