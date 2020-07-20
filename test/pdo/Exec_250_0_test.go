package pdo

import (
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
	var input2 = input.(map[string]interface{})
	var a = goorm.Pdoconfig{User: input2["user"].(string), Password: input2["password"].(string), DB: input2["db"].(string), Tns: input2["tns"].(string), Port: input2["port"].(int)}
	db := goorm.Pdo{Pdoconfig: a}
	return db.Exec(input2["sql"].(string), input2["binds"].([]interface{}))
}
