package Pdo

import (
	"encoding/json"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/pdo/Exec_250_0_test"
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
	//配置还原成对象
	var pdoconfig goorm.Pdoconfig;
	json.Unmarshal(input.([]byte),&pdoconfig)
	//生成链接对象
	pdo := goorm.Pdo{Pdoconfig: pdoconfig}
	pdo.Begin()
	var arg2=arg.(map[string]interface{})
	exec := pdo.Exec(arg2["sql"].(string), arg2["binds"].([]interface{}))
	pdo.Commit()
	return exec
}
