package goorm

import (
	"database/sql"
)

// 对象必须实现的接口方法
type PdoconfigInterface interface {
	/**    链接数据库，拼接起来的字符串    */
	LinkString() string
	/**    链接池    */
	SqldbPool() *sql.DB
	/**    独立的新的数据库连接池    */
	NewSqldb() *sql.DB
	/**    从json的字符串中，生成数据库连接池对象    */
	SqldbPoolFromBytes(bytes []byte) *sql.DB
	/**    新开事务线程    */
	NewTX() *sql.Tx
}

// 定义函数的结构体，方便扩展成中间件接
type Pdoconfig_LinkStringHandleFunc func() string
type Pdoconfig_SqldbPoolHandleFunc func() *sql.DB
type Pdoconfig_NewSqldbHandleFunc func() *sql.DB
type Pdoconfig_SqldbPoolFromBytesHandleFunc func(bytes []byte) *sql.DB
type Pdoconfig_NewTXHandleFunc func() *sql.Tx

/**
数据库配置;
*/
type Pdoconfig struct {
	/*服务器地址*/
	Tns string
	/*账户*/
	User string
	/*密码*/
	Password string
	/*数据库*/
	DB string
	/*端口*/
	Port int
	/*数据库链接句柄*/
	Sqldb *sql.DB
}

/**
* 中间件的扩展类
 */
type PdoconfigMiddleware struct {
	LinkStringindex               int
	LinkStringHandleFuncs         []Pdoconfig_LinkStringHandleFunc
	SqldbPoolindex                int
	SqldbPoolHandleFuncs          []Pdoconfig_SqldbPoolHandleFunc
	NewSqldbindex                 int
	NewSqldbHandleFuncs           []Pdoconfig_NewSqldbHandleFunc
	SqldbPoolFromBytesindex       int
	SqldbPoolFromBytesHandleFuncs []Pdoconfig_SqldbPoolFromBytesHandleFunc
	NewTXindex                    int
	NewTXHandleFuncs              []Pdoconfig_NewTXHandleFunc
	Pdoconfig                     Pdoconfig
}

func (this *PdoconfigMiddleware) Add_LinkString(middlewares ...Pdoconfig_LinkStringHandleFunc) Pdoconfig_LinkStringHandleFunc {
	if this.LinkStringHandleFuncs == nil {
		this.LinkStringHandleFuncs = make([]Pdoconfig_LinkStringHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.LinkStringHandleFuncs = append(this.LinkStringHandleFuncs, mid)
	}
	return this.Next_LinkString
}
func (this *PdoconfigMiddleware) Next_LinkString() string {
	index := this.LinkStringindex
	if this.LinkStringindex >= len(this.LinkStringHandleFuncs) {
		return ""
	}

	this.LinkStringindex++
	return this.LinkStringHandleFuncs[index]()
}

func (this *PdoconfigMiddleware) Add_SqldbPool(middlewares ...Pdoconfig_SqldbPoolHandleFunc) Pdoconfig_SqldbPoolHandleFunc {
	if this.SqldbPoolHandleFuncs == nil {
		this.SqldbPoolHandleFuncs = make([]Pdoconfig_SqldbPoolHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SqldbPoolHandleFuncs = append(this.SqldbPoolHandleFuncs, mid)
	}
	return this.Next_SqldbPool
}
func (this *PdoconfigMiddleware) Next_SqldbPool() *sql.DB {
	index := this.SqldbPoolindex
	if this.SqldbPoolindex >= len(this.SqldbPoolHandleFuncs) {
		return nil
	}

	this.SqldbPoolindex++
	return this.SqldbPoolHandleFuncs[index]()
}

func (this *PdoconfigMiddleware) Add_NewSqldb(middlewares ...Pdoconfig_NewSqldbHandleFunc) Pdoconfig_NewSqldbHandleFunc {
	if this.NewSqldbHandleFuncs == nil {
		this.NewSqldbHandleFuncs = make([]Pdoconfig_NewSqldbHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.NewSqldbHandleFuncs = append(this.NewSqldbHandleFuncs, mid)
	}
	return this.Next_NewSqldb
}
func (this *PdoconfigMiddleware) Next_NewSqldb() *sql.DB {
	index := this.NewSqldbindex
	if this.NewSqldbindex >= len(this.NewSqldbHandleFuncs) {
		return nil
	}

	this.NewSqldbindex++
	return this.NewSqldbHandleFuncs[index]()
}

func (this *PdoconfigMiddleware) Add_SqldbPoolFromBytes(middlewares ...Pdoconfig_SqldbPoolFromBytesHandleFunc) Pdoconfig_SqldbPoolFromBytesHandleFunc {
	if this.SqldbPoolFromBytesHandleFuncs == nil {
		this.SqldbPoolFromBytesHandleFuncs = make([]Pdoconfig_SqldbPoolFromBytesHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.SqldbPoolFromBytesHandleFuncs = append(this.SqldbPoolFromBytesHandleFuncs, mid)
	}
	return this.Next_SqldbPoolFromBytes
}
func (this *PdoconfigMiddleware) Next_SqldbPoolFromBytes(bytes []byte) *sql.DB {
	index := this.SqldbPoolFromBytesindex
	if this.SqldbPoolFromBytesindex >= len(this.SqldbPoolFromBytesHandleFuncs) {
		return nil
	}

	this.SqldbPoolFromBytesindex++
	return this.SqldbPoolFromBytesHandleFuncs[index](bytes)
}

func (this *PdoconfigMiddleware) Add_NewTX(middlewares ...Pdoconfig_NewTXHandleFunc) Pdoconfig_NewTXHandleFunc {
	if this.NewTXHandleFuncs == nil {
		this.NewTXHandleFuncs = make([]Pdoconfig_NewTXHandleFunc, 0)
	}
	for _, mid := range middlewares {
		this.NewTXHandleFuncs = append(this.NewTXHandleFuncs, mid)
	}
	return this.Next_NewTX
}
func (this *PdoconfigMiddleware) Next_NewTX() *sql.Tx {
	index := this.NewTXindex
	if this.NewTXindex >= len(this.NewTXHandleFuncs) {
		return nil
	}

	this.NewTXindex++
	return this.NewTXHandleFuncs[index]()
}

// 检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PdoconfigInterface = new(Pdoconfig)
