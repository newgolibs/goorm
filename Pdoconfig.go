package goorm

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
)

func (this *Pdoconfig) LinkString() string {
	return this.User + ":" + this.Password + "@tcp(" + this.Tns + ":" + strconv.Itoa(this.Port) + ")/" + this.DB + "?charset=utf8"
}

var db *sql.DB
var oncedb sync.Once

/**
和数据库建立持久链接，万一中途被断开了呢？
*/
func (this *Pdoconfig) SqldbPool() *sql.DB {
	oncedb.Do(func() {
		var err error
		db, err = sql.Open("mysql", this.User+":"+this.Password+"@tcp("+this.Tns+":"+strconv.Itoa(this.Port)+")/"+this.DB+"?charset=utf8")
		if err != nil {
			log.Printf("链接数据库错误,配置：%+v", this)
			log.Panic(err)
		}
		//这个是web服务，所以链接上去了，别想着关闭了。
		//defer db.Close()
		// 设置最大连接数
		db.SetMaxOpenConns(10)
		// 设置最大空闲连接数
		db.SetMaxIdleConns(2)
	})
	return db
}
