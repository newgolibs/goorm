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

	for _, v := range LinkString_248_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], LinkString_248_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (LinkString_248_0) run(input, arg interface{}) interface{} {
	var input2 = input.(map[string]interface{})
	var a = goorm.Pdoconfig{User: input2["user"].(string), Password: input2["password"].(string), DB: input2["db"].(string), Tns: input2["tns"].(string), Port: input2["port"].(int)}
	return a.LinkString()
}
