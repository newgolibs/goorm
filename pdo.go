package goorm

import (
	"log"
	"math/rand"
	"time"
)

/**
  执行指定的SQL语句;
*/
func (this *pdo) Exec() {
	stmt, err := this._pdoconfig.SqldbPool().Prepare(this._sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	rand.Seed(time.Now().UnixNano())
	stmt.Exec(this._bindarray)
}

/**
  写入数据;
*/
func (this *pdo) Insert() {
}
