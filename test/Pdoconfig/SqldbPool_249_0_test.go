package Pdoconfig

import (
	"fmt"
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
	a := goorm.Pdoconfig{}
	a.SqldbPoolFromBytes(input.([]byte))
	fmt.Printf("%+v\n", []interface{}{a.LinkString(), a.SqldbPool(),
		gostring.JsonMarshalIndent(a.SqldbPool().Stats())})

	b := goorm.Pdoconfig{}
	b.SqldbPoolFromBytes(input.([]byte))
	fmt.Printf("%+v\n", []interface{}{gostring.JsonMarshalIndent(b.NewSqldbPool().Stats())})

	return reflect.TypeOf(a.SqldbPool()).String()
}
