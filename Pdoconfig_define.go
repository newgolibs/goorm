package goorm
import (
        "database/sql"
        "sync"
)

//对象必须实现的接口方法
type PdoconfigInterface interface {
    /**    链接数据库，拼接起来的字符串    */
    LinkString()string
    /**    链接池    */
    SqldbPool()*sql.DB
    /**    独立的新的数据库连接池    */
    NewSqldbPool()*sql.DB
    /**    从json的字符串中，生成数据库连接池对象    */
    SqldbPoolFromBytes(bytes []byte)*sql.DB

}

/**
数据库配置;
*/
type Pdoconfig struct
{
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
    sqldb *sql.DB
    /*确保只实例化一次的锁*/
    once_instance sync.Once
}













//检测接口是否被完整的实现了，如果没有实现，那么编译不通过
var _ PdoconfigInterface =new(Pdoconfig)
