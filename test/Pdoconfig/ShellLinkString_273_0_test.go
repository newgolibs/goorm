package Pdoconfig

import (
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdoconfig/ShellLinkString_273_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type ShellLinkString_273_0 struct {
}

func TestPdoconfig_ShellLinkString_273_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range ShellLinkString_273_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], ShellLinkString_273_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (ShellLinkString_273_0) run(input, arg interface{}) interface{} {
	pdoconfig := goorm.NewPdoconfigFromBytes(input.([]byte))
	linkString := (*pdoconfig).ShellLinkString()
	log.Printf("返回值：%+v", linkString)
	return linkString
}
