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

	// 配置还原成对象
	var pdoconfig *goorm.Pdoconfig = goorm.NewPdoconfigFromBytes(input.([]byte))
	// 生成链接对象
	pdo := pdoconfig.NewPdo()
	defer pdo.Commit()

	db := pdoconfig.Sqldb
	fmt.Printf("%+v\n", []interface{}{"pdoconfig.SqldbPool()", pdoconfig.Sqldb, db.Stats()})

	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	v, _ := pdo.SelectOne(arg2["sql"].(string), arg2["binds"].([]interface{}))
	marshal, _ := json.Marshal(v)

	return string(marshal)
}
