package Pdo

import (
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/selectvar_270_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Selectvar_270_0 struct {
}

func TestPdo_selectvar_270_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range selectvar_270_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Selectvar_270_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Selectvar_270_0) run(input, arg interface{}) interface{} {
	// 配置还原成对象
	var pdoconfig goorm.Pdoconfig
	pdoconfig.SqldbPoolFromBytes(input.([]byte))
	// 生成链接对象
	pdo := goorm.Pdo{TX: pdoconfig.NewTX()}
	var a []interface{}
	num, _ := pdo.SelectVar(arg.(string), a)
	return num
}
