package pdoconfig

import (
    "/test/pdoconfig/LinkString_248_0_test"
	"github.com/newgolibs/goorm"
	"log"
    "testing"
    "github.com/stretchr/testify/assert"
    )

type  LinkString_248_0 struct {

}

func Testpdoconfig_LinkString_248_0(t *testing.T){
    log.SetFlags(log.Lshortfile | log.LstdFlags)

    for _, v := range LinkString_248_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], LinkString_248_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (LinkString_248_0) run(input, arg interface{}) interface{} {
	var _pdoconfig = goorm.pdoconfig{DB: input["tns"]}
}