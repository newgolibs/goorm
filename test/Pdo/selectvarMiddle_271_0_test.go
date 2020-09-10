package Pdo

import (
	"fmt"
	"github.com/newgolibs/goorm"
	"github.com/newgolibs/goorm/test/Pdo/selectvarMiddle_271_0_test"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

type SelectvarMiddle_271_0 struct {
}

func TestPdo_selectvarMiddle_271_0(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for _, v := range selectvarMiddle_271_0_test.DataProvider() {
		// fmt.Printf("k=%+v，v=%+v\n", k, v)
		assert.Equal(t, v.([]interface{})[1], SelectvarMiddle_271_0{}.run(v.([]interface{})[0], v.([]interface{})[2]))
	}
}

func (SelectvarMiddle_271_0) run(input, arg interface{}) interface{} {
	pdoconfig := goorm.Pdoconfig{}
	pdoconfig.SqldbPoolFromBytes(input.([]byte))
	pdo := goorm.Pdo{TX: pdoconfig.NewTX()}

	logFile, err := os.OpenFile("darkhold_go.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：darkhold_go", err)
	}
	getwd, _ := os.Getwd()
	fmt.Printf("%+v\n", []interface{}{"当前路径", getwd})
	fmt.Printf("%+v\n", []interface{}{logFile.Name()})
	MainLogger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	fmt.Printf("%+v\n", []interface{}{"===========>Emptymiddleware<==============="})
	Emptymiddleware := goorm.PdoMiddleware{Pdo: pdo, Logger: MainLogger}
	s0, _ := Emptymiddleware.Next_SelectVar("select count(*) from code_logic", []interface{}{})
	fmt.Printf("%+v\n", []interface{}{"code_logic", "=", s0})

	middleware := goorm.PdoMiddleware{Pdo: pdo, Logger: MainLogger}
	fmt.Printf("%+v\n", []interface{}{"middleware.SelectVarindex = ", middleware.SelectVarindex})

	fmt.Printf("%+v\n", []interface{}{"===========>middleware<==============="})
	next := middleware.Add_SelectVar(func(sql string, bindarray []interface{}) (string, error) {
		defer func(start time.Time) {
			tc := time.Since(start)
			fmt.Printf("time cost = %v (%v)\n", tc.Milliseconds(), tc)
		}(time.Now())
		fmt.Printf("%+v\n", []interface{}{"开始执行之前", sql})
		selectVar, err := middleware.Next_SelectVar(sql, bindarray)
		fmt.Printf("%+v\n", []interface{}{"开始执行之后"})
		return selectVar, err
	}, func(sql string, bindarray []interface{}) (string, error) {
		fmt.Printf("%+v\n", []interface{}{"准备暂停一会"})
		time.Sleep(10 * time.Microsecond)
		selectVar, err := middleware.Next_SelectVar(sql, bindarray)
		fmt.Printf("%+v\n", []interface{}{"暂停结束，看数据库链接数", pdoconfig.Sqldb.Stats()})
		return selectVar, err
	}) // 👈👈---- 注意，这里追加了 pdo.SelectVar 源函数，而中间件也追加了一个源函数，但是不会被执行2次。为什么呢？ 因为源函数里面没有调起NEXT!
	s, _ := next(arg.(string), []interface{}{})
	fmt.Printf("%+v\n", []interface{}{arg, "=", s})

	// 重置
	// middleware = goorm.PdoMiddleware{Pdo: pdo}
	fmt.Printf("%+v\n", []interface{}{"===========>middleware.2<==============="})
	middleware.SelectVarindex = 0
	// 复用上一次的中间件
	s2, _ := next("select count(*) from app_setting", []interface{}{})
	fmt.Printf("%+v\n", []interface{}{"app_setting", "=", s2})
	atoi, _ := strconv.Atoi(s)
	return atoi
}
