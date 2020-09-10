package goorm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"strconv"
)

/**
mysql专用的链接字符串
*/
func (this *Pdoconfig) LinkString() string {
	return this.User + ":" + this.Password + "@tcp(" + this.Tns + ":" + strconv.Itoa(this.Port) + ")/" + this.DB + "?charset=utf8mb4"
}

/**    独立的新的数据库连接池    */
func (this *Pdoconfig) NewSqldb() *sql.DB {
	// 这里数据库账户密码，ip，端口。配置错误，都不会导致崩溃。崩溃是产生在查询的时候
	sqldb, err := sql.Open("mysql", this.LinkString())
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		panic(fmt.Sprintf("\033[41;36merr:%+v %+v:%+v\033[0m\n", []interface{}{err}, file, line))
	}
	return sqldb
}

/**    新开事务线程    */
func (this *Pdoconfig) NewTX() *sql.Tx {
	begin, err := this.Sqldb.Begin() // 👈👈----在原来的线程池上，单开一个事务进程
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		panic(fmt.Sprintf("\033[0;31merr:%+v %+v:%+v\033[0m\n", []interface{}{err}, file, line))
	}
	return begin
}

/**
和数据库建立持久链接，万一中途被断开了呢？
*/
func (this *Pdoconfig) SqldbPool() *sql.DB {
	this.Sqldb = this.NewSqldb()
	// 这个是web服务，所以链接上去了，别想着关闭了。
	// defer db.Close()
	// 设置最大连接数
	this.Sqldb.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	this.Sqldb.SetMaxIdleConns(2)
	return this.Sqldb
}

/**    从json的字符串中，生成数据库连接池对象    */
func (this *Pdoconfig) SqldbPoolFromBytes(bytes []byte) *sql.DB {
	// 配置还原成对象
	json.Unmarshal(bytes, this)
	this.Sqldb = this.SqldbPool()
	return this.Sqldb
}
