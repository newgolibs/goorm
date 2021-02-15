package Pdo

import (
	"encoding/json"
	"fmt"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/SelectOne_252_0_test"
	"github.com/stretchr/testify/assert"
	"log"
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

	log.Printf("input:%#v", string(input.([]byte)))
	// 配置还原成对象
	var pdoconfig *goorm.Pdoconfig = goorm.NewPdoconfigFromBytes(input.([]byte))
	// 生成链接对象
	pdoconfig.MakeDbPool()
	pdo := pdoconfig.NewPdo()
	defer pdo.Commit(recover())

	db := pdoconfig.Sqldb
	fmt.Printf("%+v\n", []interface{}{"pdoconfig.SqldbPool()", pdoconfig.Sqldb, db.Stats()})

	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	log.Printf("测试sql：%+v，测试参数:%+v", arg2["sql"].(string), arg2["binds"].([]interface{}))

	v1, e1 := pdo.SelectOne("select id,a1,dd from  a where id=-1", []interface{}{})
	log.Printf("--->v1:%#v,e1:%#v<----", v1, e1)

	v, _ := pdo.SelectOne(arg2["sql"].(string), arg2["binds"].([]interface{}))
	marshal, _ := json.Marshal(v)

	return string(marshal)
}
