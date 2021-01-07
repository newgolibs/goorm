package Pdoconfig

import (
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdoconfig/LinkString_248_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type LinkString_248_0 struct {
}

func TestPdoconfig_LinkString_248_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for k, v := range LinkString_248_0_test.DataProvider() {
		//log.Printf("k=%+v，v=%+v\n", k, v)
		log.Printf("第%d个测试",k)
		assert.Equal(t, v.([]interface{})[1], LinkString_248_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (LinkString_248_0) run(input, arg interface{}) interface{} {
	pdoconfig := goorm.NewPdoconfigFromBytes(input.([]byte))
	log.Printf("数据库配置,%+v",*pdoconfig)
	linkString := (*pdoconfig).LinkString()
	log.Printf("返回值：%+v",linkString)
	return linkString
}
