package Pdo

import (
	"encoding/json"
	"github.com/newgolibs/goorm"
    "github.com/newgolibs/goorm/test/Pdo/SelectOneMiddle_272_0_test"
    "log"
    "testing"
    "github.com/stretchr/testify/assert"
    )

type  SelectOneMiddle_272_0 struct {

}

func TestPdo_SelectOneMiddle_272_0(t *testing.T){
    log.SetFlags(log.Lshortfile | log.LstdFlags)

    for _, v := range SelectOneMiddle_272_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], SelectOneMiddle_272_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (SelectOneMiddle_272_0) run(input, arg interface{}) interface{} {
	pdoconfig := goorm.NewPdoconfigFromBytes(input.([]byte))
	middleware := pdoconfig.NewPdoMiddleware()
	// 初始化一个空壳的对象
	var arg2 = arg.(map[string]interface{})
	log.Printf("测试sql：%+v，测试参数:%+v",arg2["sql"].(string),arg2["binds"].([]interface{}))

	v, _ := middleware.Next_SelectOne(arg2["sql"].(string), arg2["binds"].([]interface{}))
	marshal, _ := json.Marshal(v)

	return string(marshal)
}