package DB

import (
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/DB/Insert_276_0_test"
	zlog "github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type Insert_276_0 struct {
}

func TestDB_Insert_276_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range Insert_276_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Insert_276_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Insert_276_0) run(input, arg interface{}) interface{} {
	pdoconfig := goorm.NewPdoconfigFromBytes(input.([]byte))
	pdoMiddleware := pdoconfig.NewDBMiddleware(&zlog.Logger)

	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	log.Printf("测试sql：%+v，测试参数:%+v", arg2["sql"].(string), arg2["binds"].([]interface{}))
	insert, _ := pdoMiddleware.Insert(arg2["sql"].(string), arg2["binds"].([]interface{}))

	log.Printf("insert:%#v", insert)
	return insert > 0

}
