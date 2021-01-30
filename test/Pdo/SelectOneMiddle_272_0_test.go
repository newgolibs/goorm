package Pdo

import (
	"encoding/json"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/SelectOneMiddle_272_0_test"
	zlog "github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type SelectOneMiddle_272_0 struct {
}

func TestPdo_SelectOneMiddle_272_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range SelectOneMiddle_272_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], SelectOneMiddle_272_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

type l struct {
}

/**    输出调试信息    */
func (l) Debug(format string, v ...interface{}) {
	log.Printf("Debug-"+format, v...)
}

/**    输出错误信息    */
func (l) Error(format string, v ...interface{}) {
	log.Printf("Error-"+format, v...)
}

func (SelectOneMiddle_272_0) run(input, arg interface{}) interface{} {
	pdoconfig := goorm.NewPdoconfigFromBytes(input.([]byte))
	//pdoMiddleware := pdoconfig.NewPdoMiddleware(l{})
	pdoMiddleware := pdoconfig.NewPdoMiddleware(&zlog.Logger)
	defer pdoMiddleware.Commit(recover())
	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	log.Printf("测试sql：%+v，测试参数:%+v", arg2["sql"].(string), arg2["binds"].([]interface{}))

	v, _ := pdoMiddleware.SelectOne(arg2["sql"].(string), arg2["binds"].([]interface{}))
	marshal, _ := json.Marshal(v)

	return string(marshal)
}
