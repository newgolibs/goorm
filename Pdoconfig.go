package goorm

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

/**
mysql专用的链接字符串
*/
func (this *Pdoconfig) LinkString() string {
	return this.User + ":" + this.Password + "@tcp(" + this.Tns + ":" + strconv.Itoa(this.Port) + ")/" + this.DB + "?charset=utf8mb4"
}

/**    返回命令行下的连接字符串    */
func (this *Pdoconfig) ShellLinkString() string {
	return fmt.Sprintf("-h%s -P%d -u%s -p%s --default-character-set=utf8mb4 %s", this.Tns, this.Port, this.User, this.Password, this.DB)
}

/**    生成新的pdo对象    */
func (this *Pdoconfig) NewPdoMiddleware() *PdoMiddleware {
	pdo := &Pdo{TX: this.MakeTX(), Pdoconfig: this}
	return pdo.NewPdoMiddleware()
}

/**    生成新的pdo对象    */
func (this *Pdoconfig) NewPdo() *Pdo {
	return &Pdo{TX: this.MakeTX(), Pdoconfig: this}
}

/**    新开事务线程    */
func (this *Pdoconfig) MakeTX() *sql.Tx {
	if this.Sqldb == nil {
		this.MakeDbPool()
	}
	//log.Printf("打开数据库事务")
	begin, err := this.Sqldb.Begin() // 👈👈----在原来的线程池上，单开一个事务进程
	if err != nil {
		panic(err.Error())
	}
	return begin
}

/**    打开数据库连接    */
func (this *Pdoconfig) MakeSqldb() *Pdoconfig {
	if this.Sqldb == nil {
		//log.Printf("打开数据库池")
		// 这里数据库账户密码，ip，端口。配置错误，都不会导致崩溃。崩溃是产生在查询的时候
		sqldb, err := sql.Open("mysql", this.LinkString())
		if err != nil {
			panic(err.Error())
		}
		this.Sqldb = sqldb
	}
	return this
}

/**
和数据库建立持久链接，万一中途被断开了呢？
*/
func (this *Pdoconfig) MakeDbPool() *Pdoconfig {
	if this.Sqldb == nil {
		this.MakeSqldb()
	}
	// 这个是web服务，所以链接上去了，别想着关闭了。
	// defer db.Close()
	// 设置最大连接数
	this.Sqldb.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	this.Sqldb.SetMaxIdleConns(2)
	return this
}
