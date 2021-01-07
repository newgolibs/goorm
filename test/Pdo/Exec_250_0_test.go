package Pdo

import (
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/Exec_250_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Exec_250_0 struct {
}

func TestPdo_Exec_250_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range Exec_250_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Exec_250_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Exec_250_0) run(input, arg interface{}) interface{} {
	// 配置还原成对象,新开一个数据库事务
	var pdo  = goorm.NewPdoconfigFromBytes(input.([]byte)).NewPdo()
	defer pdo.Commit(recover())
	var arg2 = arg.(map[string]interface{})
	log.Printf("测试sql：%+v，测试参数:%+v",arg2["sql"].(string),arg2["binds"].([]interface{}))
	exec, _ := pdo.Exec(arg2["sql"].(string), arg2["binds"].([]interface{}))
	log.Printf("数据库情况:%+v",pdo.Pdoconfig.Sqldb.Stats())
	return exec
}
