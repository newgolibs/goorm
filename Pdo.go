package goorm

import (
	"fmt"
	"log"
)

/**
  执行指定的SQL语句;
*/
func (this *Pdo) Exec(sql string, bindarray []interface{}) int {
	stmt, err := this.Pdoconfig.SqldbPool().Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", []interface{}{"==>", sql, bindarray})
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	Result, err := stmt.Exec(bindarray...)
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	num, _ := Result.RowsAffected()
	return int(num)
}

/**
  写入数据;
*/
func (this *Pdo) Insert() {
}
