package Pdoconfig

import (
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdoconfig/SqldbPool_249_0_test"
	"github.com/newgolibs/gostring"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"testing"
)

type SqldbPool_249_0 struct {
}

func TestPdoconfig_SqldbPool_249_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range SqldbPool_249_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], SqldbPool_249_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (SqldbPool_249_0) run(input, arg interface{}) interface{} {
	pdoconfig_a := &goorm.Pdoconfig{}
	pdoconfig_a = goorm.NewPdoconfigFromBytes(input.([]byte))
	pdoconfig_a.MakeDbPool()
	log.Printf("%+v\n", []interface{}{
		pdoconfig_a.LinkString(),
		pdoconfig_a.Sqldb,
		gostring.JsonMarshalIndent(pdoconfig_a.Sqldb.Stats()),
	})

	return reflect.TypeOf(pdoconfig_a.Sqldb).String()
}
