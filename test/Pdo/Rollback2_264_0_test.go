package Pdo

import (
	"fmt"
	"github.com/newgolibs/goorm"
    "github.com/newgolibs/goorm/test/Pdo/Rollback2_264_0_test"
    "log"
    "testing"
    "github.com/stretchr/testify/assert"
    )

type  Rollback2_264_0 struct {

}

func TestPdo_Rollback2_264_0(t *testing.T){
    log.SetFlags(log.Lshortfile | log.LstdFlags)

    for _, v := range Rollback2_264_0_test.DataProvider() {
		//fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], Rollback2_264_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (Rollback2_264_0) run(input, arg interface{}) interface{} {
	var a =goorm.Pdoconfig{}
	a.SqldbPoolFromBytes(input.([]byte))
	var pdo =goorm.Pdo{TX: a.NewTX()}
	arg2:=arg.([]map[string]interface{})
	pdo.Exec(arg2[0]["sql"].(string),arg2[0]["binds"].([]interface{}))
	pdo.Rollback()
	num ,_:= pdo.Exec(arg2[1]["sql"].(string), arg2[1]["binds"].([]interface{}))
	one ,_:= pdo.SelectOne("select * from a where id=?", []interface{}{14})
	fmt.Printf("%+v\n", []interface{}{"one=>",one})
	fmt.Printf("%+v\n", []interface{}{a.Sqldb.Stats()})

	var b =goorm.Pdoconfig{}
	b.SqldbPoolFromBytes(input.([]byte))
	var pdo2 =goorm.Pdo{TX: b.NewTX()}
	fmt.Printf("%+v\n", []interface{}{b.Sqldb.Stats()})
	one2 ,_:= pdo2.SelectOne("select * from a where id=?", []interface{}{14})
	fmt.Printf("%+v\n", []interface{}{"one2=>",one2})

	pdo.Commit()

	return num
}