package goorm

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
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

/** 生成数据库查询对象 */
func (this *Pdoconfig) NewDBMiddleware(l *zerolog.Logger) *DBMiddleware {
	PdoconfigMiddlewarevar := &PdoconfigMiddleware{Pdoconfig: this}

	if l == nil { // 👈👈---- 不带日志的查询
		PdoconfigMiddlewarevar.MakeDbPool()
		db := &DB{TX: this.Sqldb, Pdoconfig: PdoconfigMiddlewarevar}
		return db.NewDBMiddleware()
	}
	// 👇👇---- 带日志的查询
	PdoconfigMiddlewarevar.SetZloger(l)
	PdoconfigMiddlewarevar.MakeDbPool()
	db := &DB{TX: this.Sqldb, Pdoconfig: PdoconfigMiddlewarevar}
	DBMiddlewarevar := &DBMiddleware{DB: db}
	DBMiddlewarevar.SetZloger(l)
	return DBMiddlewarevar
}

/**    生成新的pdo对象    */
func (this *Pdoconfig) NewPdoMiddleware(l *zerolog.Logger) *PdoMiddleware {
	PdoconfigMiddlewarevar := &PdoconfigMiddleware{Pdoconfig: this}
	if l == nil { // 👈👈---- 不带日志的查询
		PdoconfigMiddlewarevar.MakeDbPool()
		pdo := &Pdo{TX: PdoconfigMiddlewarevar.MakeTX(), Pdoconfig: PdoconfigMiddlewarevar}
		return pdo.NewPdoMiddleware()
	}
	// 👇👇---- 带日志的查询
	PdoconfigMiddlewarevar.SetZloger(l)
	PdoconfigMiddlewarevar.MakeDbPool()
	pdo := &Pdo{TX: PdoconfigMiddlewarevar.MakeTX(), Pdoconfig: PdoconfigMiddlewarevar}
	PdoMiddlewarevar := &PdoMiddleware{Pdo: pdo}
	PdoMiddlewarevar.SetZloger(l)
	return PdoMiddlewarevar
}

/**    生成新的pdo对象    */
func (this *Pdoconfig) NewPdo() *Pdo {
	return &Pdo{TX: this.MakeTX(), Pdoconfig: &PdoconfigMiddleware{Pdoconfig: this}}
}

/**    新开事务线程    */
func (this *Pdoconfig) MakeTX() *sql.Tx {
	//log.Printf("打开数据库事务")
	begin, err := this.Sqldb.Begin() // 👈👈----在原来的线程池上，单开一个事务进程
	if err != nil {
		panic("MakeTX error:" + err.Error())
	}
	return begin
}

/**
和数据库建立持久链接，万一中途被断开了呢？还是能自动恢复连接的
*/
func (this *Pdoconfig) MakeDbPool() *Pdoconfig {
	if this.Sqldb == nil {
		// 这里数据库账户密码，ip，端口。配置错误，都不会导致崩溃。崩溃是产生在查询的时候
		sqldb, err := sql.Open("mysql", this.LinkString())
		if err != nil {
			panic("MakeDbPool error:" + err.Error())
		}
		this.Sqldb = sqldb
		// 这个是web服务，所以链接上去了，别想着关闭了。
		// defer db.Close()
		// 设置最大连接数
		this.Sqldb.SetMaxOpenConns(10)
		// 设置最大空闲连接数
		this.Sqldb.SetMaxIdleConns(2)
	}
	return this
}
