package DB

import (
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/DB/Update_275_0_test"
	zlog "github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Update_275_0 struct {
}

func TestDB_Update_275_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range Update_275_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Update_275_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Update_275_0) run(input, arg interface{}) interface{} {
	pdoconfig := goorm.NewPdoconfigFromBytes(input.([]byte))
	pdoMiddleware := pdoconfig.NewDBMiddleware(&zlog.Logger)
	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	log.Printf("测试sql：%+v，测试参数:%+v", arg2["sql"].(string), arg2["binds"].([]interface{}))
	pdoMiddleware.Exec(arg2["sql"].(string), arg2["binds"].([]interface{}))
	v, _ := pdoMiddleware.SelectVar("select a1 from a where id=1", []interface{}{})
	log.Printf("vvvvv:%#v", v)
	return v
}
